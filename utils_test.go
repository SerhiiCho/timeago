package timeago

import (
	"testing"
)

func TestParseLangSet(t *testing.T) {
	t.Run("function returns Lang model with needed values", func(test *testing.T) {
		lang, err := parseLangSet("langs/ru.json")
		if err != nil {
			t.Errorf("want no error, got %v instead", err)
		}

		if lang.Ago != "назад" {
			t.Errorf("want %q, got %q instead", "назад", lang.Ago)
		}
	})
}
