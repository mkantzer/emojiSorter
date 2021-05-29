package main

import (
	"context"
	"log"
	"os"

	"github.com/dstotijn/go-notion"
	"go.uber.org/zap"
)

func main() {

	// logger, err := zap.NewDevelopment()
	logger, err := zap.NewProduction()

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

	// create notion database client
	notionAPIKey, ok := os.LookupEnv("NOTION_API_KEY")
	if !ok {
		sugar.Fatal("NOTION_API_KEY not set")
	}
	client := notion.NewClient(notionAPIKey)

	dbID, ok := os.LookupEnv("NOTION_DB_ID")
	if !ok {
		sugar.Fatal("NOTION_DB_ID not set")
	}

	// get the database page
	database, err := client.FindDatabaseByID(context.Background(), dbID)
	if err != nil {
		sugar.Errorf("Unable to find database: %v", err)
	}

	sugar.Infow("retrieved Notion database",
		"dbID", dbID,
		"Title", database.Title[0].PlainText,
	)
}
