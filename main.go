package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewDevelopment()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()

	sugar.Info("logger constrcution succeeded")

	// sugar.Infow("failed to fetch URL",
	// 	// Structured context as loosely typed key-value pairs.
	// 	"url", url,
	// 	"attempt", 3,
	// 	"backoff", time.Second,
	// )

	// 	sugar.Infof("Failed to fetch URL: %s", url)

}
