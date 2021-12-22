package focused_fixture_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestFocused_fixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Focused_fixture Suite")
}
