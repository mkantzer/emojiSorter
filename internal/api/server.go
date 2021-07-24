package api

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	// "github.com/newrelic/go-agent/v3/newrelic"
)

const shutdownTimeout = 30 * time.Second

type Dependencies struct {
	Logger *zap.Logger
	// Application *newrelic.Application
}

type Server struct {
	Deps *Dependencies
	Addr string

	server *http.Server
}

// Create new server
func NewServer(deps *Dependencies, addr string) *Server {
	return &Server{
		Deps: deps,
		Addr: addr,
	}
}

// Listen and server requests
func (s *Server) Start() {
	// mux := http.NewServeMux()
	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.RequestID)
	chiRouter.Use(middleware.RealIP)
	chiRouter.Use(s.ZapLogger)

	chiRouter.Get("/", s.HelloServer)
	chiRouter.Get("/healthz", HealthCheck)
	chiRouter.Get("/unhealthz", UnhealthCheck)

	// By listening outside of the serve goroutine we
	// avoid a race condition in our tests
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		s.Deps.Logger.Fatal(err.Error())
	}

	s.server = &http.Server{
		Addr:         s.Addr,
		Handler:      chiRouter,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	go func() {
		err := s.server.Serve(listener)
		if err != nil && err != http.ErrServerClosed {
			s.Deps.Logger.Error(err.Error())
		}
	}()
}

// Gracefully shutdown server
func (s *Server) Shutdown() {
	if s.server == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		s.Deps.Logger.Error(err.Error())
	}
}
