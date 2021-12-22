package eventually_failing_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestEventuallyFailing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EventuallyFailing Suite")
}
