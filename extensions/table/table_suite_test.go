package table_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestTable(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Table Suite")
}
