package progress_fixture_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestProgressFixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProgressFixture Suite")
}
