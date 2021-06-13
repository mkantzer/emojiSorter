package main

import (
	"fmt"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/mkantzer/emojiSorter/internal/core/services/emojisrv"
	"github.com/mkantzer/emojiSorter/internal/handlers/emojihdl"
	"github.com/mkantzer/emojiSorter/internal/repositories/emojirepo"
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

var (
	listenAddr  string
	storageType string
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {

	/*
		############################################################
		################## LOGGER CONFIG ###########################
		############################################################
	*/

	logger, err := zap.NewProduction()
	if err != nil {
		return fmt.Errorf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync() // flushes buffer, if any

	app := &app_info{
		App:     "emojiSorter",
		Version: os.Getenv("APP_VERSION"),
	}

	sugar := logger.Sugar().With(
		zap.Object("app_info", app))

	zap.ReplaceGlobals(logger)

	sugar.Info("logger constrcution succeeded")

	/*
		############################################################
		################### EMOJI CONFIG ###########################
		############################################################
	*/

	repo, err := emojirepo.NewNotionDB()
	if err != nil {
		return err
	}
	srv := emojisrv.New(repo)

	hdl := emojihdl.NewHTTPHandler(srv)

	/*
		############################################################
		##################### GIN CONFIG ###########################
		############################################################
	*/

	r := gin.New()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.GET("/emoji/:name", hdl.Get)
	r.GET("/healthz", hdl.Health)
	r.Run(":8080")
	return nil
}

// func run() error {
// 	flag.StringVar(&listenAddr, "listen-addr", ":5000", "server listen address")
// 	flag.Parse()

// 	store := new(memory.Storage)

// 	var lister listing.Service
// 	var adder adding.Service
// 	var opener opening.Service
// 	var remover removing.Service

// 	lister = listing.NewService(store)
// 	adder = adding.NewService(store)
// 	opener = opening.NewService(store)
// 	remover = removing.NewService(store)

// 	//seed the database
// 	adder.AddSampleItem(seed.PopulateItems())

// 	// set up the HTTP server
// 	h := rest.NewHandlers(lister, adder, opener, remover)
// 	server := h.GetServer(listenAddr)

// 	//channel to listen for errors coming from the listener.
// 	serverErrors := make(chan error, 1)

// 	go func() {
// 		log.Printf("main : API listening on %s", listenAddr)
// 		serverErrors <- server.ListenAndServe()
// 	}()

// 	//channel to listen for an interrupt or terminate signal from the OS.
// 	shutdown := make(chan os.Signal, 1)
// 	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

// 	//blocking run and waiting for shutdown.
// 	select {
// 	case err := <-serverErrors:
// 		return fmt.Errorf("error: starting server: %s", err)

// 	case <-shutdown:
// 		log.Println("main : Start shutdown")

// 		//give outstanding requests a deadline for completion.
// 		const timeout = 5 * time.Second
// 		ctx, cancel := context.WithTimeout(context.Background(), timeout)
// 		defer cancel()

// 		// asking listener to shutdown
// 		err := server.Shutdown(ctx)
// 		if err != nil {
// 			log.Printf("main : Graceful shutdown did not complete in %v : %v", timeout, err)
// 			err = server.Close()
// 		}

// 		if err != nil {
// 			return fmt.Errorf("main : could not stop server gracefully : %v", err)
// 		}
// 	}

// 	return nil
// }
