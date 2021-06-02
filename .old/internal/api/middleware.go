package api

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// source: https://blog.questionable.services/article/guide-logging-middleware-go/

// This one may be better: https://pmihaylov.com/go-structured-logs/

// responseWriter is a minimal wrapper for http.ResponseWriter that allows the
// written HTTP status code to be captured for logging.
type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

type reqLog struct {
	Status   int
	Method   string
	Path     string
	Duration time.Duration
}

func (rl reqLog) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("status", rl.Status)
	enc.AddString("method", rl.Method)
	enc.AddString("endpoint", rl.Path)
	enc.AddDuration("duration", rl.Duration)
	return nil
}

// LoggingMiddleware logs the incoming HTTP request & its duration.
func LoggingMiddleware(logger *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error(err)
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			// Log request info, but only non-healthcheck things
			req := &reqLog{
				Status:   wrapped.status,
				Method:   r.Method,
				Path:     r.URL.EscapedPath(),
				Duration: time.Since(start),
			}
			if wrapped.status >= 400 {
				logger.Errorw("",
					zap.Object("req", req),
					zap.Error(fmt.Errorf("inbound request failed with status %d", req.Status)))
			} else if r.URL.EscapedPath() != "/healthz" {
				logger.Infow("inbound request succeeded",
					zap.Object("req", req))
			}

		}

		return http.HandlerFunc(fn)
	}
}
