package D_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "D Suite")
}
