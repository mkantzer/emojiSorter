package main

import (
	"context"
	"log"
	"os"

	"github.com/dstotijn/go-notion"
	"github.com/mkantzer/emojiSorter/db"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar().With(
		"app", "emojiSorter",
		"version", os.Getenv("APP_VERSION"),
	)
	sugar.Info("logger constrcution succeeded")

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

	nextVote, err := db.FindVoteTarget(context.Background(), client, dbID)
	if err != nil {
		sugar.Fatalf("Error finding a vote target: %v", err)
	}

	sugar.Infow("vote found",
		"emojiID", nextVote.ID,
		"name", nextVote.Name,
		"imageURL", nextVote.ImageURL,
		"AliasFor", nextVote.AliasFor,
	)
}
