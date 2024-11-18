package runner

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/microsoft/typescript-go/internal/compiler"
	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/testutil/baseline"
	"github.com/microsoft/typescript-go/internal/tspath"
)

var (
	compilerBaselineRegex = regexp.MustCompile(`\.tsx?$`)
	requireRegex          = regexp.MustCompile(`require\(`)
	referencesRegex       = regexp.MustCompile(`reference\spath`)
)

var srcFolder = "/.src" // !!! Move this to vfs or equivalent of `vfsUtils.ts`

type CompilerTestType int

const (
	Conformance CompilerTestType = iota
	Regression
)

type CompilerBaselineRunner struct {
	testFiles    []string
	basePath     string
	testSuitName string
}

var _ runner = (*CompilerBaselineRunner)(nil)

func NewCompilerBaselineRunner(testType CompilerTestType) *CompilerBaselineRunner {
	var testSuitName string
	if testType == Regression {
		testSuitName = "compiler"
	} else {
		testSuitName = "conformance"
	}
	basePath := fmt.Sprintf("tests/cases/%s", testSuitName)
	return &CompilerBaselineRunner{
		basePath:     basePath,
		testSuitName: testSuitName,
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
		r.runTest(t, filename)
	}
}

var compilerVaryBy []string

func (r *CompilerBaselineRunner) runTest(t *testing.T, filename string) {
	t.Parallel()
	test := getCompilerFileBasedTest(filename)
	if len(test.configurations) > 0 {
		for _, config := range test.configurations {
			description := "" // !!!
			t.Run(fmt.Sprint("%s tests for %s%s", r.testSuitName, filename, description), func(t *testing.T) { runSingleConfigTest(t, test, config) })
		}
	} else {
		t.Run(fmt.Sprint("%s tests for %s", r.testSuitName, filename), func(t *testing.T) { runSingleConfigTest(t, test, nil) })
	}
}

func runSingleConfigTest(t *testing.T, test *compilerFileBasedTest, config fileBasedTestConfiguration) {
	t.Parallel()
	payload := makeUnitsFromTest(test.content, test.filename)
	compilerTest := NewCompilerTest(test.filename, &payload, config)

	compilerTest.VerifyDiagnostics(t)
	// !!! Verify all baselines
}

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
	settings := extractCompilerSettings(content)
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

type compilerTest struct {
	filename       string
	basename       string
	configuredName string // name with configuration description, e.g. `file`
	options        core.CompilerOptions
	result         any // !!! compile files result
	tsConfigFiles  []*baseline.TestFile
	toBeCompiled   []*baseline.TestFile // equivalent to the files that will be passed on the command line
	otherFiles     []*baseline.TestFile // equivalent to other files on the file system not directly passed to the compiler (ie things that are referenced by other files)
	hasNonDtsFiles bool
}

type testCaseContentWithConfig struct {
	testCaseContent
	configuration fileBasedTestConfiguration
}

func NewCompilerTest(filename string, testContent *testCaseContent, configuration fileBasedTestConfiguration) *compilerTest {
	absoluteRootDir := srcFolder
	basename := tspath.GetBaseFileName(filename)
	configuredName := basename
	if configuration != nil {
		// Compute name with configuration description, e.g. `filename(target=esnext).ts`
		var configNameBuilder strings.Builder
		keys := slices.Sorted(maps.Keys(configuration))
		for i, key := range keys {
			if i > 0 {
				configNameBuilder.WriteRune(',')
			}
			fmt.Fprintf(&configNameBuilder, "%s=%s", strings.ToLower(key), strings.ToLower(configuration[key]))
		}
		configName := configNameBuilder.String()
		if len(configName) > 0 {
			extname := tspath.GetAnyExtensionFromPath(basename, nil, false)
			extensionlessBasename := basename[:len(basename)-len(extname)]
			configuredName = fmt.Sprintf("%s(%s)%s", extensionlessBasename, configName, extname)
		}
	}

	testCaseContentWithConfig := testCaseContentWithConfig{
		testCaseContent: *testContent,
		configuration:   configuration,
	}

	units := testCaseContentWithConfig.testUnitData
	var toBeCompiled []*baseline.TestFile
	var otherFiles []*baseline.TestFile
	var tsConfigOptions *core.CompilerOptions
	hasNonDtsFiles := core.Some(units, func(unit *testUnit) bool { return !tspath.FileExtensionIs(unit.name, compiler.ExtensionDts) })
	harnessConfig := testCaseContentWithConfig.configuration
	// var tsConfigFiles []*baseline.TestFile // !!!
	if testCaseContentWithConfig.tsConfig != nil {
		// !!!
	} else {
		baseUrl, ok := harnessConfig["baseUrl"]
		if ok && !tspath.IsRootedDiskPath(baseUrl) {
			harnessConfig["baseUrl"] = tspath.GetNormalizedAbsolutePath(baseUrl, absoluteRootDir)
		}

		lastUnit := units[len(units)-1]
		// We need to assemble the list of input files for the compiler and other related files on the 'filesystem' (ie in a multi-file test)
		// If the last file in a test uses require or a triple slash reference we'll assume all other files will be brought in via references,
		// otherwise, assume all files are just meant to be in the same compilation session without explicit references to one another.

		if testCaseContentWithConfig.configuration["noImplicitReferences"] != "" ||
			requireRegex.MatchString(lastUnit.content) ||
			referencesRegex.MatchString(lastUnit.content) {
			toBeCompiled = append(toBeCompiled, createHarnessTestFile(lastUnit))
			for _, unit := range units[:len(units)-1] {
				otherFiles = append(otherFiles, createHarnessTestFile(unit))
			}
		} else {
			toBeCompiled = core.Map(units, createHarnessTestFile)
		}
	}

	if tsConfigOptions != nil && tsConfigOptions.ConfigFilePath != "" {
		// tsConfigOptions.configFile!.fileName = tsConfigOptions.configFilePath; // !!!
	}

	result := compileFiles(
		toBeCompiled,
		otherFiles,
		harnessConfig,
		tsConfigOptions,
		harnessConfig["currentDirectory"],
		testCaseContentWithConfig.symlinks,
	)

	return &compilerTest{
		filename:       filename,
		basename:       basename,
		configuredName: configuredName,
		// options: result.options, // !!!
		result: result,
		// tsConfigFiles: tsConfigFiles, // !!!
		toBeCompiled:   toBeCompiled,
		otherFiles:     otherFiles,
		hasNonDtsFiles: hasNonDtsFiles,
	}
}

func (c *compilerTest) VerifyDiagnostics(t *testing.T) {
	// pretty := c.options.pretty
	pretty := false                                   // !!! Add `pretty` to compiler options
	baseline.DoErrorBaseline(t, "", nil, nil, pretty) // !!!
}

func createHarnessTestFile(unit *testUnit) *baseline.TestFile {
	return &baseline.TestFile{
		UnitName:    unit.name,
		Content:     unit.content,
		FileOptions: unit.fileOptions,
	}
}
