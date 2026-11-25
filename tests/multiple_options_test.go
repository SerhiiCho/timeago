package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v4"
)

func TestParseWithMultipleFlags(t *testing.T) {
	type TestCase struct {
		date time.Time
		res  string
	}

	cases := []struct {
		date time.Time
		res  string
	}{
		{subMins(1), "1 minute"},
		{subMins(5), "5 minutes"},
		{subMins(10), "10 minutes"},
		{subMins(30), "30 minutes"},
		{subMins(31), "31 minutes"},
		{subHours(1), "1 hour"},
		{subHours(11), "11 hours"},
		{subHours(20), "20 hours"},
		{subDays(6), "6 days"},
		{subWeeks(1), "1 week"},
		{subWeeks(2), "2 weeks"},
		{subWeeks(3), "3 weeks"},
		{subMonths(1), "1 month"},
		{subMonths(6), "6 months"},
		{subMonths(8), "8 months"},
		{subMonths(9), "9 months"},
		{subMonths(11), "11 months"},
		{subMonths(1), "1 month"},
		{subYears(2), "2 years"},
		{subYears(44), "44 years"},
	}

	for i := 0; i < 60; i++ {
		cases = append(cases, TestCase{subSecs(time.Duration(i)), "Online"})
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangEn})

			res, err := ago.Parse(tc.date, ago.OptOnline, ago.OptNoSuffix)

			if err != nil {
				test.Errorf("want nil error, got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("want %q, got %q instead", tc.res, res)
			}
		})
	}
}
