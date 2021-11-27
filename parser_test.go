package timeago

import (
	"fmt"
	"testing"

	"github.com/SerhiiCho/timeago/utils"
)

func TestGetFiles(t *testing.T) {

}

func TestGetAllLanguages(t *testing.T) {
	cases := []string{"ru", "en"}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Returns slice with %s language", tc), func(test *testing.T) {
			result := getAllLanguages()
			contains := utils.SliceContains(result, tc)

			if !contains {
				t.Errorf("Function has to return slice with %s, instead got: %v", tc, result)
			}
		})
	}
}

func TestParseNeededFile(t *testing.T) {
	t.Run("Function returns Lang model with needed values", func(test *testing.T) {
		result := parseNeededFile("langs/ru.json")

		if result.Ago != "назад" {
			t.Errorf("Function needs to return model with value назад, but returned %v", result.Ago)
		}
	})
}
