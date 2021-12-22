package suite_command_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestSuiteCommand(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Suite Command Suite")
}
