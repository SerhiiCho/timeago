package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v4"
)

func TestParseWithNoSuffixFlag(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
		lang string
	}{
		{subMins(1), "1 minute", ago.LangEn},
		{subMins(2), "2 minutes", ago.LangEn},
		{subMins(3), "3 minutes", ago.LangEn},
		{subMins(4), "4 minutes", ago.LangEn},
		{subMins(5), "5 minutes", ago.LangEn},
		{subMins(10), "10 minutes", ago.LangEn},
		{subMins(11), "11 minutes", ago.LangEn},
		{subMins(20), "20 minutes", ago.LangEn},
		{subMins(21), "21 minutes", ago.LangEn},
		{subMins(22), "22 minutes", ago.LangEn},
		{subMins(30), "30 minutes", ago.LangEn},
		{subMins(31), "31 minutes", ago.LangEn},
		{subHours(1), "1 hour", ago.LangEn},
		{subHours(2), "2 hours", ago.LangEn},
		{subHours(9), "9 hours", ago.LangEn},
		{subHours(10), "10 hours", ago.LangEn},
		{subHours(11), "11 hours", ago.LangEn},
		{subHours(20), "20 hours", ago.LangEn},
		{subHours(21), "21 hours", ago.LangEn},
		{subHours(23), "23 hours", ago.LangEn},
		{subDays(1), "1 day", ago.LangEn},
		{subDays(2), "2 days", ago.LangEn},
		{subDays(6), "6 days", ago.LangEn},
		{subWeeks(1), "1 week", ago.LangEn},
		{subWeeks(2), "2 weeks", ago.LangEn},
		{subWeeks(3), "3 weeks", ago.LangEn},
		{subMonths(1), "1 month", ago.LangEn},
		{subMonths(2), "2 months", ago.LangEn},
		{subMonths(3), "3 months", ago.LangEn},
		{subMonths(4), "4 months", ago.LangEn},
		{subMonths(5), "5 months", ago.LangEn},
		{subMonths(6), "6 months", ago.LangEn},
		{subMonths(7), "7 months", ago.LangEn},
		{subMonths(8), "8 months", ago.LangEn},
		{subMonths(9), "9 months", ago.LangEn},
		{subMonths(10), "10 months", ago.LangEn},
		{subMonths(11), "11 months", ago.LangEn},
		{subMonths(12), "1 year", ago.LangEn},
		{subYears(4), "4 years", ago.LangEn},
		{subYears(5), "5 years", ago.LangEn},
		{subYears(11), "11 years", ago.LangEn},
		{subYears(29), "29 years", ago.LangEn},
		{subYears(92), "92 years", ago.LangEn},
		{subMins(1), "1 минута", ago.LangRu},
		{subMins(2), "2 минуты", ago.LangRu},
		{subMonths(3), "3 месяца", ago.LangRu},
		{subWeeks(2), "2 недели", ago.LangRu},
		{subYears(77), "77 лет", ago.LangRu},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: tc.lang})

			res, err := ago.Parse(tc.date, ago.OptNoSuffix)

			if err != nil {
				test.Errorf("want nil error, got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("want %q, got %q instead", tc.res, res)
			}
		})
	}
}
