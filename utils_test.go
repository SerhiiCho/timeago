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

func TestFileExists(t *testing.T) {
	t.Run("fileExists return false if file doesn't exist", func(test *testing.T) {
		result, _ := fileExists("somerandompath")

		if result {
			t.Errorf("Function fileExists must return false, because filepath points to a file that doesn't exist")
		}
	})

	t.Run("fileExists return true if file exist", func(test *testing.T) {
		result, _ := fileExists("timeago.go")

		if result == false {
			t.Errorf("Function fileExists must return true, because filepath points to a file that exists")
		}
	})
}
