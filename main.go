package main

import (
	"net/http"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"github.com/zikaeroh/strawrank/internal/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var args struct {
	Debug bool `long:"debug" env:"SR_DEBUG" description:"Enables debug mode, including extra routes and logging"`
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

	// TODO: flag
	secureKey := []byte("a-32-byte-long-key-goes-here")

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

	a, err := app.New(&app.Config{
		Logger:       logger,
		CookieKey:    secureKey,
		HIDMinLength: 5,
		HIDSalt:      "PJSalt",
	})
	if err != nil {
		logger.Fatal("creating app", zap.Error(err))
	}

	if err := http.ListenAndServe(":3000", a); err != nil {
		logger.Fatal("exiting", zap.Error(err))
	}
}
