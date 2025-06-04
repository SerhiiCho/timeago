package timeago

import (
	"encoding/json"
	"testing"
)

func TestParseJsonIntoLang(t *testing.T) {
	t.Run("Function returns Lang model with needed values", func(test *testing.T) {
		res, err := parseLangSet("langs/ru.json", langsFS)

		if err != nil {
			t.Errorf("Function parseLangSet must return Lang model, but returned error %v", err)
		}

		if res.Ago != "назад" {
			t.Errorf("Function needs to return model with value назад, but returned %v", res.Ago)
		}
	})
}

func TestIsFilePresent(t *testing.T) {
	t.Run("isFilePresent return false if file doesn't exist", func(test *testing.T) {
		res, _ := isFilePresent("somerandompath")

		if res {
			t.Error("isFilePresent must return false, because filepath points to a file that doesn't exist")
		}
	})

	t.Run("isFilePresent return true if file exist", func(test *testing.T) {
		ok, err := isFilePresent("timeago.go")

		if err != nil {
			t.Errorf("isFilePresent must return true, because filepath points to a file that exists, but returned error %v", err)
		}

		if !ok {
			t.Error("isFilePresent must return true, because filepath points to a file that exists")
		}
	})
}

func TestReadFile(t *testing.T) {
	t.Run("readFile returns content of the file", func(test *testing.T) {
		res := readFile("langs/en.json")

		var js json.RawMessage
		err := json.Unmarshal(res, &js)

		if err != nil {
			t.Errorf("Function readFile must return JSON object but %s returned", string(res))
		}
	})
}
