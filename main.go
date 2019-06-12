package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/zikaeroh/strawrank/internal/app"
	"github.com/zikaeroh/strawrank/internal/db/migrations"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	_ "github.com/lib/pq" // For PostgreSQL support.
)

var args = struct {
	Addr string `long:"addr" env:"SR_ADDR" description:"Address to listen at"`

	CookieKey string `long:"cookie-key" env:"SR_COOKIE_KEY" description:"Cookie encryption key, hex encoded"`

	HIDMinLength int    `long:"hid-min-length" env:"SR_HID_MIN_LENGTH" description:"HashID minimum length"`
	HIDSalt      string `long:"hid-salt" env:"SR_HID_SALT" description:"HashID salt"`

	Database      string `long:"database" env:"SR_DATABASE" description:"Database connection string" required:"true"`
	MigrateUp     bool   `long:"migrate-up" env:"SR_MIGRATE_UP" description:"Migrate the database up before starting"`
	MigrateReset  bool   `long:"migrate-reset" env:"SR_MIGRATE_RESET" description:"Reset the database before starting"`
	DatabaseDebug bool   `long:"database-debug" env:"SR_DATABASE_DEBUG" description:"Enable SQLBoiler debug logging"`

	RealIP bool `long:"real-ip" env:"SR_REAL_IP" description:"Enable RealIP middleware to read IP from headers"`

	Debug bool `long:"debug" env:"SR_DEBUG" description:"Enables debug mode, including extra routes and logging"`
}{
	Addr:         ":3000",
	CookieKey:    "612D33322D627974652D6C6F6E672D6B65792D676F65732D68657265", // a-32-byte-long-key-goes-here
	HIDMinLength: 5,
	HIDSalt:      "PJSalt",
}

func main() {
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}

	if _, err := flags.Parse(&args); err != nil {
		// Default flag parser prints messages, so just exit.
		os.Exit(1)
	}

	cookieKey, err := hex.DecodeString(args.CookieKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error decoding cookie key: %v", err)
		os.Exit(1)
	}

	if args.MigrateUp && args.MigrateReset {
		fmt.Fprintf(os.Stderr, "migrate-up and migrate-reset cannot both be set")
		os.Exit(1)
	}

	var logConfig zap.Config

	if args.Debug {
		logConfig = zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		logConfig = zap.NewProductionConfig()
	}

	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}

	if args.DatabaseDebug {
		r, w := io.Pipe()

		boil.DebugMode = true
		boil.DebugWriter = w

		go func() {
			scanner := bufio.NewScanner(r)

			for scanner.Scan() {
				line := scanner.Text()
				line = strings.TrimSpace(line)

				if line == "" {
					continue
				}

				logger.Debug("sqlboiler: " + line)
			}
		}()
	}

	logger.Info("starting")

	undoStdlog := zap.RedirectStdLog(logger)
	defer undoStdlog()

	var db *sql.DB
	connected := false

	for i := 0; !connected && i < 10; i++ {
		db, err = sql.Open("postgres", args.Database)
		if err != nil {
			logger.Error("error opening database connection", zap.Error(err))
			time.Sleep(20 * time.Second)
			continue
		}

		if err := db.Ping(); err != nil {
			logger.Error("error pinging database connection", zap.Error(err))
			time.Sleep(20 * time.Second)
			continue
		}

		connected = true
	}

	if !connected {
		logger.Fatal("database could not be reached")
	}

	defer db.Close()

	debugf := func(format string, v ...interface{}) {
		logger.Sugar().Infof("migrate: "+strings.TrimSpace(format), v...)
	}

	switch {
	case args.MigrateUp:
		err = migrations.Up(args.Database, debugf)
	case args.MigrateReset:
		err = migrations.Reset(args.Database, debugf)
	}

	if err != nil {
		logger.Fatal("error migrating database", zap.Error(err))
	}

	a, err := app.New(&app.Config{
		Logger:       logger,
		DB:           db,
		CookieKey:    cookieKey,
		HIDMinLength: args.HIDMinLength,
		HIDSalt:      args.HIDSalt,
		RealIP:       args.RealIP,
		Debug:        args.Debug,
	})
	if err != nil {
		logger.Fatal("creating app", zap.Error(err))
	}

	srv := http.Server{
		Addr:    args.Addr,
		Handler: a,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		logger.Info("shutting down server")

		if err := srv.Shutdown(context.Background()); err != nil {
			logger.Error("error shutting down server", zap.Error(err))
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			logger.Error("listen and serve error", zap.Error(err))
		}
	}

	<-idleConnsClosed
	logger.Info("server shut down")
}
