package spec_iterator_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestSpecIterator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SpecIterator Suite")
}
