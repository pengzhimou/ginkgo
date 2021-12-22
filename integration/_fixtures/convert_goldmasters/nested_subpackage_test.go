package subpackage

import (
	. "github.com/pengzhimou/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("nested sub packages", func() {
		GinkgoT().Fail(true)
	})
})
