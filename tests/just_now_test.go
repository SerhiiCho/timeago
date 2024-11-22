package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago/v3"
)

func TestJustNowOption(t *testing.T) {
	cases := []struct {
		date time.Time
	}{
		{subSeconds(0)},
		{subSeconds(1)},
		{subSeconds(2)},
		{subSeconds(9)},
		{subSeconds(10)},
		{subSeconds(11)},
		{subSeconds(20)},
		{subSeconds(21)},
		{subSeconds(22)},
		{subSeconds(30)},
		{subSeconds(58)},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			timeago.Reconfigure(timeago.Config{Language: langEn})

			res, err := timeago.Parse(tc.date, "justNow")

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != "Just now" {
				test.Errorf("Result must be %q, but got %s instead", "Just now", res)
			}
		})
	}
}
