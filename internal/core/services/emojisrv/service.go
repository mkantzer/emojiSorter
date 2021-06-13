/*
Package emojisrv contains the entrypoint and implementation of the EmojiService port
*/

package emojisrv

import (
	"fmt"

	"github.com/mkantzer/emojiSorter/internal/core/domain"
	"github.com/mkantzer/emojiSorter/internal/core/ports"
)

type service struct {
	er ports.EmojiRepository
}

func New(er ports.EmojiRepository) *service {
	return &service{
		er: er,
	}
}

func (srv *service) Get(name string) (domain.Emoji, error) {
	emoji, err := srv.er.Get(name)
	if err != nil {
		return domain.Emoji{}, fmt.Errorf("get emoji from repository has failed: %w", err)
	}
	return emoji, nil
}

func (srv *service) GetAll() ([]domain.Emoji, error) {
	emoji, err := srv.er.GetAll()
	if err != nil {
		return []domain.Emoji{}, fmt.Errorf("get all emoji from repostiory has failed: %w", err)
	}
	return emoji, nil
}

func (srv *service) Vote(name string, v domain.Vote) error {
	emoji, err := srv.Get(name)
	if err != nil {
		return err
	}

	if err := emoji.Vote(v); err != nil {
		return err
	}

	if err := srv.er.Save(emoji); err != nil {
		return fmt.Errorf("error saving vote: %w", err)
	}

	return nil
}

// Get the next emoji to vote on
func (srv *service) Next() (domain.Emoji, error) {
	emoji, err := srv.er.Next()
	if err != nil {
		return domain.Emoji{}, fmt.Errorf("get next vote from repository has failed: %w", err)
	}
	return emoji, nil
}
