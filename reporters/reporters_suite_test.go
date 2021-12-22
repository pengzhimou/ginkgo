package reporters_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestReporters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reporters Suite")
}
