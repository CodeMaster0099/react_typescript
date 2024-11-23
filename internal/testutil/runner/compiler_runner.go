package runner

import (
	"fmt"
	"maps"
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/microsoft/typescript-go/internal/ast"
	"github.com/microsoft/typescript-go/internal/compiler"
	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/testutil/baseline"
	"github.com/microsoft/typescript-go/internal/tspath"
	"github.com/microsoft/typescript-go/internal/vfs"
)

var (
	compilerBaselineRegex = regexp.MustCompile(`\.tsx?$`)
	requireRegex          = regexp.MustCompile(`require\(`)
	referencesRegex       = regexp.MustCompile(`reference\spath`)
)

var (
	// Posix-style path to sources under test
	srcFolder = "/.src" // !!! Move this to vfs or equivalent of `vfsUtils.ts`
	// Posix-style path to the TypeScript compiler build outputs (including tsc.js, lib.d.ts, etc.)
	builtFolder = "/.ts"
	// Posix-style path to additional test libraries
	testLibFolder = "/.lib"
)

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
	// !!! Basepath has to account for the fact that the cwd is the package root when the tests are run
	basePath := fmt.Sprintf("tests/cases/%s", testSuitName)
	return &CompilerBaselineRunner{
		basePath:     basePath,
		testSuitName: testSuitName,
	}
}

// >> TODO: right now, this returns absolute, normalized file paths, but maybe we want paths relative to the cwd i.e. test folder
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
	files := r.EnumerateTestFiles()
	for _, filename := range files {
		r.runTest(t, filename)
	}
}

var compilerVaryBy []string

func (r *CompilerBaselineRunner) runTest(t *testing.T, filename string) {
	test := getCompilerFileBasedTest(filename)
	if len(test.configurations) > 0 {
		for _, config := range test.configurations {
			description := "" // !!!
			t.Run(fmt.Sprintf("%s tests for %s%s", r.testSuitName, filename, description), func(t *testing.T) { runSingleConfigTest(t, test, config) })
		}
	} else {
		// !!! Fix filename for printing
		t.Run(fmt.Sprintf("%s tests for %s", r.testSuitName, filename), func(t *testing.T) { runSingleConfigTest(t, test, nil) })
	}
}

func runSingleConfigTest(t *testing.T, test *compilerFileBasedTest, config fileBasedTestConfiguration) {
	t.Parallel()
	payload := makeUnitsFromTest(test.content, test.filename)
	compilerTest := NewCompilerTest(test.filename, &payload, config)

	compilerTest.VerifyDiagnostics(t)
	// !!! Verify all baselines
}

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

	allValues := getAllValuesForOption(option)
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

// >> End of utils/common stuff to move

type compilerTest struct {
	filename       string
	basename       string
	configuredName string // name with configuration description, e.g. `file`
	options        *core.CompilerOptions
	result         *CompileFilesResult
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
	hasNonDtsFiles := core.Some(units, func(unit *testUnit) bool { return !tspath.FileExtensionIs(unit.name, tspath.ExtensionDts) })
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
	pretty := false // !!! Add `pretty` to compiler options
	files := core.Concatenate(c.tsConfigFiles, core.Concatenate(c.toBeCompiled, c.otherFiles))
	baseline.DoErrorBaseline(t, c.configuredName, files, c.result.Diagnostics, pretty)
}

func createHarnessTestFile(unit *testUnit) *baseline.TestFile {
	return &baseline.TestFile{
		UnitName:    unit.name,
		Content:     unit.content,
		FileOptions: unit.fileOptions,
	}
}

// >> Below: move

type harnessOptions struct {
	useCaseSensitiveFileNames bool
	includeBuiltFile          string
	baselineFile              string
	libFiles                  []string
	noTypesAndSymbols         bool
	captureSuggestions        bool
}

// >> TODO: also move this to common harness

type CompileFilesResult struct {
	Diagnostics []*ast.Diagnostic
}

