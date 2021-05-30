package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/dstotijn/go-notion"
	"go.uber.org/zap"
)

func main() {

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
		sugar.Fatalf("Unable to find database: %v", err)
	}

	sugar.Infow("retrieved Notion database",
		"dbID", dbID,
		"Title", database.Title[0].PlainText,
	)

	nextVote, err := FindVoteTarget(context.Background(), client, dbID)
	if err != nil {
		sugar.Fatalf("Error finding a vote target: %v", err)
	}

	sugar.Infow("vote found",
		"name", nextVote.name,
		"imageURL", nextVote.imageURL,
	)
}

type emoji struct {
	name     string
	imageURL string
	aliasFor string
}

// FindVoteTarget queries the datastore for an emoji to vote on. It retrieves a single emoji that has the lowest number of total votes
func FindVoteTarget(ctx context.Context, client *notion.Client, dbID string) (emoji, error) {
	// Query for a single emoji that fits "least number of total votes"
	query := notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Property: "Name",
			Text: &notion.TextDatabaseQueryFilter{
				Equals: "000",
			},
			Number: &notion.NumberDatabaseQueryFilter{
				Equals: notion.IntPtr(0),
			},
		},
		Sorts: []notion.DatabaseQuerySort{{
			Property:  "Total Votes",
			Direction: "descending",
		}},
		StartCursor: "",
		PageSize:    1,
	}

	response, err := client.QueryDatabase(ctx, dbID, &query)
	if err != nil {
		return emoji{}, err
	}
	if len(response.Results) == 0 {
		return emoji{}, errors.New("did not find a vote target")
	}
	if len(response.Results) > 1 {
		return emoji{}, errors.New("found too many vote targets")
	}

	emojiPage := response.Results[0]
	emojiMap := emojiPage.Properties.(notion.DatabasePageProperties)

	nameProperty, ok := emojiMap["Name"]
	if !ok {
		return emoji{}, errors.New("returned page does not have section 'name'")
	}
	if nameProperty.Type != notion.DBPropTypeTitle {
		return emoji{}, fmt.Errorf("returned 'Name' property does not have Type %s", notion.DBPropTypeTitle)
	}

	imageProperty, ok := emojiMap["Image"]
	if !ok {
		return emoji{}, errors.New("returned page does not have section 'Image'")
	}
	if imageProperty.Type != notion.DBPropTypeFiles {
		return emoji{}, fmt.Errorf("returned 'Image' property does not have Type %s", notion.DBPropTypeFiles)
	}

	return emoji{
		name:     nameProperty.Title[0].PlainText,
		imageURL: imageProperty.Files[0].Name,
	}, nil
}
