package timeago

import (
	"testing"
)

func Test_parseJsonIntoLang(t *testing.T) {
	t.Run("Function returns Lang model with needed values", func(test *testing.T) {
		res, err := parseLangSet("langs/ru.json")

		if err != nil {
			t.Errorf("Function parseLangSet must return Lang model, but returned error %v", err)
		}

		if res.Ago != "назад" {
			t.Errorf("Function needs to return model with value назад, but returned %v", res.Ago)
		}
	})
}
