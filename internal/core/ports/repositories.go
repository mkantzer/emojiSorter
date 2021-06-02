package ports

import "github.com/mkantzer/emojiSorter/internal/core/domain"

type EmojiRepository interface {
	Get(name string) (domain.Emoji, error)
	Save(domain.Emoji) error
}
