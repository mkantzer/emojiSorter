package db

import (
	"context"

	"github.com/mkantzer/emojiSorter/internal/core/emoji"
)

// FindVoteTarget queries the datastore for an emoji to vote on.
// It returns an Emoji that has the lowest number of total votes.
func (n NotionDB) FindVoteTarget(ctx context.Context) (emoji.Emoji, error) {}
