package tmp_test

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"
)

func TestConvertFixtures(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConvertFixtures Suite")
}
