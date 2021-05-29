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
	// Querry database for all non-voted emoji
	query := notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Property: "Total Votes",
			Number: &notion.NumberDatabaseQueryFilter{
				Equals: notion.IntPtr(0),
			},
		},
		// Sorts:       []notion.DatabaseQuerySort{},
		StartCursor: "",
		PageSize:    0,
	}

	response, err := client.QueryDatabase(context.Background(), dbID, &query)
	if err != nil {
		sugar.Errorw("query failed",
			"error", err,
			"query", query,
			"dbID", dbID,
		)
	}
	sugar.Infow("query succeeded",
		"numResults", len(response.Results),
	)
	// sugar.Infow("Retrieved non-voted emoji",

}
