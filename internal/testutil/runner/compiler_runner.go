package runner

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"strings"
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
	configurations := getFileBasedTestConfigurations(settings, compilerVaryBy)
	return &compilerFileBasedTest{
		filename:       filename,
		content:        content,
		configurations: configurations,
	}
}

func getFileBasedTestConfigurations(settings map[string]string, option []string) []fileBasedTestConfiguration {
	var optionEntries [][]string
	variationCount := 1
	for _, optionKey := range option {
		value, ok := settings[optionKey]
		if ok {
			entries := splitOptionValues(value, optionKey)
			if len(entries) > 0 {
				variationCount *= len(entries)
				if variationCount > 25 {
					panic(fmt.Sprintf("Provided test options exceeded the maximum number of variations: %s", strings.Join(option, ", ")))
				}
				optionEntries = append(optionEntries, []string{optionKey, value})
			}
		}
	}

	if len(optionEntries) == 0 {
		return nil
	}

	return computeFileBasedTestConfigurationVariations(variationCount, optionEntries)
}

func splitOptionValues(value string, option string) []string {
	if len(value) == 0 {
		return nil
	}

	star := false
	var includes []string
	var excludes []string
	for _, s := range strings.Split(value, ",") {
		s = strings.ToLower(strings.TrimSpace(s))
		if len(s) == 0 {
			continue
		}
		if s == "*" {
			star = true
		} else if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "!") {
			excludes = append(excludes, s[1:])
		} else {
			includes = append(includes, s)
		}
	}

	// do nothing if the setting has no variations
	if len(includes) <= 1 && !star && len(excludes) == 0 {
		return nil
	}

	// !!! We should dedupe the variations by their normalized values instead of by name
	var variations map[string]struct{}

	// add (and deduplicate) all included entries
	for _, include := range includes {
		// value := getValueOfSetting(setting, include)
		variations[include] = struct{}{}
	}

	allValues := getAllValuesForSetting(option)
	if star && len(allValues) > 0 {
		// add all entries
		for _, value := range allValues {
			variations[value] = struct{}{}
		}
	}

	// remove all excluded entries
	for _, exclude := range excludes {
		delete(variations, exclude)
	}

	if len(variations) == 0 {
		panic(fmt.Sprintf("Variations in test option '@%s' resulted in an empty set.", option))
	}

	return slices.Collect(maps.Keys(variations))
}

func getAllValuesForOption(option string) []string {
	// !!!
	return nil
}

// Leave setting values as strings, for now, then normalize later?
// however we'll need to deduplicate them, and there's a chance diff strings map to same setting value.
// Also need to validate the values?

// // >> TODO: this should not return a string, but let's pretend for now
// func getValueOfSetting(setting string, value string) string {}

func computeFileBasedTestConfigurationVariations(variationCount int, optionEntries [][]string) []fileBasedTestConfiguration {
	configurations := make([]fileBasedTestConfiguration, 0, variationCount)
	computeFileBasedTestConfigurationVariationsWorker(&configurations, optionEntries, 0, make(map[string]string))
	return configurations
}

func computeFileBasedTestConfigurationVariationsWorker(
	configurations *[]fileBasedTestConfiguration,
	optionEntries [][]string,
	index int,
	variationState fileBasedTestConfiguration) {
	if index >= len(optionEntries) {
		*configurations = append(*configurations, maps.Clone(variationState))
		return
	}

	optionKey := optionEntries[index][0]
	entries := optionEntries[index][1:]
	for _, entry := range entries {
		// set or overwrite the variation, then compute the next variation
		variationState[optionKey] = entry
		computeFileBasedTestConfigurationVariationsWorker(configurations, optionEntries, index+1, variationState)
	}
}
