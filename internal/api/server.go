package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mkantzer/emojiSorter/internal/db"
	"go.uber.org/zap"
)

type Dependencies struct {
	Logger   *zap.Logger
	Database db.Emojistore
}

type Server struct {
	Deps *Dependencies
	Addr string

	server *http.Server
}

func NewServer(deps *Dependencies, addr string) *Server {
	return &Server{
		Deps: deps,
		Addr: addr,
	}
}

func (a *Server) Start() {
	r := gin.New()
	r.GET("/", HelloServer)
	r.GET("/healthz", HealthCheck)
	r.GET("/unhealthz", UnhealthCheck)
	r.GET("/bruh", a.testLog)

	a.server = &http.Server{
		Addr:         a.Addr,
		Handler:      r,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	go func() {
		err := a.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			a.Deps.Logger.Error(err.Error())
		}
	}()
}

func (a *Server) Shutdown() {
	if a.server == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := a.server.Shutdown(ctx)
	if err != nil {
		a.Deps.Logger.Error(err.Error())
	}
}
