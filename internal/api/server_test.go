package api_test

import (
	"fmt"
	"net/http"

	"github.com/mkantzer/emojiSorter/internal/api"
	"go.uber.org/zap"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("API Server", func() {
	logger, _ := zap.NewDevelopment()
	deps := api.Dependencies{
		Logger: logger,
	}
	addr := "localhost:9001"

	It("should be able to start and stop server", func() {
		server := api.NewServer(&deps, addr)
		server.Start()

		resp, err := http.Get(fmt.Sprintf("http://%s/healthz", addr))
		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		server.Shutdown()

		_, err = http.Get(fmt.Sprintf("http://%s/healthz", addr))
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(ContainSubstring("connection refused"))
	})

})
