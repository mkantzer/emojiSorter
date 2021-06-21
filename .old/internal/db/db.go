package db

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/dstotijn/go-notion"
	"go.uber.org/zap"
)

// This folder defines the database structure, and how it can be passed around?

type NotionDB struct {
	client *notion.Client
	dbID   string
}

// New returns a new NotionDB struct
func New() (NotionDB, error) {

	// create notion database client
	notionAPIKey, ok := os.LookupEnv("NOTION_API_KEY")
	if !ok {
		return NotionDB{}, errors.New("env var NOTION_API_KEY not set")
	}
	client := notion.NewClient(notionAPIKey)

	dbID, ok := os.LookupEnv("NOTION_DB_ID")
	if !ok {
		return NotionDB{}, errors.New("env var NOTION_DB_ID not set")
	}

	return NotionDB{
		client: client,
		dbID:   dbID,
	}, nil
}

// QueryDatabase wraps the *notion.client.QueryDatabase call for our db struct.
// It adds some useful debugger logging, (metrics handling and gathering), and (something else).
// It does NOT modify the DatabaseQueryResponse or error, and returns them directly.
func (n NotionDB) queryDatabase(ctx context.Context, query *notion.DatabaseQuery) (notion.DatabaseQueryResponse, error) {
	start := time.Now()
	zlog := zap.S().With(
		"dbType", "notion",
		"dbID", n.dbID,
		// note: this is probably a bad idea for security reasons eventually
		"query", query,
	)
	zlog.Debug("querying datastore")
	response, err := n.client.QueryDatabase(ctx, n.dbID, query)
	if err != nil {
		zlog.Errorw("query failed",
			"duration", time.Since(start),
			"error", err.Error(),
		)
	}
	zlog.Debugw("query suceeded",
		"duration", time.Since(start),
		"numResults", len(response.Results),
	)
	return response, err
}
