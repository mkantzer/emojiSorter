package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mkantzer/emojiSorter/internal/api"
)

var _ = Describe("Health Checks", func() {
	Context("HealthCheck", func() {
		It("returns OK", func() {
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

			api.HealthCheck(c)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body).ToNot(BeNil())

			body, _ := io.ReadAll(w.Body)
			Expect(string(body)).To(Equal("This seems fine\n"))
		})
	})

	Context("HealthCheck", func() {
		It("returns Not OK", func() {
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

			api.UnhealthCheck(c)

			Expect(w.Code).To(Equal(http.StatusInternalServerError))
			Expect(w.Body).ToNot(BeNil())

			body, _ := io.ReadAll(w.Body)
			Expect(string(body)).To(Equal("This seems Not Fine\n"))
		})
	})
})

var _ = Describe("Hello Endpoint", func() {
	Context("HelloServer", func() {
		It("returns OK w/ Hello World!", func() {
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

			api.HelloServer(c)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body).ToNot(BeNil())

			body, _ := io.ReadAll(w.Body)
			Expect(string(body)).To(Equal("Hello World!\n"))
		})
	})
})