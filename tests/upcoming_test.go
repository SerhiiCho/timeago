package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v4"
)

func TestParseWithUpcomingFlag(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
		lang string
	}{
		{subMins(1), "1 minute", "en"},
		{subMins(2), "2 minutes", "en"},
		{subMins(3), "3 minutes", "en"},
		{subMins(4), "4 minutes", "en"},
		{subMins(5), "5 minutes", "en"},
		{subMins(10), "10 minutes", "en"},
		{subMins(11), "11 minutes", "en"},
		{subMins(20), "20 minutes", "en"},
		{subMins(21), "21 minutes", "en"},
		{subMins(22), "22 minutes", "en"},
		{subMins(30), "30 minutes", "en"},
		{subMins(31), "31 minutes", "en"},
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
		{subMins(1), "1 минута", "ru"},
		{subMins(2), "2 минуты", "ru"},
		{subMonths(3), "3 месяца", "ru"},
		{subWeeks(2), "2 недели", "ru"},
		{subYears(77), "77 лет", "ru"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: tc.lang})

			res, err := ago.Parse(tc.date, ago.OptUpcoming)

			if err != nil {
				test.Errorf("want nil error, got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("want %q, got %q instead", tc.res, res)
			}
		})
	}
}
