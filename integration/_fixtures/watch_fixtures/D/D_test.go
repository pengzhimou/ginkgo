package D_test

import (
	. "$ROOT_PATH$/D"

	. "github.com/pengzhimou/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("D", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
