package main

import (
	app "inventory/app"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, _ := loggerConfig.Build()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Infof("Initialized App Log test: %s", "L34")

	app.Run()
}
