package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mkantzer/emojiSorter/internal/api"
)

var _ = Describe("Health Checks", func() {
	Context("HealthCheck", func() {
		It("returns OK", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()

			api.HealthCheck(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))
			Expect(res.Body).ToNot(BeNil())

			body, _ := io.ReadAll(res.Body)
			Expect(string(body)).To(Equal("This seems fine\n"))
		})
	})

	Context("HealthCheck", func() {
		It("returns OK", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()

			api.UnhealthCheck(res, req)
			Expect(res.Code).To(Equal(http.StatusInternalServerError))
			Expect(res.Body).ToNot(BeNil())

			body, _ := io.ReadAll(res.Body)
			Expect(string(body)).To(Equal("This seems Not Fine\n"))
		})
	})
})
