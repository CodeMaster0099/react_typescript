package runner

import (
	"testing"
)

func TestCompilerBaselines(t *testing.T) {
	t.Parallel()
	testTypes := []CompilerTestType{Regression, Conformance}
	for _, testType := range testTypes {
		cleanUpLocalCompilerTests(testType)
		runner := NewCompilerBaselineRunner(testType)
		runner.RunTests(t)
	}
}
