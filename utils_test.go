package timeago

import (
	"testing"
)

func TestParseJsonIntoLang(t *testing.T) {
	t.Run("Function returns Lang model with needed values", func(test *testing.T) {
		res, err := parseLangSet("langs/ru.json")
		if err != nil {
			t.Errorf("want no error, got %v instead", err)
		}

		if res.Ago != "назад" {
			t.Errorf("want %q, got %q instead", "назад", res.Ago)
		}
	})
}
