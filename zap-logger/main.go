package zaplogger

import (
	"os"

	"go.uber.org/zap"
)

func main() {
	// config
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		"stdout",          // print logs in console
		"./logs/log.json", // print logs in json file
	}
	config.Encoding = "json"

	// create logger
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("Application started", zap.String("module", "main"))

	fakeDataLogs(logger)
}

func fakeDataLogs(logger *zap.Logger) {
	// info log
	logger.Info("Processing user data",
		zap.String("username", "ali-daei"),
		zap.Int("userID", 1234),
		zap.String("status", "active"),
	)

	// warning log
	logger.Warn("API rate limit reached",
		zap.String("api_endpoint", "/getUserData"),
		zap.Int("userID", 1234),
		zap.Int("attempts", 5),
	)

	// error log
	logger.Error("Failed to connect to database",
		zap.String("db_host", "localhost"),
		zap.Int("retry_count", 3),
		zap.Error(os.ErrNotExist),
	)

	// debug log
	logger.Debug("Debugging session start",
		zap.String("debug_id", "dbg123"),
		zap.String("description", "Starting debug session for testing purposes"),
	)

	// fatal log
	logger.Fatal("Critical error, shutting down",
		zap.String("module", "main"),
		zap.String("reason", "Unable to recover from critical error"),
	)
}
