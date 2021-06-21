package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/dstotijn/go-notion"
	"github.com/mkantzer/emojiSorter/internal/core"
)

// FindVoteTarget queries the datastore for an emoji to vote on.
// It returns an Emoji that has the lowest number of total votes.
func (n NotionDB) FindVoteTarget(ctx context.Context) (core.Emoji, error) {
	// Query for a single emoji that fits "least number of total votes"
	query := notion.DatabaseQuery{
		Sorts: []notion.DatabaseQuerySort{{
			Property:  "Total Votes",
			Direction: "descending",
		}},
		StartCursor: "",
		PageSize:    1,
	}

	n.Deps.Logger.Sugar().Debug("finding next vote target")
	response, err := n.queryDatabase(ctx, &query)
	if err != nil {
		n.Deps.Logger.Sugar().Debugf("error finding next vote target: %w", err)
		return core.Emoji{}, err
	}
	if len(response.Results) == 0 {
		n.Deps.Logger.Sugar().Debug("returned zero results")
		return core.Emoji{}, errors.New("did not find a vote target")
	}
	if len(response.Results) > 1 {
		n.Deps.Logger.Sugar().Debug("returned too many resulsts")
		return core.Emoji{}, errors.New("found too many vote targets")
	}

	// Extract Emoji Data
	emojiPage := response.Results[0]
	emojiData, err := n.extractEmojiFromPage(ctx, emojiPage)
	if err != nil {
		return core.Emoji{}, fmt.Errorf("error extracting emoji data from page: %w", err)
	}
	return emojiData, nil
}
