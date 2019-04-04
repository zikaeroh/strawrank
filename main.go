package main

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"github.com/zikaeroh/strawrank/internal/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var args = struct {
	CookieKey string `long:"cookie-key" env:"SR_COOKIE_KEY" description:"Cookie encryption key, hex encoded" required:"true"`

	HIDMinLength int    `long:"hid-min-length" env:"SR_HID_MIN_LENGTH" description:"HashID minimum length"`
	HIDSalt      string `long:"hid-salt" env:"SR_HID_SALT" description:"HashID salt"`

	Debug bool `long:"debug" env:"SR_DEBUG" description:"Enables debug mode, including extra routes and logging"`
}{
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

	a, err := app.New(&app.Config{
		Logger:       logger,
		CookieKey:    cookieKey,
		HIDMinLength: args.HIDMinLength,
		HIDSalt:      args.HIDSalt,
	})
	if err != nil {
		logger.Fatal("creating app", zap.Error(err))
	}

	if err := http.ListenAndServe(":3000", a); err != nil {
		logger.Fatal("exiting", zap.Error(err))
	}
}
