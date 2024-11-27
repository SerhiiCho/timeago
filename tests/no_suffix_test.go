package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseWithNoSuffixFlag(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
		lang string
	}{
		{utils.SubMinutes(1), "1 minute", ago.LangEn},
		{utils.SubMinutes(2), "2 minutes", ago.LangEn},
		{utils.SubMinutes(3), "3 minutes", ago.LangEn},
		{utils.SubMinutes(4), "4 minutes", ago.LangEn},
		{utils.SubMinutes(5), "5 minutes", ago.LangEn},
		{utils.SubMinutes(10), "10 minutes", ago.LangEn},
		{utils.SubMinutes(11), "11 minutes", ago.LangEn},
		{utils.SubMinutes(20), "20 minutes", ago.LangEn},
		{utils.SubMinutes(21), "21 minutes", ago.LangEn},
		{utils.SubMinutes(22), "22 minutes", ago.LangEn},
		{utils.SubMinutes(30), "30 minutes", ago.LangEn},
		{utils.SubMinutes(31), "31 minutes", ago.LangEn},
		{utils.SubHours(1), "1 hour", ago.LangEn},
		{utils.SubHours(2), "2 hours", ago.LangEn},
		{utils.SubHours(9), "9 hours", ago.LangEn},
		{utils.SubHours(10), "10 hours", ago.LangEn},
		{utils.SubHours(11), "11 hours", ago.LangEn},
		{utils.SubHours(20), "20 hours", ago.LangEn},
		{utils.SubHours(21), "21 hours", ago.LangEn},
		{utils.SubHours(23), "23 hours", ago.LangEn},
		{utils.SubDays(1), "1 day", ago.LangEn},
		{utils.SubDays(2), "2 days", ago.LangEn},
		{utils.SubDays(6), "6 days", ago.LangEn},
		{utils.SubWeeks(1), "1 week", ago.LangEn},
		{utils.SubWeeks(2), "2 weeks", ago.LangEn},
		{utils.SubWeeks(3), "3 weeks", ago.LangEn},
		{utils.SubMonths(1), "1 month", ago.LangEn},
		{utils.SubMonths(2), "2 months", ago.LangEn},
		{utils.SubMonths(3), "3 months", ago.LangEn},
		{utils.SubMonths(4), "4 months", ago.LangEn},
		{utils.SubMonths(5), "5 months", ago.LangEn},
		{utils.SubMonths(6), "6 months", ago.LangEn},
		{utils.SubMonths(7), "7 months", ago.LangEn},
		{utils.SubMonths(8), "8 months", ago.LangEn},
		{utils.SubMonths(9), "9 months", ago.LangEn},
		{utils.SubMonths(10), "10 months", ago.LangEn},
		{utils.SubMonths(11), "11 months", ago.LangEn},
		{utils.SubMonths(12), "1 year", ago.LangEn},
		{utils.SubYears(4), "4 years", ago.LangEn},
		{utils.SubYears(5), "5 years", ago.LangEn},
		{utils.SubYears(11), "11 years", ago.LangEn},
		{utils.SubYears(29), "29 years", ago.LangEn},
		{utils.SubYears(92), "92 years", ago.LangEn},
		{utils.SubMinutes(1), "1 минута", ago.LangRu},
		{utils.SubMinutes(2), "2 минуты", ago.LangRu},
		{utils.SubMonths(3), "3 месяца", ago.LangRu},
		{utils.SubWeeks(2), "2 недели", ago.LangRu},
		{utils.SubYears(77), "77 лет", ago.LangRu},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: tc.lang})

			res, err := ago.Parse(tc.date, ago.OptNoSuffix)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("Result must be %s, but got %s instead", tc.res, res)
			}
		})
	}
}
