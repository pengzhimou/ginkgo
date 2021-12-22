package failer_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestFailer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Failer Suite")
}
