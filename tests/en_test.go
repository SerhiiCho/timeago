package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseEn(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{utils.SubMinutes(1), "1 minute ago"},
		{utils.SubMinutes(2), "2 minutes ago"},
		{utils.SubMinutes(5), "5 minutes ago"},
		{utils.SubMinutes(9), "9 minutes ago"},
		{utils.SubMinutes(10), "10 minutes ago"},
		{utils.SubMinutes(11), "11 minutes ago"},
		{utils.SubMinutes(20), "20 minutes ago"},
		{utils.SubMinutes(21), "21 minutes ago"},
		{utils.SubMinutes(22), "22 minutes ago"},
		{utils.SubMinutes(30), "30 minutes ago"},
		{utils.SubMinutes(31), "31 minutes ago"},
		{utils.SubMinutes(59), "59 minutes ago"},
		{utils.SubHours(1), "1 hour ago"},
		{utils.SubHours(2), "2 hours ago"},
		{utils.SubHours(9), "9 hours ago"},
		{utils.SubHours(10), "10 hours ago"},
		{utils.SubHours(11), "11 hours ago"},
		{utils.SubHours(20), "20 hours ago"},
		{utils.SubHours(21), "21 hours ago"},
		{utils.SubHours(23), "23 hours ago"},
		{utils.SubDays(1), "1 day ago"},
		{utils.SubDays(2), "2 days ago"},
		{utils.SubDays(4), "4 days ago"},
		{utils.SubDays(5), "5 days ago"},
		{utils.SubDays(6), "6 days ago"},
		{utils.SubWeeks(1), "1 week ago"},
		{utils.SubWeeks(2), "2 weeks ago"},
		{utils.SubWeeks(3), "3 weeks ago"},
		{utils.SubMonths(1), "1 month ago"},
		{utils.SubMonths(2), "2 months ago"},
		{utils.SubMonths(9), "9 months ago"},
		{utils.SubMonths(11), "11 months ago"},
		{utils.SubYears(1), "1 year ago"},
		{utils.SubYears(2), "2 years ago"},
		{utils.SubYears(21), "21 years ago"},
		{utils.SubYears(31), "31 years ago"},
		{utils.SubYears(100), "100 years ago"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangEn})

			res, err := ago.Parse(tc.date)

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
		{utils.SubSeconds(0), []string{"0 seconds ago", "1 second ago"}},
		{utils.SubSeconds(1), []string{"1 second ago", "2 seconds ago"}},
		{utils.SubSeconds(2), []string{"2 seconds ago", "3 seconds ago"}},
		{utils.SubSeconds(9), []string{"9 seconds ago", "10 seconds ago"}},
		{utils.SubSeconds(10), []string{"10 seconds ago", "11 seconds ago"}},
		{utils.SubSeconds(11), []string{"11 seconds ago", "12 seconds ago"}},
		{utils.SubSeconds(20), []string{"20 seconds ago", "21 seconds ago"}},
		{utils.SubSeconds(21), []string{"21 seconds ago", "22 seconds ago"}},
		{utils.SubSeconds(22), []string{"22 seconds ago", "23 seconds ago"}},
		{utils.SubSeconds(30), []string{"30 seconds ago", "31 seconds ago"}},
		{utils.SubSeconds(59), []string{"59 seconds ago", "1 minute ago"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangEn})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res[0] && res != tc.res[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.res[0], tc.res[1], res)
			}
		})
	}
}
