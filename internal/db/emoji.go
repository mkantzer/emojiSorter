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
