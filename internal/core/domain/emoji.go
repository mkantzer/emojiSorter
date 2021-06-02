/*
Package domain contains the go struct definitions and functions on each entity that is part of the domain problem.

The functions and structures here are used across the application, and should not contain other internal dependancies.
*/

package domain

import (
	"github.com/mkantzer/emojiSorter/pkg/apperrors"
)

// Votes being exported probably isn't great here.
type Emoji struct {
	// ID       string
	Name     string
	ImageURL string
	AliasFor string
	Votes    map[Vote]uint
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
