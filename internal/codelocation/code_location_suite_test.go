package codelocation_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestCodelocation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CodeLocation Suite")
}
