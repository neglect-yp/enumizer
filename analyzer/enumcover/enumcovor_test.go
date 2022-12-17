package enumcover_test

import (
	"github.com/neglect-yp/enumizer/analyzer/enumcover"
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestEnumcoverAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, enumcover.Analyzer, "a", "across_package")
}
