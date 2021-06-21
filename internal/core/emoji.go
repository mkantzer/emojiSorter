package core

import (
	"github.com/mkantzer/emojiSorter/pkg/apperrors"
)

type Emoji struct {
	Name     string `json:"name"`
	ImageURL string `json:"imageurl"`
	AliasFor string `json:"aliasfor"`
	// Unicode  string
	Votes map[Vote]uint `json:"votes"`
}

type Vote uint

const (
	vote_beg Vote = iota
	Keep
	Delete
	Transparent
	Rename
	Duplicate
	vote_end
)

// IsVote returns true for values corresponding to votes; it returns false otherwise.
func (v Vote) IsVote() bool { return vote_beg < v && v < vote_end }

// Vote
func (emoji Emoji) Vote(vote Vote) error {
	if vote.IsVote() {
		emoji.Votes[vote] += 1
		return nil
	}
	return apperrors.ErrInvalidVote
}

func (v Vote) String() string {
	return [...]string{
		"vote_beg",
		"Keep",
		"Delete",
		"Transparent",
		"Rename",
		"Duplicate",
		"vote_end",
	}[v]
}

// maybe do one of these: https://stackoverflow.com/questions/54735113/how-to-use-json-string-value-to-get-iota-value
