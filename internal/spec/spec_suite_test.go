package spec_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spec Suite")
}
