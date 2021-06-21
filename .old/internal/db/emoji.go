package db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/dstotijn/go-notion"
)

type Emoji struct {
	ID       string
	Name     string
	ImageURL string
	AliasFor string
}

// FindVoteTarget queries the datastore for an emoji to vote on.
// It returns an Emoji that has the lowest number of total votes
func (n NotionDB) FindVoteTarget(ctx context.Context) (Emoji, error) {
	// Query for a single emoji that fits "least number of total votes"
	query := notion.DatabaseQuery{
		Sorts: []notion.DatabaseQuerySort{{
			Property:  "Total Votes",
			Direction: "descending",
		}},
		StartCursor: "",
		PageSize:    1,
	}

	response, err := n.queryDatabase(ctx, &query)
	if err != nil {
		return Emoji{}, err
	}
	if len(response.Results) == 0 {
		return Emoji{}, errors.New("did not find a vote target")
	}
	if len(response.Results) > 1 {
		return Emoji{}, errors.New("found too many vote targets")
	}

	// Extract Emoji Data
	emojiPage := response.Results[0]
	emojiData, err := n.extractEmojiFromPage(ctx, emojiPage)
	if err != nil {
		return Emoji{}, fmt.Errorf("error extracting emoji data from page: %w", err)
	}
	return emojiData, nil
}

func (n NotionDB) GetEmojiByName(ctx context.Context, name string) (Emoji, error) {
	// Query for a single emoji that fits "least number of total votes"
	query := notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Property: "Name",
			Text: &notion.TextDatabaseQueryFilter{
				Equals: name,
			},
		},
		StartCursor: "",
		PageSize:    1,
	}
	response, err := n.queryDatabase(ctx, &query)
	if err != nil {
		return Emoji{}, err
	}
	if len(response.Results) == 0 {
		return Emoji{}, errors.New("did not find an emoji with this name")
	}
	if len(response.Results) > 1 {
		return Emoji{}, fmt.Errorf("found %d emoji with this name", len(response.Results))
	}

	// Extract Emoji Data
	emojiPage := response.Results[0]
	emojiData, err := n.extractEmojiFromPage(ctx, emojiPage)
	if err != nil {
		return Emoji{}, fmt.Errorf("error extracting emoji data from page: %w", err)
	}
	return emojiData, nil
}

func (n NotionDB) extractEmojiFromPage(ctx context.Context, emojiPage notion.Page) (Emoji, error) {
	emojiMap := emojiPage.Properties.(notion.DatabasePageProperties)

	nameProperty, ok := emojiMap["Name"]
	if !ok {
		return Emoji{}, errors.New("returned page does not have section 'Name'")
	}
	if nameProperty.Type != notion.DBPropTypeTitle {
		return Emoji{}, fmt.Errorf("returned 'Name' property does not have Type %s", notion.DBPropTypeTitle)
	}
	nameHolder := nameProperty.Title[0].PlainText

	imageProperty, ok := emojiMap["Image"]
	if !ok {
		return Emoji{}, errors.New("returned page does not have section 'Image'")
	}
	if imageProperty.Type != notion.DBPropTypeFiles {
		return Emoji{}, fmt.Errorf("returned 'Image' property does not have Type %s", notion.DBPropTypeFiles)
	}

	imageURLHolder := imageProperty.Files[0].Name
	aliasHolder := ""

	// If the image URL starts with "alias:", we need to go find the ACTUAL image URL
	if strings.HasPrefix(imageURLHolder, "alias:") {
		aliasHolder = imageProperty.Files[0].Name
		aliasedEmoji, err := n.GetEmojiByName(ctx, strings.TrimPrefix(imageURLHolder, "alias:"))
		if err != nil {
			return Emoji{}, fmt.Errorf(
				"error retrieving emoji %s aliased by %s: %w",
				strings.TrimPrefix(imageURLHolder, "alias:"),
				nameHolder,
				err,
			)
		}
		imageURLHolder = aliasedEmoji.ImageURL
	}

	return Emoji{
		ID:       emojiPage.ID,
		Name:     nameHolder,
		ImageURL: imageURLHolder,
		AliasFor: aliasHolder,
	}, nil
}
