package langs

import "testing"

func TestGetRussian(t *testing.T) {
	if ru["weeks2"] != "недель" {
		t.Error("getRussian must return map of strings with translations for russian language")
	}

	if ru["days"] != "дня" {
		t.Error("getRussian must return map of strings with translations for russian language")
	}
}
