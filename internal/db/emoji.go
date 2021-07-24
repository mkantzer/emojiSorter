package db

import (
	"context"
	"fmt"

	"github.com/dstotijn/go-notion"
	"github.com/mkantzer/emojiSorter/internal/core"
	"github.com/mkantzer/emojiSorter/pkg/apperrors"
)

// GetEmojiByName returns a single emoji with the given name
func (n NotionDB) GetEmojiByName(ctx context.Context, name string) (core.Emoji, error) {
	// Query for a single emoji with a given name
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
	response, err := n.queryDatabase(ctx, &query)
	if err != nil {
		return core.Emoji{}, err
	}
	if len(response.Results) == 0 {
		return core.Emoji{}, apperrors.ErrEmojiNotFound
	}
	if len(response.Results) > 1 {
		return core.Emoji{}, fmt.Errorf("found %d emoji with this name", len(response.Results))
	}

	// Extract Emoji Data
	emojiPage := response.Results[0]
	emojiData, err := n.extractEmojiFromPage(ctx, emojiPage)
	if err != nil {
		return core.Emoji{}, fmt.Errorf("error extracting emoji data from page: %w", err)
	}
	return emojiData, nil
}

func (n NotionDB) GetAllEmoji(ctx context.Context) ([]core.Emoji, error) {
	n.Deps.Logger.Debug("Whelp, guess I gotta go get all the emoji")

	// Query for all emoji
	query := notion.DatabaseQuery{
		Sorts: []notion.DatabaseQuerySort{
			{
				Property:  "Name",
				Timestamp: "created_time",
				Direction: "ascending",
			},
		},
		StartCursor: "",
		PageSize:    100,
	}

	response, err := n.queryDatabase(ctx, &query)
	if err != nil {
		return []core.Emoji{}, err
	}

	// loop through response, and extract the emoji
	allEmoji := []core.Emoji{}

	for _, page := range response.Results {
		emojiData, err := n.extractEmojiFromPage(ctx, page)
		if err != nil {
			return []core.Emoji{}, fmt.Errorf("error extracting emoji data from page id %s: %w", page.ID, err)
		}
		allEmoji = append(allEmoji, emojiData)
	}

	// handle pagination
	for response.HasMore {
		n.Deps.Logger.Sugar().Infow("query indicates more data available",
			"count", len(allEmoji),
		)
		query.StartCursor = *response.NextCursor
		response, err = n.queryDatabase(ctx, &query)
		if err != nil {
			return []core.Emoji{}, err
		}
		for _, page := range response.Results {
			emojiData, err := n.extractEmojiFromPage(ctx, page)
			if err != nil {
				return []core.Emoji{}, fmt.Errorf("error extracting emoji data from page id %s: %w", page.ID, err)
			}
			allEmoji = append(allEmoji, emojiData)
		}
	}
	n.Deps.Logger.Sugar().Infow("getAllEmoji complete",
		"count", len(allEmoji),
	)
	return allEmoji, nil
}
