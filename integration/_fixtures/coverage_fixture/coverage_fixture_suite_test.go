package coverage_fixture_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestCoverageFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CoverageFixture Suite")
}
