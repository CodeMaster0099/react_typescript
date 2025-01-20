package baseline

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/microsoft/typescript-go/internal/repo"
	"github.com/pkg/diff"
)

type Options struct {
	Subfolder string
	IsDiff    bool
}

const NoContent = "<no content>"

func Run(t *testing.T, fileName string, actual string, opts Options) {
	if opts.IsDiff {
		diff := getBaselineDiff(t, actual, fileName)
		diffFileName := fileName + ".diff"
		opts.Subfolder = "diff/" + opts.Subfolder
		// Write original baseline in addition to diff baseline
		if actual != NoContent {
			localFileName := localPath(fileName, opts.Subfolder)
			if err := os.MkdirAll(filepath.Dir(localFileName), 0o755); err != nil {
				t.Fatal(fmt.Errorf("failed to create directories for the local baseline file %s: %w", localFileName, err))
			}
			if err := os.WriteFile(localFileName, []byte(actual), 0o644); err != nil {
				t.Fatal(fmt.Errorf("failed to write the local baseline file %s: %w", localFileName, err))
			}
		}
		writeComparison(t, diff, diffFileName, opts)
	} else {
		writeComparison(t, actual, fileName, opts)
	}
}

func getBaselineDiff(t *testing.T, actual string, fileName string) string {
	expected := NoContent
	refFileName := tsBaselinePath(fileName)
	if content, err := os.ReadFile(refFileName); err == nil {
		expected = string(content)
	}
	if actual == expected {
		return NoContent
	}
	var b strings.Builder
	err := diff.Text("", "", actual, expected, &b)
	if err != nil {
		t.Fatalf("failed to diff the actual and expected content: %v", err)
	}
	return b.String()
}

func tsBaselinePath(fileName string) string {
	return filepath.Join(repo.TestDataPath, "..", "_submodules", "TypeScript", "tests", "baselines", "reference", fileName)
}

func writeComparison(t *testing.T, actual string, relativeFileName string, opts Options) {
	if actual == "" {
		panic("The generated content was \"\". Return 'baseline.NoContent' if no baselining is required.")
	}
	var (
		localFileName     string
		referenceFileName string
	)

	localFileName = localPath(relativeFileName, opts.Subfolder)
	referenceFileName = referencePath(relativeFileName, opts.Subfolder)
	expected := getExpectedContent(relativeFileName, opts)
	if _, err := os.Stat(localFileName); err == nil {
		if err := os.Remove(localFileName); err != nil {
			t.Fatal(fmt.Errorf("failed to remove the local baseline file %s: %w", localFileName, err))
		}
	}
	if actual != expected {
		if err := os.MkdirAll(filepath.Dir(localFileName), 0o755); err != nil {
			t.Fatal(fmt.Errorf("failed to create directories for the local baseline file %s: %w", localFileName, err))
		}
		if actual == NoContent {
			if err := os.WriteFile(localFileName+".delete", []byte{}, 0o644); err != nil {
				t.Fatal(fmt.Errorf("failed to write the local baseline file %s: %w", localFileName+".delete", err))
			}
		} else if err := os.WriteFile(localFileName, []byte(actual), 0o644); err != nil {
			t.Fatal(fmt.Errorf("failed to write the local baseline file %s: %w", localFileName, err))
		}

		if _, err := os.Stat(referenceFileName); err != nil {
			t.Errorf("New baseline created at %s.", localFileName)
		} else {
			t.Errorf("The baseline file %s has changed. (Run `hereby baseline-accept` if the new baseline is correct.)", relativeFileName)
		}
	}
}

func getExpectedContent(relativeFileName string, opts Options) string {
	refFileName := referencePath(relativeFileName, opts.Subfolder)
	expected := NoContent
	if content, err := os.ReadFile(refFileName); err == nil {
		expected = string(content)
	}
	return expected
}

func localPath(fileName string, subfolder string) string {
	return filepath.Join(repo.TestDataPath, "baselines", "local", subfolder, fileName)
}

func referencePath(fileName string, subfolder string) string {
	return filepath.Join(repo.TestDataPath, "baselines", "reference", subfolder, fileName)
}
