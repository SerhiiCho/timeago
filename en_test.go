package timeago

import "testing"

func TestGetEnglish(t *testing.T) {
	if getEnglish()["weeks2"] != "weeks" {
		t.Error("getEnglish must return map of strings with translations for english language")
	}

	if getEnglish()["days"] != "days" {
		t.Error("getEnglish must return map of strings with translations for english language")
	}
}
