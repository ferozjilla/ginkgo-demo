package reporter

import (
	"fmt"

	"github.com/onsi/ginkgo/types"
	ginkgoConfig "github.com/onsi/ginkgo/config"
)

type DummyReport struct {
}

func (report *DummyReport) BeforeSuiteDidRun(summary *types.SetupSummary) {
	fmt.Println("DUMMY BeforeSuiteDidRun")
}

func (report *DummyReport) SpecWillRun(summary *types.SpecSummary) {
	fmt.Println("DUMMY SpecWillRun")
}

func (report *DummyReport) SpecDidComplete(summary *types.SpecSummary) {
	fmt.Println("DUMMY SpecDidComplete")
}

func (report *DummyReport) AfterSuiteDidRun(summary *types.SetupSummary) {
	fmt.Println("DUMMY AfterSuiteDidRun SUCCESS!")
}

func (report *DummyReport) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	fmt.Println("DUMMY SpecSuiteDidEnd")
}

func (report *DummyReport) SpecSuiteWillBegin(config ginkgoConfig.GinkgoConfigType, summary *types.SuiteSummary) {
	fmt.Println("DUMMY SpecSuiteWillBegin")
}
