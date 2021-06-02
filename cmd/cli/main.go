package main

import (
	"fmt"
	"log"

	"github.com/mkantzer/emojiSorter/internal/core/domain"
)

func main() {
	emoji := domain.Emoji{
		Name:     "",
		ImageURL: "",
		AliasFor: "",
		Votes:    map[domain.Vote]uint{},
	}
	log.Println(emoji)
	fmt.Println(emoji.Vote(domain.Rename))
	fmt.Println(emoji.Vote(domain.Transparent))
	fmt.Println(emoji.Vote(domain.Delete))
	fmt.Println(emoji.Vote(10))
	emoji.Votes[domain.Keep] = 234
	fmt.Println(emoji.Vote(domain.Delete))
	fmt.Println(emoji.Vote(domain.Delete))
	fmt.Println(emoji.Vote(domain.Delete))
	log.Println(emoji)

}
