package lumberjack

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// loggs config for lumberjack
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "./logs/app.log", // route for logg file
		MaxSize:    5,                // max file size
		MaxBackups: 3,                // max file backup number
		MaxAge:     28,               // max number of days saving backup file
		Compress:   true,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(lumberjackLogger),
		zap.InfoLevel,
	)

	logger := zap.New(core)
	defer logger.Sync() // finally sync the logs

	logger.Info("Application started")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")

	for i := 0; i < 1000; i++ {
		logger.Info("Logging fake data",
			zap.Int("iteration", i),
			zap.Time("timestamp", time.Now()),
			zap.String("status", "running"),
		)
	}

	logger.Info("Application finished")
}
