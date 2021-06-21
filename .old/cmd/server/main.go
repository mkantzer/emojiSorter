package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/mkantzer/emojiSorter/internal/initializer"
)

const serviceName = "emojiSorter"

// Will be overwitten by build process
var gitHash = "development"

func main() {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		log.Print("ENVIRONMENT not set, using development")
		environment = "development"
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(fmt.Errorf("problem getting hostname: %v", err))
	}

	logger, err := initializer.Logging(environment, hostname, serviceName, gitHash)
	if err != nil {
		log.Fatal(fmt.Errorf("problem setting up logger: %v", err))
	}

	// Flush/sync logs when we exit
	defer func() {
		_ = logger.Sync() // flushes buffer, if any
	}()

	apiServer, err := initializer.ApiServer(logger)
	if err != nil {
		logger.Fatal(fmt.Errorf("problem setting up api server: %v", err).Error())
	}

	logger.Info("staring api server")
	apiServer.Start()

	// Setup signal capture
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Wait for interupt signal
	<-stop

	logger.Info("shuting down api server")
	apiServer.Shutdown()

	logger.Info("gracefully shutdown")
}
