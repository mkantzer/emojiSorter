package initializer_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mkantzer/emojiSorter/internal/initializer"
)

var _ = Describe("ApiServer", func() {
	logger, err := initializer.Logging("development", "hostname", "service", "hash")
	Expect(err).To(BeNil())

	It("should setup api server on 0.0.0.0:8080", func() {
		server, err := initializer.ApiServer(logger)
		Expect(err).To(BeNil())
		Expect(server.Addr).To(Equal("0.0.0.0:8080"))
	})

	It("should allow setting port via env var", func() {
		prevPort := os.Getenv("PORT")
		os.Setenv("PORT", "9001")
		defer os.Setenv("PORT", prevPort)

		server, err := initializer.ApiServer(logger)
		Expect(err).To(BeNil())
		Expect(server.Addr).To(Equal("0.0.0.0:9001"))
	})

	It("should error if PORT env var not valid", func() {
		prevPort := os.Getenv("PORT")
		os.Setenv("PORT", "notanumber")
		defer os.Setenv("PORT", prevPort)

		_, err := initializer.ApiServer(logger)
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("problem getting port: strconv.Atoi: parsing \"notanumber\": invalid syntax"))
	})
})