func compileFiles(
	inputFiles []*baseline.TestFile,
	otherFiles []*baseline.TestFile,
	rawHarnessConfig fileBasedTestConfiguration,
	compilerOptions *core.CompilerOptions,
	currentDirectory string,
	symlinks any) *CompileFilesResult {
	// originalCurrentDirectory := currentDirectory
	var options core.CompilerOptions
	if compilerOptions != nil {
		options = *compilerOptions
	}
	harnessOptions := getHarnessOptions(rawHarnessConfig)

	if currentDirectory == "" {
		currentDirectory = srcFolder
	}

	var typescriptVersion string

	// Parse settings
	if rawHarnessConfig != nil { // >> TODO: why do we need this if we've already parsed ts config options in `NewCompilerTest`?
		setCompilerOptionsFromHarnessConfig(rawHarnessConfig, &options) // >> TODO: validate that options are either harness or compiler options
		typescriptVersion = rawHarnessConfig["typescriptVersion"]
	}

	// useCaseSensitiveFileNames := harnessOptions.useCaseSensitiveFileNames // >> TODO: this is wrong because this should default to true
	var programFileNames []string
	for _, file := range inputFiles {
		// >> TODO: why aren't they always normalized and absolute?
		// When a tsconfig is present, root names passed to createProgram should already be absolute
		var fileName string
		// if compilerOptions.ConfigFilePath != "" {
		// 	fileName = tspath.GetNormalizedAbsolutePath(file.UnitName, currentDirectory)
		// } else {
		// 	fileName = file.UnitName
		// }
		fileName = tspath.GetNormalizedAbsolutePath(file.UnitName, currentDirectory)

		if !tspath.FileExtensionIs(fileName, tspath.ExtensionJson) {
			programFileNames = append(programFileNames, fileName)
		}
	}

	// !!! Note: lib files are not going to be in `built/local`.
	// In addition, not all files that used to be in `built/local` are going to exist.
	// Files from built\local that are requested by test "@includeBuiltFiles" to be in the context.
	// Treat them as library files, so include them in build, but not in baselines.
	// if harnessOptions.includeBuiltFile != "" {
	// 	programFileNames = append(programFileNames, tspath.CombinePaths(builtFolder, harnessOptions.includeBuiltFile))
	// }

	// !!! This won't work until we have the actual lib files
	// // Files from tests\lib that are requested by "@libFiles"
	// if len(harnessOptions.libFiles) > 0 {
	// 	for _, libFile := range harnessOptions.libFiles {
	// 		programFileNames = append(programFileNames, tspath.CombinePaths(testLibFolder, libFile))
	// 	}
	// }

	// !!!
	// docs := append(inputFiles, otherFiles...) // !!! Convert to `TextDocument`
	// const fs = vfs.createFromFileSystem(IO, !useCaseSensitiveFileNames, { documents: docs, cwd: currentDirectory });
	// if (symlinks) {
	// 	fs.apply(symlinks);
	// }

	// ts.assign(options, ts.convertToOptionsWithAbsolutePaths(options, path => ts.getNormalizedAbsolutePath(path, currentDirectory)));

	// !!! Port vfs usage closer to original

	// Create fake FS for testing
	testfs := fstest.MapFS{}
	for _, file := range inputFiles {
		fileName := tspath.GetNormalizedAbsolutePath(file.UnitName, currentDirectory)
		testfs[fileName] = &fstest.MapFile{
			Data: []byte(file.Content),
		}
	}
	for _, file := range otherFiles {
		fileName := tspath.GetNormalizedAbsolutePath(file.UnitName, currentDirectory)
		testfs[fileName] = &fstest.MapFile{
			Data: []byte(file.Content),
		}
	}
	fs := vfs.FromIOFS(true, testfs)

	host := createCompilerHost(fs, &options, currentDirectory)
	result := compileFilesWithHost(host, programFileNames, &options, typescriptVersion, harnessOptions.captureSuggestions)

	return result
}

func getHarnessOptions(harnessConfig fileBasedTestConfiguration) harnessOptions {
	// !!!
	// >> TODO: split and trim libFiles by comma
	return harnessOptions{}
}

func setCompilerOptionsFromHarnessConfig(harnessConfig fileBasedTestConfiguration, options *core.CompilerOptions) {
	for name, value := range harnessConfig {
		if value == "" {
			panic(fmt.Sprintf("Cannot have undefined value for compiler option '%s'", name))
		}
		if name == "typescriptversion" {
			continue
		}

		// !!! Implement this once command line parsing is done
		// const option = getCommandLineOption(name);
		// if (option) {
		// 	const errors: ts.Diagnostic[] = [];
		// 	options[option.name] = optionValue(option, value, errors);
		// 	if (errors.length > 0) {
		// 		throw new Error(`Unknown value '${value}' for compiler option '${name}'.`);
		// 	}
		// }
		// else {
		// 	throw new Error(`Unknown compiler option '${name}'.`);
		// }
	}
}

func createCompilerHost(fs vfs.FS, options *core.CompilerOptions, currentDirectory string) compiler.CompilerHost {
	return compiler.NewCompilerHost(options, false, currentDirectory, fs)
}

