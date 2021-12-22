package containernode_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestContainernode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Containernode Suite")
}
