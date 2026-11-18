package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseWithUpcomingFlag(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
		lang string
	}{
		{utils.SubMinutes(1), "1 minute", "en"},
		{utils.SubMinutes(2), "2 minutes", "en"},
		{utils.SubMinutes(3), "3 minutes", "en"},
		{utils.SubMinutes(4), "4 minutes", "en"},
		{utils.SubMinutes(5), "5 minutes", "en"},
		{utils.SubMinutes(10), "10 minutes", "en"},
		{utils.SubMinutes(11), "11 minutes", "en"},
		{utils.SubMinutes(20), "20 minutes", "en"},
		{utils.SubMinutes(21), "21 minutes", "en"},
		{utils.SubMinutes(22), "22 minutes", "en"},
		{utils.SubMinutes(30), "30 minutes", "en"},
		{utils.SubMinutes(31), "31 minutes", "en"},
		{utils.SubHours(1), "1 hour", "en"},
		{utils.SubHours(2), "2 hours", "en"},
		{utils.SubHours(9), "9 hours", "en"},
		{utils.SubHours(10), "10 hours", "en"},
		{utils.SubHours(11), "11 hours", "en"},
		{utils.SubHours(20), "20 hours", "en"},
		{utils.SubHours(21), "21 hours", "en"},
		{utils.SubHours(23), "23 hours", "en"},
		{utils.SubDays(1), "1 day", "en"},
		{utils.SubDays(2), "2 days", "en"},
		{utils.SubDays(6), "6 days", "en"},
		{utils.SubWeeks(1), "1 week", "en"},
		{utils.SubWeeks(2), "2 weeks", "en"},
		{utils.SubWeeks(3), "3 weeks", "en"},
		{utils.SubMonths(1), "1 month", "en"},
		{utils.SubMonths(2), "2 months", "en"},
		{utils.SubMonths(3), "3 months", "en"},
		{utils.SubMonths(4), "4 months", "en"},
		{utils.SubMonths(5), "5 months", "en"},
		{utils.SubMonths(6), "6 months", "en"},
		{utils.SubMonths(7), "7 months", "en"},
		{utils.SubMonths(8), "8 months", "en"},
		{utils.SubMonths(9), "9 months", "en"},
		{utils.SubMonths(10), "10 months", "en"},
		{utils.SubMonths(11), "11 months", "en"},
		{utils.SubMonths(12), "1 year", "en"},
		{utils.SubYears(4), "4 years", "en"},
		{utils.SubYears(5), "5 years", "en"},
		{utils.SubYears(11), "11 years", "en"},
		{utils.SubYears(29), "29 years", "en"},
		{utils.SubYears(92), "92 years", "en"},
		{utils.SubMinutes(1), "1 минута", "ru"},
		{utils.SubMinutes(2), "2 минуты", "ru"},
		{utils.SubMonths(3), "3 месяца", "ru"},
		{utils.SubWeeks(2), "2 недели", "ru"},
		{utils.SubYears(77), "77 лет", "ru"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: tc.lang})

			res, err := ago.Parse(tc.date, ago.OptUpcoming)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("Result must be %s, but got %s instead", tc.res, res)
			}
		})
	}
}
