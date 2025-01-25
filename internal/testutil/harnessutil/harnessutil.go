package harnessutil

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/microsoft/typescript-go/internal/ast"
	"github.com/microsoft/typescript-go/internal/bundled"
	"github.com/microsoft/typescript-go/internal/compiler"
	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/repo"
	"github.com/microsoft/typescript-go/internal/tsoptions"
	"github.com/microsoft/typescript-go/internal/tspath"
	"github.com/microsoft/typescript-go/internal/vfs"
	"github.com/microsoft/typescript-go/internal/vfs/vfstest"
)

type TestFile struct {
	UnitName    string
	Content     string
	FileOptions map[string]string
}

type CompilationResult struct {
	Diagnostics []*ast.Diagnostic
	Program     *compiler.Program
	Options     *core.CompilerOptions
	// !!! outputs
}

// This maps a compiler setting to its string value, after splitting by commas,
// handling inclusions and exclusions, and deduplicating.
// For example, if a test file contains:
//
//	// @target: esnext, es2015
//
// Then the map will map "target" to "esnext", and another map will map "target" to "es2015".
type TestConfiguration = map[string]string

type harnessOptions struct {
	allowNonTsExtensions      bool
	useCaseSensitiveFileNames bool
	baselineFile              string
	includeBuiltFile          string
	fileName                  string
	libFiles                  []string
	noErrorTruncation         bool
	suppressOutputPathCheck   bool
	noImplicitReferences      bool
	currentDirectory          string
	symlink                   string
	link                      string
	noTypesAndSymbols         bool
	fullEmitPaths             bool
	noCheck                   bool
	reportDiagnostics         bool
	captureSuggestions        bool
	typescriptVersion         string
}

