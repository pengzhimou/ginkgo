package B_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "B Suite")
}
