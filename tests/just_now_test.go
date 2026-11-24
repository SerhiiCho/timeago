package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v4"
)

func TestJustNowOption(t *testing.T) {
	cases := []struct {
		date time.Time
	}{
		{subSecs(0)},
		{subSecs(1)},
		{subSecs(2)},
		{subSecs(9)},
		{subSecs(10)},
		{subSecs(11)},
		{subSecs(20)},
		{subSecs(21)},
		{subSecs(22)},
		{subSecs(30)},
		{subSecs(58)},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangEn})

			res, err := ago.Parse(tc.date, ago.OptJustNow)

			if err != nil {
				test.Errorf("want nil error, got %v instead", err)
			}

			if res != "Just now" {
				test.Errorf("want %q, got %q instead", "Just now", res)
			}
		})
	}
}
