/*
Package ports contains the interface definitions used to communicate with actors.

It does not actively implement them.
*/
package ports

import "github.com/mkantzer/emojiSorter/internal/core/domain"

type EmojiRepository interface {
	Get(name string) (domain.Emoji, error)
	Save(domain.Emoji) error
}

type EmojiService interface {
	Get(name string) (domain.Emoji, error)
	Vote(name string, v domain.Vote) error
}
