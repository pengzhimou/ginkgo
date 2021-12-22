package hanging_suite_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestHangingSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HangingSuite Suite")
}