func compileFilesWithHost(
	host compiler.CompilerHost,
	rootFiles []string,
	options *core.CompilerOptions,
	typescriptVersion string,
	captureSuggestions bool) *CompileFilesResult {
	// !!!
	// if (compilerOptions.project || !rootFiles || rootFiles.length === 0) {
	// 	const project = readProject(host.parseConfigHost, compilerOptions.project, compilerOptions);
	// 	if (project) {
	// 		if (project.errors && project.errors.length > 0) {
	// 			return new CompilationResult(host, compilerOptions, /*program*/ undefined, /*result*/ undefined, project.errors);
	// 		}
	// 		if (project.config) {
	// 			rootFiles = project.config.fileNames;
	// 			compilerOptions = project.config.options;
	// 		}
	// 	}
	// 	delete compilerOptions.project;
	// }

	// establish defaults (aligns with old harness)
	if options.NewLine == core.NewLineKindNone {
		options.NewLine = core.NewLineKindCRLF
	}
	// !!!
	// if options.SkipDefaultLibCheck == core.TSUnknown {
	// 	options.SkipDefaultLibCheck = core.TSTrue
	// }
	if options.NoErrorTruncation == core.TSUnknown {
		options.NoErrorTruncation = core.TSTrue
	}

	// pre-emit/post-emit error comparison requires declaration emit twice, which can be slow. If it's unlikely to flag any error consistency issues
	// and if the test is running `skipLibCheck` - an indicator that we want the tets to run quickly - skip the before/after error comparison, too
	// skipErrorComparison := len(rootFiles) >= 100 || options.SkipLibCheck == core.TSTrue && options.Declaration == core.TSTrue
	// var preProgram *compiler.Program
	// if !skipErrorComparison {
	// !!! Need actual program for this
	// preProgram = ts.createProgram({ rootNames: rootFiles || [], options: { ...compilerOptions, configFile: compilerOptions.configFile, traceResolution: false }, host, typeScriptVersion })
	// }
	// let preErrors = preProgram && ts.getPreEmitDiagnostics(preProgram);
	// if (preProgram && captureSuggestions) {
	//     preErrors = ts.concatenate(preErrors, ts.flatMap(preProgram.getSourceFiles(), f => preProgram.getSuggestionDiagnostics(f)));
	// }

	// const program = ts.createProgram({ rootNames: rootFiles || [], options: compilerOptions, host, typeScriptVersion });
	// const emitResult = program.emit();
	// let postErrors = ts.getPreEmitDiagnostics(program);
	// if (captureSuggestions) {
	//     postErrors = ts.concatenate(postErrors, ts.flatMap(program.getSourceFiles(), f => program.getSuggestionDiagnostics(f)));
	// }
	// const longerErrors = ts.length(preErrors) > postErrors.length ? preErrors : postErrors;
	// const shorterErrors = longerErrors === preErrors ? postErrors : preErrors;
	// const errors = preErrors && (preErrors.length !== postErrors.length) ? [
	//     ...shorterErrors!,
	//     ts.addRelatedInfo(
	//         ts.createCompilerDiagnostic({
	//             category: ts.DiagnosticCategory.Error,
	//             code: -1,
	//             key: "-1",
	//             message: `Pre-emit (${preErrors.length}) and post-emit (${postErrors.length}) diagnostic counts do not match! This can indicate that a semantic _error_ was added by the emit resolver - such an error may not be reflected on the command line or in the editor, but may be captured in a baseline here!`,
	//         }),
	//         ts.createCompilerDiagnostic({
	//             category: ts.DiagnosticCategory.Error,
	//             code: -1,
	//             key: "-1",
	//             message: `The excess diagnostics are:`,
	//         }),
	//         ...ts.filter(longerErrors!, p => !ts.some(shorterErrors, p2 => ts.compareDiagnostics(p, p2) === ts.Comparison.EqualTo)),
	//     ),
	// ] : postErrors;
	program := createProgram(host, options)
	var diagnostics []*ast.Diagnostic
	diagnostics = append(diagnostics, program.GetSyntacticDiagnostics(nil)...)
	diagnostics = append(diagnostics, program.GetBindDiagnostics(nil)...)
	diagnostics = append(diagnostics, program.GetSemanticDiagnostics(nil)...)
	diagnostics = append(diagnostics, program.GetGlobalDiagnostics()...)
	return &CompileFilesResult{
		Diagnostics: diagnostics,
	}
}

// !!! Temporary while we don't have the real `createProgram`
func createProgram(host compiler.CompilerHost, options *core.CompilerOptions) *compiler.Program {
	programOptions := compiler.ProgramOptions{
		Host:           host,
		Options:        options,
		SingleThreaded: true,
		RootPath:       host.GetCurrentDirectory(),
	}
	program := compiler.NewProgram(programOptions)
	return program
}
