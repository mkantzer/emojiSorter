package core

import (
	"context"
)

type Emoji struct {
	Name     string `json:"name"`
	ImageURL string `json:"imageurl"`
	AliasFor string `json:"aliasfor"`
	// Unicode  string
	Votes map[Vote]uint `json:"votes"`
}

// EmojiStore defines the behaviors needed to be a valid emoji storage tool.
type EmojiStore interface {
	GetEmoji(context.Context, string) (Emoji, error)
	GetAllEmoji(context.Context) ([]Emoji, error)
	GetNextVote(context.Context) (Emoji, error)
	Vote(context.Context, Emoji, Vote) error
}
