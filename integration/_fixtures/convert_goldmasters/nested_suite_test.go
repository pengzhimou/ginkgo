package nested_test

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"
)

func TestNested(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nested Suite")
}
