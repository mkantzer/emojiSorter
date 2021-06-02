package api

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

func DummyHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	zlog := zap.S().With(
		// zap.Object("req", reqLog{Name: "alice"}),
		"handler", "dummy",
	)
	zlog.Info("i'm right gere my dude")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": "banana"}`)
}

// func ElectionHandler(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	io.WriteString(w, `{"emoji": "look-at-me"}`)

// }

// https://www.nicolasmerouze.com/share-values-between-middlewares-context-golang
// Passing logs around with something like this.
