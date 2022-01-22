package timeago

import (
	"encoding/json"
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
			t.Error("Function fileExists must return false, because filepath points to a file that doesn't exist")
		}
	})

	t.Run("fileExists return true if file exist", func(test *testing.T) {
		result, _ := fileExists("timeago.go")

		if result == false {
			t.Error("Function fileExists must return true, because filepath points to a file that exists")
		}
	})
}

func TestGetFileContent(t *testing.T) {
	t.Run("getFileContent returns content of the file", func(test *testing.T) {
		result := getFileContent("langs/en.json")

		var js json.RawMessage
		err := json.Unmarshal(result, &js)

		if err != nil {
			t.Errorf("Function getFileContent must return JSON object but %s returned", string(result))
		}
	})
}

func TestGetLastNumberDigit(t *testing.T) {
	cases := []struct {
		number int
		result int
		name   string
	}{
		{0, 0, "must return 0"},
		{1, 1, "must return 1"},
		{9, 9, "must return 9"},
		{20, 0, "must return 0"},
		{253, 3, "must return 3"},
		{23423252, 2, "must return 2"},
		{223424342325, 5, "must return 5"},
		{23423521562348, 8, "must return 8"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(test *testing.T) {
			if res := getLastNumberDigit(tc.number); res != tc.result {
				test.Errorf("Result must be %d, but got %d instead", tc.result, res)
			}
		})
	}
}
