package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dstotijn/go-notion"
	"github.com/gorilla/mux"
	"github.com/mkantzer/emojiSorter/db"
	"go.uber.org/zap"
)

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
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar().With(
		"app", "emojiSorter",
		"version", os.Getenv("APP_VERSION"),
	)
	sugar.Info("logger constrcution succeeded")

	/*
		############################################################
		################## DATASTORE CONFIG ########################
		############################################################
	*/

	// create notion database client
	notionAPIKey, ok := os.LookupEnv("NOTION_API_KEY")
	if !ok {
		sugar.Fatal("NOTION_API_KEY not set")
	}
	client := notion.NewClient(notionAPIKey)

	dbID, ok := os.LookupEnv("NOTION_DB_ID")
	if !ok {
		sugar.Fatal("NOTION_DB_ID not set")
	}

	/*
		############################################################
		################## DATASTORE USAGE? ########################
		############################################################
	*/

	nextVote, err := db.FindVoteTarget(context.Background(), client, dbID)
	if err != nil {
		sugar.Fatalf("Error finding a vote target: %v", err)
	}

	sugar.Infow("vote found",
		"emojiID", nextVote.ID,
		"name", nextVote.Name,
		"imageURL", nextVote.ImageURL,
		"AliasFor", nextVote.AliasFor,
	)

	/*
		############################################################
		################## SERVER CONFIG ########################
		############################################################
	*/
	// for graceful shutdown
	wait := time.Duration(time.Second * 5)

	r := mux.NewRouter()
	r.HandleFunc("/healthz", HealthCheckHandler)
	r.HandleFunc("/", DummyHandler)
	// r.HandleFunc("/", VotingHandler)
	// r.HandleFunc("/emoji/{name}", SpecificEmojiHandler)
	// r.HandleFunc("/results", ResultHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			sugar.Errorf("server shutdown with error:", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	sugar.Info("listening on port 8080")

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	sugar.Info("shutting down")
	sugar.Sync()
	os.Exit(0)
}
