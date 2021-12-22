package example_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Example Suite")
}
