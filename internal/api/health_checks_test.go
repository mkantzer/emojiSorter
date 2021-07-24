package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/mkantzer/emojiSorter/internal/api"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var _ = Describe("Health Checks", func() {
	server := api.Server{
		Deps: &api.Dependencies{
			Logger: zap.NewNop(),
		},
	}

	Context("HealthCheck", func() {
		It("returns OK", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()

			server.HealthCheck(res, req)
			Expect(res.Result().StatusCode).To(Equal(http.StatusOK))
			Expect(res.Result().Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
			Expect(res.Result().Body).ToNot(BeNil())

			body, _ := io.ReadAll(res.Result().Body)
			Expect(string(body)).To(Equal("{\"status\":\"ok\"}\n"))
		})
	})

	Context("UnhealthCheck", func() {
		It("returns 500", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()

			server.UnhealthCheck(res, req)
			Expect(res.Result().StatusCode).To(Equal(http.StatusInternalServerError))
			Expect(res.Result().Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
			Expect(res.Result().Body).ToNot(BeNil())

			body, _ := io.ReadAll(res.Result().Body)
			Expect(string(body)).To(Equal("{\"status\":\"error\",\"reason\":\"This seems Not Fine\"}\n"))
		})
	})

	Context("Contract tests", func() {
		Context("/healthz", func() {
			It("returns OK", func() {
				res, err := http.Get("http://localhost:8080/healthz")
				Expect(err).To(BeNil())

				body, _ := io.ReadAll(res.Body)
				Expect(res.StatusCode).To(Equal(http.StatusOK))
				Expect(res.Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
				Expect(string(body)).To(Equal("{\"status\":\"ok\"}\n"))
			})
		})

		Context("/unhealthz", func() {
			It("returns OK", func() {
				res, err := http.Get("http://localhost:8080/unhealthz")
				Expect(err).To(BeNil())

				body, _ := io.ReadAll(res.Body)
				Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
				Expect(res.Header.Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
				Expect(string(body)).To(Equal("{\"status\":\"error\",\"reason\":\"This seems Not Fine\"}\n"))
			})
		})
	})
})
