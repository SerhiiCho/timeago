package langs

import "testing"

func TestGetEnglish(t *testing.T) {
	if en["weeks2"] != "weeks" {
		t.Error("getEnglish must return map of strings with translations for english language")
	}

	if en["days"] != "days" {
		t.Error("getEnglish must return map of strings with translations for english language")
	}
}
