package db

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dstotijn/go-notion"
	"github.com/mkantzer/emojiSorter/internal/core"
	"go.uber.org/zap"
)

type Dependencies struct {
	Logger *zap.Logger
}

type NotionDB struct {
	Deps *Dependencies
	DbID string

	client *notion.Client
}

type Emojistore interface {
	GetEmojiByName(context.Context, string) (core.Emoji, error)
	// GetAllEmoji(context.Context) ([]core.Emoji, error)
	FindVoteTarget(context.Context) (core.Emoji, error)
}

func NewDatabase(deps *Dependencies, dbID string, apiKey string) (NotionDB, error) {
	emojiCodeMap = emojiCode()
	return NotionDB{
		Deps:   deps,
		DbID:   dbID,
		client: notion.NewClient(apiKey),
	}, nil
}

// queryDatabase wraps the *notion.client.QueryDatabase call for our db struct.
// It adds some useful debugger logging, (metrics handling and gathering), and (something else).
// It does NOT modify the DatabaseQueryResponse or error, and returns them directly.
func (n NotionDB) queryDatabase(ctx context.Context, query *notion.DatabaseQuery) (notion.DatabaseQueryResponse, error) {
	start := time.Now()
	n.Deps.Logger.Sugar().Debugw("notion query start",
		"query", query,
	)
	response, err := n.client.QueryDatabase(ctx, n.DbID, query)
	if err != nil {
		n.Deps.Logger.Sugar().Error("query failed",
			"duration", time.Since(start),
			"error", err.Error(),
		)
		return notion.DatabaseQueryResponse{}, err
	}

	n.Deps.Logger.Sugar().Debugw("query suceeded",
		"duration", time.Since(start),
		"numResults", len(response.Results),
	)
	return response, nil
}

func (n NotionDB) extractEmojiFromPage(ctx context.Context, emojiPage notion.Page) (core.Emoji, error) {
	emojiMap := emojiPage.Properties.(notion.DatabasePageProperties)

	nameProperty, ok := emojiMap["Name"]
	if !ok {
		return core.Emoji{}, errors.New("returned page does not have section 'Name'")
	}
	if nameProperty.Type != notion.DBPropTypeTitle {
		return core.Emoji{}, fmt.Errorf("returned 'Name' property does not have Type %s", notion.DBPropTypeTitle)
	}
	nameHolder := nameProperty.Title[0].PlainText

	imageProperty, ok := emojiMap["Image"]
	if !ok {
		return core.Emoji{}, errors.New("returned page does not have section 'Image'")
	}
	if imageProperty.Type != notion.DBPropTypeFiles {
		return core.Emoji{}, fmt.Errorf("returned 'Image' property does not have Type %s", notion.DBPropTypeFiles)
	}

	imageURLHolder := imageProperty.Files[0].Name
	aliasHolder := ""

	// If the image URL starts with "alias:", we need to go find the ACTUAL image URL
	if strings.HasPrefix(imageURLHolder, "alias:") {
		aliasHolder = imageProperty.Files[0].Name

		aliasName := strings.TrimPrefix(imageURLHolder, "alias:")
		// check if predefined. If so, return unicode code
		if code, ok := emojiCodeMap[":"+aliasName+":"]; ok {
			imageURLHolder = code

		} else {
			aliasedEmoji, err := n.GetEmojiByName(ctx, strings.TrimPrefix(imageURLHolder, "alias:"))
			if err != nil {
				return core.Emoji{}, fmt.Errorf(
					"error retrieving emoji %s aliased by %s: %w",
					strings.TrimPrefix(imageURLHolder, "alias:"),
					nameHolder,
					err,
				)
			}
			imageURLHolder = aliasedEmoji.ImageURL
		}
	}

	return core.Emoji{
		Name:     nameHolder,
		ImageURL: imageURLHolder,
		AliasFor: aliasHolder,
	}, nil
}
