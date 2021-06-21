package initializer

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mkantzer/emojiSorter/internal/api"
	"github.com/mkantzer/emojiSorter/internal/db"
	"go.uber.org/zap"
)

func ApiServer(logger *zap.Logger, database db.NotionDB) (*api.Server, error) {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		logger.Info("PORT not set, using 8080")
		portStr = "8080"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("problem getting port: %v", err)
	}

	server := api.NewServer(&api.Dependencies{
		Logger:   logger,
		Database: database,
	}, fmt.Sprintf("0.0.0.0:%d", port))

	return server, nil
}
