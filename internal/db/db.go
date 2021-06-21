package db

import (
	"context"
	"time"

	"github.com/dstotijn/go-notion"
	"go.uber.org/zap"
)

type Dependencies struct {
	Logger *zap.Logger
}

type NotionDB struct {
	Deps *Dependencies
	DbID string

	client *notion.Client
}

func NewDatabase(deps *Dependencies, dbID string, apiKey string) (NotionDB, error) {
	return NotionDB{
		Deps:   deps,
		DbID:   dbID,
		client: notion.NewClient(apiKey),
	}, nil
}

// queryDatabase wraps the *notion.client.QueryDatabase call for our db struct.
// It adds some useful debugger logging, (metrics handling and gathering), and (something else).
// It does NOT modify the DatabaseQueryResponse or error, and returns them directly.
func (n NotionDB) queryDatabase(ctx context.Context, query *notion.DatabaseQuery) (notion.DatabaseQueryResponse, error) {
	start := time.Now()
	n.Deps.Logger.Sugar().Debugw("notion query start",
		"query", query,
	)
	response, err := n.client.QueryDatabase(ctx, n.DbID, query)
	if err != nil {
		n.Deps.Logger.Sugar().Error("query failed",
			"duration", time.Since(start),
			"error", err.Error(),
		)
		return notion.DatabaseQueryResponse{}, err
	}

	n.Deps.Logger.Sugar().Debugw("query suceeded",
		"duration", time.Since(start),
		"numResults", len(response.Results),
	)
	return response, nil
}
