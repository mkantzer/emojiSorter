package initializer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mkantzer/emojiSorter/internal/initializer"
)

var _ = Describe("Logger", func() {
	It("should setup logger development logger", func() {
		_, err := initializer.Logging("development", "hostname", "service", "hash")
		Expect(err).To(BeNil())
	})
})
