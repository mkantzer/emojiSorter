package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

const RequestIdFieldKey = "request_id"

// Get request-scoped logger (includes request id and other details)
func GetLogger(parent *zap.Logger, r *http.Request) *zap.Logger {
	return parent.With(
		zap.String(RequestIdFieldKey, middleware.GetReqID(r.Context())),
	)
}

// Chi middleware that logs request and response
func (s *Server) ZapLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := GetLogger(s.Deps.Logger, r)

		fields := []zap.Field{
			zap.String("path", r.URL.Path),
			zap.String("method", r.Method),
			zap.String("host", r.RemoteAddr),
		}

		logger.Info("http request", fields...)

		startTime := time.Now()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		responseTime := time.Since(startTime)

		// response fields
		fields = append(fields,
			zap.Int("status", ww.Status()),
			zap.Int64("response_time", responseTime.Milliseconds()),
			zap.String("path_template", chi.RouteContext(r.Context()).RoutePattern()),
		)

		if ww.Status() >= 500 {
			logger.Error("http response", fields...)
		} else {
			logger.Info("http response", fields...)
		}
	})
}
