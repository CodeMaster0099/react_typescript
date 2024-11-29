package runner

import (
	"testing"
)

func TestCompilerBaselines(t *testing.T) {
	t.Parallel()
	testTypes := []CompilerTestType{TestTypeRegression, TestTypeConformance}
	for _, testType := range testTypes {
		t.Run(testType.String(), func(t *testing.T) {
			t.Parallel()
			cleanUpLocalCompilerTests(testType)
			runner := NewCompilerBaselineRunner(testType)
			runner.RunTests(t)
		})
	}
}