func CompileFiles(
	t *testing.T,
	inputFiles []*TestFile,
	otherFiles []*TestFile,
	harnessConfig TestConfiguration,
	compilerOptions *core.CompilerOptions,
	currentDirectory string,
	symlinks any,
) *CompilationResult {
	var options core.CompilerOptions
	if compilerOptions != nil {
		options = *compilerOptions
	}
	// Set default options for tests
	if options.NewLine == core.NewLineKindNone {
		options.NewLine = core.NewLineKindCRLF
	}
	if options.SkipDefaultLibCheck == core.TSUnknown {
		options.SkipDefaultLibCheck = core.TSTrue
	}
	options.NoErrorTruncation = core.TSTrue
	harnessOptions := harnessOptions{useCaseSensitiveFileNames: true, currentDirectory: currentDirectory}

	// Parse harness and compiler options from the harness configuration
	if harnessConfig != nil {
		setOptionsFromHarnessConfig(t, harnessConfig, &options, &harnessOptions)
	}

	var programFileNames []string
	for _, file := range inputFiles {
		fileName := tspath.GetNormalizedAbsolutePath(file.UnitName, currentDirectory)

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
	// Note: the code below assumes a single root, since an FS in Go always has a single root.
	testfs := fstest.MapFS{}
	for _, file := range inputFiles {
		fileName := tspath.GetNormalizedAbsolutePath(file.UnitName, currentDirectory)
		rootLen := tspath.GetRootLength(fileName)
		fileName = fileName[rootLen:]
		testfs[fileName] = &fstest.MapFile{
			Data: []byte(file.Content),
		}
	}
	for _, file := range otherFiles {
		fileName := tspath.GetNormalizedAbsolutePath(file.UnitName, currentDirectory)
		rootLen := tspath.GetRootLength(fileName)
		fileName = fileName[rootLen:]
		testfs[fileName] = &fstest.MapFile{
			Data: []byte(file.Content),
		}
	}

	fs := vfstest.FromMapFS(testfs, harnessOptions.useCaseSensitiveFileNames)
	fs = bundled.WrapFS(fs)

	host := createCompilerHost(fs, &options, currentDirectory)
	result := compileFilesWithHost(host, programFileNames, &options, harnessOptions.typescriptVersion, harnessOptions.captureSuggestions)

	return result
}

func setOptionsFromHarnessConfig(t *testing.T, harnessConfig TestConfiguration, compilerOptions *core.CompilerOptions, harnessOptions *harnessOptions) {
	for name, value := range harnessConfig {
		if name == "typescriptversion" {
			continue
		}

		commandLineOption := getCommandLineOption(name)
		if commandLineOption != nil {
			parsedValue := optionValue(t, commandLineOption, value)
			tsoptions.ParseCompilerOptions(commandLineOption.Name, parsedValue, compilerOptions)
			continue
		}
		harnessOption := getHarnessOption(name)
		if harnessOption != nil {
			parsedValue := optionValue(t, harnessOption, value)
			parseHarnessOption(name, parsedValue, harnessOptions)
			continue
		}

		t.Fatalf("Unknown compiler option '%s'.", name)
	}
}

var harnessCommandLineOptions = []*tsoptions.CommandLineOption{
	{
		Name: "allowNonTsExtensions",
		Kind: tsoptions.CommandLineOptionTypeBoolean,
		// defaultValueDescription: false,
	},
	{
		Name: "useCaseSensitiveFileNames",
		Kind: "boolean",
		// defaultValueDescription: false,
	},
	{
		Name: "baselineFile",
		Kind: "string",
	},
	{
		Name: "includeBuiltFile",
		Kind: "string",
	},
	{
		Name: "fileName",
		Kind: "string",
	},
	{
		Name: "libFiles",
		Kind: "string",
	},
	{
		Name: "noErrorTruncation",
		Kind: "boolean",
		// defaultValueDescription: false,
	},
	{
		Name: "suppressOutputPathCheck",
		Kind: "boolean",
		// defaultValueDescription: false,
	},
	{
		Name: "noImplicitReferences",
		Kind: "boolean",
		// defaultValueDescription: false,
	},
	{
		Name: "currentDirectory",
		Kind: "string",
	},
	{
		Name: "symlink",
		Kind: "string",
	},
	{
		Name: "link",
		Kind: "string",
	},
	{
		Name: "noKindsAndSymbols",
		Kind: "boolean",
		// defaultValueDescription: false,
	},
	// Emitted js baseline will print full paths for every output file
	{
		Name: "fullEmitPaths",
		Kind: "boolean",
		// defaultValueDescription: false,
	},
	{
		Name: "noCheck",
		Kind: "boolean",
		// defaultValueDescription: false,
	},
	// used to enable error collection in `transpile` baselines
	{
		Name: "reportDiagnostics",
		Kind: "boolean",
		// defaultValueDescription: false,
	},
	// Adds suggestion diagnostics to error baselines
	{
		Name: "captureSuggestions",
		Kind: "boolean",
		// defaultValueDescription: false,
	},
}

func getHarnessOption(name string) *tsoptions.CommandLineOption {
	return core.Find(harnessCommandLineOptions, func(option *tsoptions.CommandLineOption) bool {
		return strings.ToLower(option.Name) == strings.ToLower(name)
	})
}

func parseHarnessOption(key string, value any, options *harnessOptions) {
	switch key {
	case "allowNonTsExtensions":
		options.allowNonTsExtensions = value.(bool)
	case "useCaseSensitiveFileNames":
		options.useCaseSensitiveFileNames = value.(bool)
	case "baselineFile":
		options.baselineFile = value.(string)
	case "includeBuiltFile":
		options.includeBuiltFile = value.(string)
	case "fileName":
		options.fileName = value.(string)
	case "libFiles":
		options.libFiles = value.([]string)
	case "noErrorTruncation":
		options.noErrorTruncation = value.(bool)
	case "suppressOutputPathCheck":
		options.suppressOutputPathCheck = value.(bool)
	case "noImplicitReferences":
		options.noImplicitReferences = value.(bool)
	case "currentDirectory":
		options.currentDirectory = value.(string)
	case "symlink":
		options.symlink = value.(string)
	case "link":
		options.link = value.(string)
	case "noTypesAndSymbols":
		options.noTypesAndSymbols = value.(bool)
	case "fullEmitPaths":
		options.fullEmitPaths = value.(bool)
	case "noCheck":
		options.noCheck = value.(bool)
	case "reportDiagnostics":
		options.reportDiagnostics = value.(bool)
	case "captureSuggestions":
		options.captureSuggestions = value.(bool)
	case "typescriptVersion":
		options.typescriptVersion = value.(string)
	}
}

func optionValue(t *testing.T, option *tsoptions.CommandLineOption, value string) tsoptions.CompilerOptionsValue {
	switch option.Kind {
	case tsoptions.CommandLineOptionTypeString:
		return value
	case tsoptions.CommandLineOptionTypeNumber:
		numVal, err := strconv.Atoi(value)
		if err != nil {
			t.Fatalf("Value for option '%s' must be a number, got: %v", option.Name, value)
		}
		return numVal
	case tsoptions.CommandLineOptionTypeBoolean:
		boolVal, err := strconv.ParseBool(value) // >> TODO: convert to Tristate?
		if err != nil {
			t.Fatalf("Value for option '%s' must be a boolean, got: %v", option.Name, value)
		}
		return boolVal
	case tsoptions.CommandLineOptionTypeEnum:
		enumVal, ok := option.EnumMap().Get(strings.ToLower(value))
		if !ok {
			t.Fatalf("Value for option '%s' must be one of %s, got: %v", option.Name, strings.Join(slices.Collect(option.EnumMap().Keys()), ","), value)
		}
		return enumVal
	case tsoptions.CommandLineOptionTypeList, tsoptions.CommandLineOptionTypeListOrElement:
		listVal, errors := tsoptions.ParseListTypeOption(option, value)
		if len(errors) > 0 {
			t.Fatalf("Unknown value '%s' for compiler option '%s'", value, option.Name)
		}
		return listVal
	case tsoptions.CommandLineOptionTypeObject:
		t.Fatalf("Object type options like '%s' are not supported", option.Name)
	}
	return nil
}

func createCompilerHost(fs vfs.FS, options *core.CompilerOptions, currentDirectory string) compiler.CompilerHost {
	return compiler.NewCompilerHost(options, currentDirectory, fs)
}

func compileFilesWithHost(
	host compiler.CompilerHost,
	rootFiles []string,
	options *core.CompilerOptions,
	typescriptVersion string,
	captureSuggestions bool,
) *CompilationResult {
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
	program := createProgram(host, options, rootFiles)
	var diagnostics []*ast.Diagnostic
	diagnostics = append(diagnostics, program.GetSyntacticDiagnostics(nil)...)
	diagnostics = append(diagnostics, program.GetBindDiagnostics(nil)...)
	diagnostics = append(diagnostics, program.GetSemanticDiagnostics(nil)...)
	diagnostics = append(diagnostics, program.GetGlobalDiagnostics()...)

	return newCompilationResult(options, program, diagnostics)
}

func newCompilationResult(
	options *core.CompilerOptions,
	program *compiler.Program,
	diagnostics []*ast.Diagnostic,
) *CompilationResult {
	if program != nil {
		options = program.Options()
	}
	return &CompilationResult{
		Diagnostics: diagnostics,
		Program:     program,
		Options:     options,
	}
}

// !!! Temporary while we don't have the real `createProgram`
func createProgram(host compiler.CompilerHost, options *core.CompilerOptions, rootFiles []string) *compiler.Program {
	programOptions := compiler.ProgramOptions{
		RootFiles:          rootFiles,
		Host:               host,
		Options:            options,
		DefaultLibraryPath: bundled.LibPath(),
	}
	program := compiler.NewProgram(programOptions)
	return program
}

func EnumerateFiles(folder string, testRegex *regexp.Regexp, recursive bool) ([]string, error) {
	files, err := listFiles(folder, testRegex, recursive)
	if err != nil {
		return nil, err
	}
	return core.Map(files, tspath.NormalizeSlashes), nil
}

func listFiles(path string, spec *regexp.Regexp, recursive bool) ([]string, error) {
	return listFilesWorker(spec, recursive, path)
}

func listFilesWorker(spec *regexp.Regexp, recursive bool, folder string) ([]string, error) {
	folder = tspath.GetNormalizedAbsolutePath(folder, repo.TestDataPath)
	entries, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	var paths []string
	for _, entry := range entries {
		path := filepath.Join(folder, entry.Name())
		if !entry.IsDir() {
			if spec == nil || spec.MatchString(path) {
				paths = append(paths, path)
			}
		} else if recursive {
			subPaths, err := listFilesWorker(spec, recursive, path)
			if err != nil {
				return nil, err
			}
			paths = append(paths, subPaths...)
		}
	}
	return paths, nil
}

func GetFileBasedTestConfigurationDescription(config TestConfiguration) string {
	var output strings.Builder
	keys := slices.Sorted(maps.Keys(config))
	for i, key := range keys {
		if i > 0 {
			output.WriteString(", ")
		}
		fmt.Fprintf(&output, "@%s: %s", key, config[key])
	}
	return output.String()
}

func GetFileBasedTestConfigurations(settings map[string]string, option []string) []TestConfiguration {
	var optionEntries [][]string
	variationCount := 1
	for _, optionKey := range option {
		value, ok := settings[optionKey]
		if ok {
			entries := splitOptionValues(value, optionKey)
			if len(entries) > 0 {
				variationCount *= len(entries)
				if variationCount > 25 {
					panic("Provided test options exceeded the maximum number of variations: " + strings.Join(option, ", "))
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

// Splits a string value into an array of strings, each corresponding to a unique value for the given option.
// Also handles the `*` value, which includes all possible values for the option, and exclusions.
// ```
//
//	splitOptionValues("esnext, es2015, es6", "target") => ["esnext", "es2015"]
//	splitOptionValues("*", "strict") => ["true", "false"]
//
// ```
func splitOptionValues(value string, option string) []string {
	if len(value) == 0 {
		return nil
	}

	star := false
	var includes []string
	var excludes []string
	for _, s := range strings.Split(value, ",") {
		s = strings.TrimSpace(s)
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

	if len(includes) == 0 && !star && len(excludes) == 0 {
		return nil
	}

	// Dedupe the variations by their normalized values
	variations := make(map[tsoptions.CompilerOptionsValue]string)

	// add (and deduplicate) all included entries
	for _, include := range includes {
		value, ok := getValueOfOptionString(option, include)
		if ok {
			variations[value] = include
		} else {
			variations[include] = include
		}
	}

	allValues := getAllValuesForOption(option)
	if star && len(allValues) > 0 {
		// add all entries
		for _, include := range allValues {
			value, ok := getValueOfOptionString(option, include)
			if ok {
				variations[value] = include
			} else {
				variations[include] = include
			}
		}
	}

	// remove all excluded entries
	for _, exclude := range excludes {
		value, ok := getValueOfOptionString(option, exclude)
		if ok {
			delete(variations, value)
		} else {
			delete(variations, exclude)
		}
	}

	if len(variations) == 0 {
		panic(fmt.Sprintf("Variations in test option '@%s' resulted in an empty set.", option))
	}
	return slices.Collect(maps.Values(variations))
}

func getValueOfOptionString(option string, value string) (tsoptions.CompilerOptionsValue, bool) {
	optionDecl := getCommandLineOption(option)
	if optionDecl == nil {
		return nil, false
	}
	switch optionDecl.Kind {
	case tsoptions.CommandLineOptionTypeString,
		tsoptions.CommandLineOptionTypeList,
		tsoptions.CommandLineOptionTypeListOrElement,
		tsoptions.CommandLineOptionTypeObject:
		return value, true
	case tsoptions.CommandLineOptionTypeEnum:
		enumVal, ok := optionDecl.EnumMap().Get(strings.ToLower(value))
		return enumVal, ok
	case tsoptions.CommandLineOptionTypeNumber:
		numVal, err := strconv.Atoi(value)
		if err != nil {
			return nil, false
		}
		return numVal, true
	case tsoptions.CommandLineOptionTypeBoolean:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return nil, false
		}
		return boolVal, true
	}
	return nil, false
}

func getCommandLineOption(option string) *tsoptions.CommandLineOption {
	return core.Find(tsoptions.OptionsDeclarations, func(optionDecl *tsoptions.CommandLineOption) bool {
		return strings.ToLower(optionDecl.Name) == strings.ToLower(option)
	})
}

func getAllValuesForOption(option string) []string {
	optionDecl := getCommandLineOption(option)
	if optionDecl == nil {
		return nil
	}
	switch optionDecl.Kind {
	case tsoptions.CommandLineOptionTypeEnum:
		return slices.Collect(optionDecl.EnumMap().Keys())
	case tsoptions.CommandLineOptionTypeBoolean:
		return []string{"true", "false"}
	}
	return nil
}

func computeFileBasedTestConfigurationVariations(variationCount int, optionEntries [][]string) []TestConfiguration {
	configurations := make([]TestConfiguration, 0, variationCount)
	computeFileBasedTestConfigurationVariationsWorker(&configurations, optionEntries, 0, make(map[string]string))
	return configurations
}

func computeFileBasedTestConfigurationVariationsWorker(
	configurations *[]TestConfiguration,
	optionEntries [][]string,
	index int,
	variationState TestConfiguration,
) {
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

func GetConfigNameFromFileName(filename string) string {
	basenameLower := strings.ToLower(tspath.GetBaseFileName(filename))
	if basenameLower == "tsconfig.json" || basenameLower == "jsconfig.json" {
		return basenameLower
	}
	return ""
}
