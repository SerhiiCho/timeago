package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestJustNowOption(t *testing.T) {
	cases := []struct {
		date time.Time
	}{
		{utils.SubSeconds(0)},
		{utils.SubSeconds(1)},
		{utils.SubSeconds(2)},
		{utils.SubSeconds(9)},
		{utils.SubSeconds(10)},
		{utils.SubSeconds(11)},
		{utils.SubSeconds(20)},
		{utils.SubSeconds(21)},
		{utils.SubSeconds(22)},
		{utils.SubSeconds(30)},
		{utils.SubSeconds(58)},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangEn})

			res, err := ago.Parse(tc.date, ago.OptJustNow)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != "Just now" {
				test.Errorf("Result must be %q, but got %s instead", "Just now", res)
			}
		})
	}
}
