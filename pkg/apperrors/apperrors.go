package apperrors

import "errors"

var (
	ErrInvalidVote   = errors.New("invalid vote type")
	ErrEmojiNotFound = errors.New("emoji not found")

// NotFound         =
// IllegalOperation =
// InvalidInput     =
// Internal         =
)
