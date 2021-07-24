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

var _ = Describe("Hello Endpoint", func() {
	server := api.Server{
		Deps: &api.Dependencies{
			Logger: zap.NewNop(),
		},
	}

	Context("HelloServer", func() {
		It("returns OK w/ Hello World!", func() {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			res := httptest.NewRecorder()

			server.HelloServer(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))
			Expect(res.Body).ToNot(BeNil())

			body, _ := io.ReadAll(res.Body)
			Expect(string(body)).To(Equal("Hello World!\n"))
		})
	})
})
