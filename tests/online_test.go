package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseWithOnlineFlag(t *testing.T) {
	type TestCase struct {
		date time.Time
		res  string
	}

	cases := []TestCase{
		{utils.SubMinutes(1), "1 minute ago"},
		{utils.SubMinutes(2), "2 minutes ago"},
		{utils.SubMinutes(9), "9 minutes ago"},
		{utils.SubMinutes(10), "10 minutes ago"},
		{utils.SubMinutes(11), "11 minutes ago"},
		{utils.SubMinutes(20), "20 minutes ago"},
		{utils.SubMinutes(21), "21 minutes ago"},
		{utils.SubMinutes(22), "22 minutes ago"},
		{utils.SubMinutes(30), "30 minutes ago"},
		{utils.SubMinutes(31), "31 minutes ago"},
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
		{utils.SubDays(6), "6 days ago"},
		{utils.SubWeeks(1), "1 week ago"},
		{utils.SubWeeks(2), "2 weeks ago"},
		{utils.SubWeeks(3), "3 weeks ago"},
		{utils.SubMonths(1), "1 month ago"},
		{utils.SubMonths(2), "2 months ago"},
		{utils.SubMonths(3), "3 months ago"},
		{utils.SubMonths(4), "4 months ago"},
		{utils.SubMonths(5), "5 months ago"},
		{utils.SubMonths(6), "6 months ago"},
		{utils.SubMonths(7), "7 months ago"},
		{utils.SubMonths(8), "8 months ago"},
		{utils.SubMonths(9), "9 months ago"},
		{utils.SubMonths(10), "10 months ago"},
		{utils.SubMonths(11), "11 months ago"},
		{utils.SubMonths(12), "1 year ago"},
		{utils.SubYears(1), "1 year ago"},
		{utils.SubYears(2), "2 years ago"},
		{utils.SubYears(21), "21 years ago"},
		{utils.SubYears(31), "31 years ago"},
		{utils.SubYears(100), "100 years ago"},
	}

	for i := 0; i < 60; i++ {
		cases = append(cases, TestCase{utils.SubSeconds(time.Duration(i)), "Online"})
	}

	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			res, err := ago.Parse(tc.date, ago.OptOnline)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("Result must be %q, but got %q instead", tc.res, res)
			}
		})
	}
}

func TestOnlineThresholdConfiguration(t *testing.T) {
	cases := []struct {
		date      time.Time
		res       string
		threshold int
	}{
		{utils.SubSeconds(10), "Online", 12},
		{utils.SubSeconds(10), "10 seconds ago", 8},
		{utils.SubSeconds(20), "Online", 22},
		{utils.SubSeconds(20), "20 seconds ago", 18},
		{utils.SubSeconds(30), "Online", 32},
		{utils.SubSeconds(30), "30 seconds ago", 28},
		{utils.SubSeconds(40), "Online", 42},
		{utils.SubSeconds(40), "40 seconds ago", 38},
		{utils.SubSeconds(50), "Online", 52},
		{utils.SubSeconds(50), "50 seconds ago", 48},
		{utils.SubSeconds(53), "53 seconds ago", 5},
		{utils.SubSeconds(57), "Online", 60},
		{utils.SubSeconds(57), "57 seconds ago", 55},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{
				Language:        ago.LangEn,
				OnlineThreshold: tc.threshold,
			})

			out, err := ago.Parse(tc.date, ago.OptOnline)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if out != tc.res {
				test.Errorf("Result must be %q, but got %q instead", tc.res, out)
			}
		})
	}
}

func TestJustNowThresholdConfiguration(t *testing.T) {
	cases := []struct {
		date      time.Time
		res       string
		threshold int
	}{
		{utils.SubSeconds(10), "Только что", 12},
		{utils.SubSeconds(10), "10 секунд назад", 8},
		{utils.SubSeconds(20), "Только что", 22},
		{utils.SubSeconds(20), "20 секунд назад", 18},
		{utils.SubSeconds(30), "Только что", 32},
		{utils.SubSeconds(30), "30 секунд назад", 28},
		{utils.SubSeconds(40), "Только что", 42},
		{utils.SubSeconds(40), "40 секунд назад", 38},
		{utils.SubSeconds(50), "Только что", 52},
		{utils.SubSeconds(50), "50 секунд назад", 48},
		{utils.SubSeconds(53), "53 секунды назад", 5},
		{utils.SubSeconds(57), "Только что", 60},
		{utils.SubSeconds(57), "57 секунд назад", 55},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{
				Language:         ago.LangRu,
				JustNowThreshold: tc.threshold,
			})

			out, err := ago.Parse(tc.date, ago.OptJustNow)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if out != tc.res {
				test.Errorf("Result must be %q, but got %q instead", tc.res, out)
			}
		})
	}
}
