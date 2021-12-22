package specrunner_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestSpecRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spec Runner Suite")
}
