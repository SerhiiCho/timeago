package timeago

import (
	"encoding/json"
	"fmt"
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
		expect int
	}{
		{0, 0},
		{1, 1},
		{9, 9},
		{20, 0},
		{300, 0},
		{-1, 1},
		{253, 3},
		{23423252, 2},
		{22, 2},
		{4444, 4},
		{-24, 4},
		{23423521562348, 8},
		{23525, 5},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Result must be %d", tc.expect), func(test *testing.T) {
			if res := getLastNumberDigit(tc.number); res != tc.expect {
				test.Errorf("Result must be %d, but got %d instead", tc.expect, res)
			}
		})
	}
}
