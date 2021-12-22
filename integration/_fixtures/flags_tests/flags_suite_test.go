package flags_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestFlags(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flags Suite")
}
