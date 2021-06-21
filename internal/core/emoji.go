package core

import (
	"github.com/mkantzer/emojiSorter/pkg/apperrors"
)

type Emoji struct {
	Name     string
	ImageURL string
	AliasFor string
	// Unicode  string
	Votes map[Vote]uint
}

type Vote uint

const (
	vote_beg Vote = iota
	Keep
	Delete
	Transparent
	Rename
	Dupelicate
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
