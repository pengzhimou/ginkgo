package C_test

import (
	. "$ROOT_PATH$/C"

	. "github.com/pengzhimou/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("C", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
