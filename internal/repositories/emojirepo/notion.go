package emojirepo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dstotijn/go-notion"
	"github.com/mkantzer/emojiSorter/internal/core/domain"
	"github.com/mkantzer/emojiSorter/pkg/apperrors"
	"go.uber.org/zap"
)

type notionDB struct {
	client *notion.Client
	dbID   string
}

// New returns a new NotionDB struct
func NewNotionDB() (notionDB, error) {

	// create notion database client
	notionAPIKey, ok := os.LookupEnv("NOTION_API_KEY")
	if !ok {
		return notionDB{}, errors.New("env var NOTION_API_KEY not set")
	}
	client := notion.NewClient(notionAPIKey)

	dbID, ok := os.LookupEnv("NOTION_DB_ID")
	if !ok {
		return notionDB{}, errors.New("env var NOTION_DB_ID not set")
	}

	return notionDB{
		client: client,
		dbID:   dbID,
	}, nil
}

func (repo *notionDB) Get(ctx context.Context, name string) (domain.Emoji, error) {
	// Query for a single emoji that fits "least number of total votes"
	query := notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Property: "Name",
			Text: &notion.TextDatabaseQueryFilter{
				Equals: name,
			},
		},
		StartCursor: "",
		// PageSize:    0,
	}
	response, err := repo.queryDatabase(ctx, &query)
	if err != nil {
		return domain.Emoji{}, err
	}
	if len(response.Results) == 0 {
		return domain.Emoji{}, apperrors.ErrEmojiNotFound
	}
	if len(response.Results) > 1 {
		return domain.Emoji{}, fmt.Errorf("found %d emoji with this name", len(response.Results))
	}

	// Extract Emoji Data
	emojiPage := response.Results[0]
	emojiData, err := repo.extractEmojiFromPage(ctx, emojiPage)
	if err != nil {
		return domain.Emoji{}, fmt.Errorf("error extracting emoji data from page: %w", err)
	}
	return emojiData, nil
}

func (n notionDB) extractEmojiFromPage(ctx context.Context, emojiPage notion.Page) (domain.Emoji, error) {
	emojiMap := emojiPage.Properties.(notion.DatabasePageProperties)

	nameProperty, ok := emojiMap["Name"]
	if !ok {
		return domain.Emoji{}, errors.New("returned page does not have section 'Name'")
	}
	if nameProperty.Type != notion.DBPropTypeTitle {
		return domain.Emoji{}, fmt.Errorf("returned 'Name' property does not have Type %s", notion.DBPropTypeTitle)
	}
	nameHolder := nameProperty.Title[0].PlainText

	imageProperty, ok := emojiMap["Image"]
	if !ok {
		return domain.Emoji{}, errors.New("returned page does not have section 'Image'")
	}
	if imageProperty.Type != notion.DBPropTypeFiles {
		return domain.Emoji{}, fmt.Errorf("returned 'Image' property does not have Type %s", notion.DBPropTypeFiles)
	}

	imageURLHolder := imageProperty.Files[0].Name
	aliasHolder := ""

	// If the image URL starts with "alias:", we need to go find the ACTUAL image URL
	if strings.HasPrefix(imageURLHolder, "alias:") {
		aliasHolder = imageProperty.Files[0].Name
		aliasedEmoji, err := n.Get(ctx, strings.TrimPrefix(imageURLHolder, "alias:"))
		if err != nil {
			return domain.Emoji{}, fmt.Errorf(
				"error retrieving emoji %s aliased by %s: %w",
				strings.TrimPrefix(imageURLHolder, "alias:"),
				nameHolder,
				err,
			)
		}
		imageURLHolder = aliasedEmoji.ImageURL
	}

	return domain.Emoji{
		// ID:       emojiPage.ID,
		Name:     nameHolder,
		ImageURL: imageURLHolder,
		AliasFor: aliasHolder,
	}, nil
}

// queryDatabase wraps the *notion.client.QueryDatabase call for our db struct.
// It adds some useful debugger logging, (metrics handling and gathering), and (something else).
// It does NOT modify the DatabaseQueryResponse or error, and returns them directly.
func (n notionDB) queryDatabase(ctx context.Context, query *notion.DatabaseQuery) (notion.DatabaseQueryResponse, error) {
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