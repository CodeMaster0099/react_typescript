package runner

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/tspath"
)

// >> TODO: do we need a runner base struct + composition?
func enumerateFiles(folder string, testRegex *regexp.Regexp, recursive bool) ([]string, error) {
	files, err := listFiles(folder, testRegex, recursive)
	if err != nil {
		return nil, err
	}
	return core.Map(files, tspath.NormalizeSlashes), nil
}

// >> TODO: use filepath.WalkDir?
func listFiles(path string, spec *regexp.Regexp, recursive bool) ([]string, error) {
	return listFilesWorker(spec, recursive, path)
}

func listFilesWorker(spec *regexp.Regexp, recursive bool, folder string) ([]string, error) {
	entries, err := os.ReadDir(folder)
	if err == nil {
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
