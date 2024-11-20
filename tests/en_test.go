package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago/v3"
)

const langEn = "en"

func TestParseEn(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{subMinutes(1), "1 minute ago"},
		{subMinutes(2), "2 minutes ago"},
		{subMinutes(5), "5 minutes ago"},
		{subMinutes(9), "9 minutes ago"},
		{subMinutes(10), "10 minutes ago"},
		{subMinutes(11), "11 minutes ago"},
		{subMinutes(20), "20 minutes ago"},
		{subMinutes(21), "21 minutes ago"},
		{subMinutes(22), "22 minutes ago"},
		{subMinutes(30), "30 minutes ago"},
		{subMinutes(31), "31 minutes ago"},
		{subMinutes(59), "59 minutes ago"},
		{subHours(1), "1 hour ago"},
		{subHours(2), "2 hours ago"},
		{subHours(9), "9 hours ago"},
		{subHours(10), "10 hours ago"},
		{subHours(11), "11 hours ago"},
		{subHours(20), "20 hours ago"},
		{subHours(21), "21 hours ago"},
		{subHours(23), "23 hours ago"},
		{subDays(1), "1 day ago"},
		{subDays(2), "2 days ago"},
		{subDays(4), "4 days ago"},
		{subDays(5), "5 days ago"},
		{subDays(6), "6 days ago"},
		{subWeeks(1), "1 week ago"},
		{subWeeks(2), "2 weeks ago"},
		{subWeeks(3), "3 weeks ago"},
		{subMonths(1), "1 month ago"},
		{subMonths(2), "2 months ago"},
		{subMonths(9), "9 months ago"},
		{subMonths(11), "11 months ago"},
		{subYears(1), "1 year ago"},
		{subYears(2), "2 years ago"},
		{subYears(21), "21 years ago"},
		{subYears(31), "31 years ago"},
		{subYears(100), "100 years ago"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			timeago.ClearCache()
			timeago.Configure(timeago.Config{Language: langEn})

			res, err := timeago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("Result must be %s, but got %s instead", tc.res, res)
			}
		})
	}
}

func TestParseEnWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{subSeconds(0), []string{"0 seconds ago", "1 second ago"}},
		{subSeconds(1), []string{"1 second ago", "2 seconds ago"}},
		{subSeconds(2), []string{"2 seconds ago", "3 seconds ago"}},
		{subSeconds(9), []string{"9 seconds ago", "10 seconds ago"}},
		{subSeconds(10), []string{"10 seconds ago", "11 seconds ago"}},
		{subSeconds(11), []string{"11 seconds ago", "12 seconds ago"}},
		{subSeconds(20), []string{"20 seconds ago", "21 seconds ago"}},
		{subSeconds(21), []string{"21 seconds ago", "22 seconds ago"}},
		{subSeconds(22), []string{"22 seconds ago", "23 seconds ago"}},
		{subSeconds(30), []string{"30 seconds ago", "31 seconds ago"}},
		{subSeconds(59), []string{"59 seconds ago", "1 minute ago"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			timeago.ClearCache()
			timeago.Configure(timeago.Config{Language: langEn})

			res, err := timeago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res[0] && res != tc.res[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.res[0], tc.res[1], res)
			}
		})
	}
}
