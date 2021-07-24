package initializer_test

import (
	"github.com/mkantzer/emojiSorter/internal/initializer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logger", func() {
	It("should setup logger development logger", func() {
		_, err := initializer.Logging("development", "hostname", "service", "hash", nil)
		Expect(err).To(BeNil())
	})
})
