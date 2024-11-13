package baseline

import (
	"regexp"
	"strings"

	"github.com/microsoft/typescript-go/internal/compiler"
	"github.com/microsoft/typescript-go/internal/tspath"
)

var testPathPrefix = regexp.MustCompile(`(?:(file:\/{3})|\/)\.(?:ts|lib|src)\/`)
var testPathCharacters = regexp.MustCompile(`[\^<>:"|?*%]`)
var testPathDotDot = regexp.MustCompile(`\.\.\/`)

// This is done so tests work on windows _and_ linux
var canonicalizeForHarness = strings.ToLower

var libFolder = "built/local/"
var builtFolder = "/.ts"

func removeTestPathPrefixes(text string, retainTrailingDirectorySeparator bool) string {
	testPathPrefix.ReplaceAllStringFunc(text, func(scheme string) string {
		if scheme != "" {
			return scheme
		}
		if retainTrailingDirectorySeparator {
			return "/"
		}
		return ""
	})
	return testPathPrefix.ReplaceAllString(text, "/")
}

func isDefaultLibraryFile(filePath string) bool {
	fileName := tspath.GetBaseFileName(filePath)
	return strings.HasPrefix(fileName, "lib.") && strings.HasSuffix(fileName, compiler.ExtensionDts)
}

func isBuiltFile(filePath string) bool {
	return strings.HasPrefix(filePath, libFolder) || strings.HasPrefix(filePath, tspath.EnsureTrailingDirectorySeparator(builtFolder))
}

func isTsConfigFile(path string) bool {
	return strings.Contains(path, "tsconfig") && strings.Contains(path, "json")
}

func sanitizeTestFilePath(name string) string {
	path := testPathCharacters.ReplaceAllString(name, "_")
	path = tspath.NormalizeSlashes(path)
	path = testPathDotDot.ReplaceAllString(path, "__dotdot/")
	path = string(tspath.ToPath(path, "", canonicalizeForHarness))
	return strings.TrimPrefix(path, "/")
}
