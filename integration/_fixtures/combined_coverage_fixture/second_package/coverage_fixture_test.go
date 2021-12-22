package second_package_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"
	. "github.com/pengzhimou/ginkgo/integration/_fixtures/combined_coverage_fixture/second_package"
)

var _ = Describe("CoverageFixture", func() {
	It("should test A", func() {
		Ω(A()).Should(Equal("A"))
	})

	It("should test B", func() {
		Ω(B()).Should(Equal("B"))
	})

	It("should test C", func() {
		Ω(C()).Should(Equal("C"))
	})

	It("should test D", func() {
		Ω(D()).Should(Equal("D"))
	})

	It("should test E", func() {
		Ω(E()).Should(Equal("E"))
	})
})
