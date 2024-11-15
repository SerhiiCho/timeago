package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/config"
)

func TestLocationChanges(t *testing.T) {
	cases := []struct {
		loc      string
		subHours time.Duration
		expect   string
	}{
		{"Europe/Kiev", 2, "2 hours ago"},
		{"America/New_York", 2, "5 hours"},
	}

	for _, tc := range cases {
		t.Run(tc.loc, func(test *testing.T) {
			timeago.Configure(&config.Config{Location: tc.loc})

			d := subHours(tc.subHours)

			if res := timeago.Parse(d); res != tc.expect {
				test.Errorf("Result must be %q, but got %q instead", tc.expect, res)
			}
		})
	}
}
