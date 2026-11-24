package tests

import "testing"

// assertResult checks the outcome of a parse call.
// Safe for goroutines: uses t.Errorf, not t.Fatalf.
func assertResult(t *testing.T, want string, got string, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Parse returned an error: %v", err)
		return
	}

	if got != want {
		t.Errorf("Result must be %q, but got %q instead", want, got)
	}
}
