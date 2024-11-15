package runner

import "testing"

type runner interface {
	EnumerateTestFiles() []string
	RunTests(t *testing.T)
}

func runTests(t *testing.T, runners []runner) {
	// !!!
	// const seen = new Map<string, string>();
	// const dupes: [string, string][] = [];
	for _, runner := range runners {
		runner.RunTests(t)
		// !!! Check for duplicates in case baselines are going in the same folder
	}
	// !!!
	// 	if (dupes.length) {
	// 		throw new Error(`${dupes.length} Tests with duplicate baseline names:
	// ${JSON.stringify(dupes, undefined, 2)}`);
	// 	}
}
