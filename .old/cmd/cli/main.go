package main

import (
	"context"
	"log"
	"os"

	"github.com/mkantzer/emojiSorter/internal/db"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type app_info struct {
	App     string
	Version string
}

func (a app_info) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("app", a.App)
	enc.AddString("version", a.Version)
	return nil
}

func main() {

	/*
		############################################################
		################## LOGGER CONFIG ###########################
		############################################################
	*/

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	zap.ReplaceGlobals(logger)
	defer logger.Sync() // flushes buffer, if any

	app := &app_info{
		App:     "emojiSorter",
		Version: os.Getenv("APP_VERSION"),
	}

	sugar := logger.Sugar().With(
		zap.Object("app_info", app))
	sugar.Info("logger constrcution succeeded")

	/*
		############################################################
		################## DATASTORE CONFIG ########################
		############################################################
	*/

	emojiDataBase, err := db.New()
	if err != nil {
		sugar.Fatalf("Error setting up database client: %w", err)
	}

	/*
		############################################################
		################## DATASTORE USAGE? ########################
		############################################################
	*/

	nextVote, err := emojiDataBase.FindVoteTarget(context.Background())
	if err != nil {
		sugar.Fatalf("Error finding a vote target: %v", err)
	}

	sugar.Infow("vote found",
		"emojiID", nextVote.ID,
		"name", nextVote.Name,
		"imageURL", nextVote.ImageURL,
		"AliasFor", nextVote.AliasFor,
	)

}
