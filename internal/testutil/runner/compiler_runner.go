package runner

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)

var compilerBaselineRegex = regexp.MustCompile(`\.tsx?$`)

type CompilerTestType int

const (
	Conformance CompilerTestType = iota
	Regression
)

type CompilerBaselineRunner struct {
	testFiles []string
	basePath  string
	testType  CompilerTestType // >> TODO: will we need this?
}

var _ runner = (*CompilerBaselineRunner)(nil)

// >> TODO: make testType an enum
func NewCompilerBaselineRunner(testType CompilerTestType) *CompilerBaselineRunner {
	var basePath string
	if testType == Regression {
		basePath = "tests/cases/compiler"
	} else {
		basePath = "tests/cases/conformance"
	}
	return &CompilerBaselineRunner{
		basePath: basePath,
		testType: testType,
	}
}

func (r *CompilerBaselineRunner) EnumerateTestFiles() []string {
	if len(r.testFiles) > 0 {
		return r.testFiles
	}
	files, err := enumerateFiles(r.basePath, compilerBaselineRegex, true)
	if err != nil {
		panic("Could not read compiler test files: " + err.Error())
	}
	r.testFiles = files
	return files
}

func (r *CompilerBaselineRunner) RunTests(t *testing.T) {
	t.Parallel()
	files := r.EnumerateTestFiles()
	for _, filename := range files {
		// >> TODO: use name based on test type
		t.Run(fmt.Sprintf("Compiler tests for %s", filename), func(t *testing.T) {
			runTest(t, filename)
		})
	}
}

var compilerVaryBy []string

func runTest(t *testing.T, filename string) {
	t.Parallel()
	test := getCompilerFileBasedTest(filename)
	// >> TODO
}

func runSingleConfigTest()

// >> TODO: move functions below to utils or somewhere else
// >> TODO: should this be alias or newtype?
type fileBasedTestConfiguration = map[string]string

type compilerFileBasedTest struct {
	filename       string
	content        string
	configurations []fileBasedTestConfiguration
}

func getCompilerFileBasedTest(filename string) *compilerFileBasedTest {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic("Could not read test file: " + err.Error())
	}
	content := string(bytes)
	// settings := TestCaseParser.extractCompilerSettings(content) // !!!
	var settings map[string]string
	configurations := getFileBasedTestConfigurations(settings, compilerVaryBy) // !!!
	return &compilerFileBasedTest{
		filename:       filename,
		content:        content,
		configurations: configurations,
	}
}

func getFileBasedTestConfigurations(settings map[string]string, varyBy []string) []fileBasedTestConfiguration { // !!!
	// !!!
	var varyByEntries [][]string
	variationCount := 1
	for _, varyByKey := range varyBy {
		value, ok := settings[varyByKey]
		if ok {
			entries := splitVaryBySettingValue(value, varyByKey)

		}
	}
	// !!!
}

func splitVaryBySettingValue(value string, varyByKey string) []string {
	// !!!
}
