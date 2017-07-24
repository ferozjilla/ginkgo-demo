package dummy

import (
	. "github.com/onsi/ginkgo"
	"github.com/ferozjilla/ginkgo-demo/reporter"

	"testing"
)

var (
	dummyReporter *reporter.DummyReport
)

func TestDummyReport(t *testing.T) {
	dummyReporter = new(reporter.DummyReport)

	reporter := []Reporter{
		Reporter(dummyReporter),
	}

	RunSpecsWithDefaultAndCustomReporters(t, "Dummy Smoke Tests", reporter)
}
