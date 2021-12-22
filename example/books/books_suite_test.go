package books_test

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}
