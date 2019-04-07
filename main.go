package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"github.com/zikaeroh/strawrank/internal/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var args = struct {
	Addr string `long:"addr" env:"SR_ADDR" description:"Address to listen at"`

	CookieKey string `long:"cookie-key" env:"SR_COOKIE_KEY" description:"Cookie encryption key, hex encoded"`

	HIDMinLength int    `long:"hid-min-length" env:"SR_HID_MIN_LENGTH" description:"HashID minimum length"`
	HIDSalt      string `long:"hid-salt" env:"SR_HID_SALT" description:"HashID salt"`

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

	logger.Info("starting")

	undoStdlog := zap.RedirectStdLog(logger)
	defer undoStdlog()

	a, err := app.New(&app.Config{
		Logger:       logger,
		CookieKey:    cookieKey,
		HIDMinLength: args.HIDMinLength,
		HIDSalt:      args.HIDSalt,
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
