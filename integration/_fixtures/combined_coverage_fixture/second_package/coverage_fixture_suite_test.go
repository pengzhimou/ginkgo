package second_package_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestCoverageFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CombinedFixture Second Suite")
}
