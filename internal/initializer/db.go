package initializer

import (
	"errors"
	"os"

	"github.com/mkantzer/emojiSorter/internal/db"
	"go.uber.org/zap"
)

func NotionDatabase(logger *zap.Logger) (db.NotionDB, error) {
	notionAPIKey, ok := os.LookupEnv("NOTION_API_KEY")
	if !ok {
		return db.NotionDB{}, errors.New("env var NOTION_API_KEY not set")
	}

	dbID, ok := os.LookupEnv("NOTION_DB_ID")
	if !ok {
		return db.NotionDB{}, errors.New("env var NOTION_DB_ID not set")
	}

	database, err := db.NewDatabase(&db.Dependencies{
		Logger: logger,
	}, dbID, notionAPIKey)
	if err != nil {
		return db.NotionDB{}, err
	}

	return database, err

}
