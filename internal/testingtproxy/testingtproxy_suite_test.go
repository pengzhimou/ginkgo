package testingtproxy_test

import (
	"testing"

	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"
)

func TestTestingtproxy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testingtproxy Suite")
}
