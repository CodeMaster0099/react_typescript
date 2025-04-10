package runner

import "testing"

type Runner interface {
	RunTests(t *testing.T)
}

func runTests(t *testing.T, runners []Runner) {
	for _, runner := range runners {
		runner.RunTests(t)
	}
}
