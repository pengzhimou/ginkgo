package tags_tests_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestTagsTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TagsTests Suite")
}
