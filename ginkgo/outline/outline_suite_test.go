package outline_test

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"
)

func TestOutline(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Outline Suite")
}
