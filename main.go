package main

import (
	"net/http"

	"github.com/zikaeroh/strawrank/internal/app"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TODO: flag
const debug = true

func main() {
	// TODO: flag
	secureKey := []byte("a-32-byte-long-key-goes-here")

	var logConfig zap.Config

	if debug {
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
		Logger:    logger,
		CookieKey: secureKey,
		HIDSalt:   "PJSalt",
	})
	if err != nil {
		logger.Fatal("creating app", zap.Error(err))
	}

	if err := http.ListenAndServe(":3000", a); err != nil {
		logger.Fatal("exiting", zap.Error(err))
	}
}
