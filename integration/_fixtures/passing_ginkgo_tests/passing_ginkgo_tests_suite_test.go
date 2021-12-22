package passing_ginkgo_tests_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestPassing_ginkgo_tests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Passing_ginkgo_tests Suite")
}
