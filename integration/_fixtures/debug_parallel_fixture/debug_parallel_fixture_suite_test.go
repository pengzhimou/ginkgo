package debug_parallel_fixture_test

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"
)

func TestDebugParallelFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DebugParallelFixture Suite")
}
