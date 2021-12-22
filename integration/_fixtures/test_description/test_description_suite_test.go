package test_description_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestTestDescription(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestDescription Suite")
}
