package tests

import (
	"testing"
	"time"

	. "github.com/SerhiiCho/timeago/v3"
)

func TestParseWithNoSuffixFlag(t *testing.T) {
	type TestCase struct {
		date   time.Time
		result string
		lang   string
	}

	cases := []TestCase{
		{subMinutes(1), "1 minute", "en"},
		{subMinutes(2), "2 minutes", "en"},
		{subMinutes(3), "3 minutes", "en"},
		{subMinutes(4), "4 minutes", "en"},
		{subMinutes(5), "5 minutes", "en"},
		{subMinutes(10), "10 minutes", "en"},
		{subMinutes(11), "11 minutes", "en"},
		{subMinutes(20), "20 minutes", "en"},
		{subMinutes(21), "21 minutes", "en"},
		{subMinutes(22), "22 minutes", "en"},
		{subMinutes(30), "30 minutes", "en"},
		{subMinutes(31), "31 minutes", "en"},
		{subHours(1), "1 hour", "en"},
		{subHours(2), "2 hours", "en"},
		{subHours(9), "9 hours", "en"},
		{subHours(10), "10 hours", "en"},
		{subHours(11), "11 hours", "en"},
		{subHours(20), "20 hours", "en"},
		{subHours(21), "21 hours", "en"},
		{subHours(23), "23 hours", "en"},
		{subDays(1), "1 day", "en"},
		{subDays(2), "2 days", "en"},
		{subDays(6), "6 days", "en"},
		{subWeeks(1), "1 week", "en"},
		{subWeeks(2), "2 weeks", "en"},
		{subWeeks(3), "3 weeks", "en"},
		{subMonths(1), "1 month", "en"},
		{subMonths(2), "2 months", "en"},
		{subMonths(3), "3 months", "en"},
		{subMonths(4), "4 months", "en"},
		{subMonths(5), "5 months", "en"},
		{subMonths(6), "6 months", "en"},
		{subMonths(7), "7 months", "en"},
		{subMonths(8), "8 months", "en"},
		{subMonths(9), "9 months", "en"},
		{subMonths(10), "10 months", "en"},
		{subMonths(11), "11 months", "en"},
		{subMonths(12), "1 year", "en"},
		{subYears(4), "4 years", "en"},
		{subYears(5), "5 years", "en"},
		{subYears(11), "11 years", "en"},
		{subYears(29), "29 years", "en"},
		{subYears(92), "92 years", "en"},
		{subMinutes(1), "1 минута", "ru"},
		{subMinutes(2), "2 минуты", "ru"},
		{subMonths(3), "3 месяца", "ru"},
		{subWeeks(2), "2 недели", "ru"},
		{subYears(77), "77 лет", "ru"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			SetConfig(Config{
				Language: tc.lang,
			})

			if res := Parse(tc.date, "noSuffix"); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}
