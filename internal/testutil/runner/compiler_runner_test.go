package runner

import (
	"testing"

	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/tspath"
	"gotest.tools/v3/assert"
)

func TestEnumerateTestFiles(t *testing.T) {
	t.Parallel()
	runner := NewCompilerBaselineRunner(Regression)
	testFiles := core.Map(runner.EnumerateTestFiles(), tspath.GetBaseFileName)
	assert.DeepEqual(t, testFiles, []string{"foo.ts"})
}

func TestRunCompilerTests(t *testing.T) {
	t.Parallel()
	runner := NewCompilerBaselineRunner(Regression)
	runner.RunTests(t)
}
