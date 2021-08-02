package db

import (
	"context"

	"github.com/dstotijn/go-notion"
	"github.com/mkantzer/emojiSorter/internal/core"
)

type notionDB interface {
	queryDB(ctx context.Context, query *notion.DatabaseQuery) (notion.DatabaseQueryResponse, error)
	extractEmojiFromPage(ctx context.Context, emojiPage notion.Page) (core.Emoji, error)
	getNextVote(ctx context.Context) (core.Emoji, error)
}
