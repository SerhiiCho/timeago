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
