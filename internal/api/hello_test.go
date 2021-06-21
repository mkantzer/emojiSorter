package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mkantzer/emojiSorter/internal/api"
)

var _ = Describe("Hello Endpoint", func() {
	Context("HelloServer", func() {
		It("returns OK w/ Hello World!", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()

			api.HelloServer(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))
			Expect(res.Body).ToNot(BeNil())

			body, _ := io.ReadAll(res.Body)
			Expect(string(body)).To(Equal("Hello World!\n"))
		})
	})
})
