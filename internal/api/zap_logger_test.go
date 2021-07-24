package api_test

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mkantzer/emojiSorter/internal/api"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Zap Logger", func() {
	Context("GetLogger", func() {
		It("should return logger with request_id", func() {
			core, recorded := observer.New(zap.InfoLevel)
			parent := zap.New(core)

			r := httptest.NewRequest(http.MethodGet, "/", nil)
			r = r.WithContext(context.WithValue(r.Context(), middleware.RequestIDKey, "arequestid"))

			logger := api.GetLogger(parent, r)
			logger.Info("test")

			logs := recorded.All()
			Expect(len(logs)).To(Equal(1))
			Expect(logs[0].Message).To(Equal("test"))

			value, ok := logs[0].ContextMap()[api.RequestIdFieldKey]
			Expect(ok).To(Equal(true))
			Expect(value).To(Equal("arequestid"))
		})
	})

	Context("ZapLogger", func() {
		It("should log request and response", func() {
			core, recorded := observer.New(zap.InfoLevel)
			logger := zap.New(core)

			server := api.NewServer(&api.Dependencies{
				Logger: logger,
			}, "localhost:9001")

			handler := server.ZapLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			routeCtx := chi.NewRouteContext()
			routeCtx.RoutePatterns = []string{"/"}
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, routeCtx))
			r = r.WithContext(context.WithValue(r.Context(), middleware.RequestIDKey, "arequestid"))

			handler.ServeHTTP(w, r)

			logs := recorded.All()
			Expect(len(logs)).To(Equal(2))
			Expect(logs[0].Message).To(Equal("http request"))
			Expect(logs[0].ContextMap()).To(Equal(map[string]interface{}{
				"request_id": "arequestid",
				"path":       "/",
				"method":     "GET",
				"host":       "192.0.2.1:1234",
			}))

			Expect(logs[1].Level).To(Equal(zap.InfoLevel))
			Expect(logs[1].Message).To(Equal("http response"))
			Expect(logs[1].ContextMap()).To(Equal(map[string]interface{}{
				"path":          "/",
				"method":        "GET",
				"host":          "192.0.2.1:1234",
				"status":        int64(http.StatusOK),
				"response_time": int64(0),
				"path_template": "/",
				"request_id":    "arequestid",
			}))
		})

		It("should log error if internal server error", func() {
			core, recorded := observer.New(zap.InfoLevel)
			logger := zap.New(core)

			server := api.NewServer(&api.Dependencies{
				Logger: logger,
			}, "localhost:9001")

			handler := server.ZapLogger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}))

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)

			routeCtx := chi.NewRouteContext()
			routeCtx.RoutePatterns = []string{"/"}
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, routeCtx))
			r = r.WithContext(context.WithValue(r.Context(), middleware.RequestIDKey, "arequestid"))

			handler.ServeHTTP(w, r)

			logs := recorded.All()
			Expect(logs[1].Level).To(Equal(zap.ErrorLevel))
			Expect(logs[1].Message).To(Equal("http response"))
			Expect(logs[1].ContextMap()).To(Equal(map[string]interface{}{
				"path":          "/",
				"method":        "GET",
				"host":          "192.0.2.1:1234",
				"status":        int64(http.StatusInternalServerError),
				"response_time": int64(0),
				"path_template": "/",
				"request_id":    "arequestid",
			}))
		})
	})
})
