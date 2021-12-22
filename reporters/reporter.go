package reporters

import (
	"github.com/pengzhimou/ginkgo/config"
	"github.com/pengzhimou/ginkgo/types"
)

type Reporter interface {
	SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary)
	BeforeSuiteDidRun(setupSummary *types.SetupSummary)
	SpecWillRun(specSummary *types.SpecSummary)
	SpecDidComplete(specSummary *types.SpecSummary)
	AfterSuiteDidRun(setupSummary *types.SetupSummary)
	SpecSuiteDidEnd(summary *types.SuiteSummary)
}
