package timeago

import (
	"testing"
)

func TestParseJsonIntoLang(t *testing.T) {
	t.Run("Function returns Lang model with needed values", func(test *testing.T) {
		result := parseJsonIntoLang("langs/ru.json")

		if result.Ago != "назад" {
			t.Errorf("Function needs to return model with value назад, but returned %v", result.Ago)
		}
	})
}
