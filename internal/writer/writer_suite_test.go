package writer_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestWriter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Writer Suite")
}
