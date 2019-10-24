package timeago

import "testing"

func TestGetRussian(t *testing.T) {
	if getRussian()["weeks2"] != "недель" {
		t.Error("getRussian must return map of strings with translations for russian lang")
	}

	if getRussian()["days"] != "дня" {
		t.Error("getRussian must return map of strings with translations for russian lang")
	}
}
