package tmp

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestTmp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tmp Suite")
}
