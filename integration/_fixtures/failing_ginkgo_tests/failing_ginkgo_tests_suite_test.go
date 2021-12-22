package failing_ginkgo_tests_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestFailing_ginkgo_tests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Failing_ginkgo_tests Suite")
}
