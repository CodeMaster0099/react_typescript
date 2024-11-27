package runner

import (
	"testing"
)

func TestCompilerBaselines(t *testing.T) {
	t.Parallel()
	runner := NewCompilerBaselineRunner(Regression)
	runner.RunTests(t)
}
