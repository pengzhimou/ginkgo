package A_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestA(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "A Suite")
}
