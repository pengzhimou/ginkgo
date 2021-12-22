package C_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "C Suite")
}
