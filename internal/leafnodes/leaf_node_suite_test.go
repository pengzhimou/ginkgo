package leafnodes_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestLeafNode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LeafNode Suite")
}
