package tests

import "testing"

// assertResult checks the outcome of a parse call.
// Safe for goroutines: uses t.Errorf, not t.Fatalf.
func assertResult(t *testing.T, want string, got string, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("want nil error, got %v instead", err)
		return
	}

	if got != want {
		t.Errorf("want %q, got %q instead", want, got)
	}
}
