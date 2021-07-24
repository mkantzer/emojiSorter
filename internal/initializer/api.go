package initializer

import (
	"fmt"

	"github.com/mkantzer/emojiSorter/internal/api"

	"os"
	"strconv"

	"go.uber.org/zap"
)

func ApiServer(logger *zap.Logger) (*api.Server, error) {
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
		Logger: logger,
	}, fmt.Sprintf("0.0.0.0:%d", port))

	return server, nil
}
