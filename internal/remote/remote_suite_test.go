package remote_test

import (
	. "github.com/onsi/gomega"
	. "github.com/pengzhimou/ginkgo"

	"testing"
)

func TestRemote(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Remote Spec Forwarding Suite")
}
