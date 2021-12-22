package testsuite_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestTestsuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testsuite Suite")
}
