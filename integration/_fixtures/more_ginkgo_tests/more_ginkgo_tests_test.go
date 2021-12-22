package more_ginkgo_tests_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"
	. "github.com/pengzhimou/ginkgo/integration/_fixtures/more_ginkgo_tests"
)

var _ = Describe("MoreGinkgoTests", func() {
	It("should pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})

	It("should always pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})
})
