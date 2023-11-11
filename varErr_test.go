package varErrChecker_test

import (
	"testing"

	"github.com/tomato3713/varErrChecker"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test_a(t *testing.T) {
	testdata := analysistest.TestData()
	t.Run("simple case", func(t *testing.T) {
		analysistest.Run(t, testdata, varErrChecker.Analyzer, "a")
	})

	t.Run("set pattern flag", func(t *testing.T) {
		varErrChecker.Analyzer.Flags.Set("pattern", `^[\d\w]+Error$`)
		analysistest.Run(t, testdata, varErrChecker.Analyzer, "b")
	})
}
