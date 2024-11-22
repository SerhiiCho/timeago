package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago/v3"
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
		{subMinutes(1), "1 minute"},
		{subMinutes(5), "5 minutes"},
		{subMinutes(10), "10 minutes"},
		{subMinutes(30), "30 minutes"},
		{subMinutes(31), "31 minutes"},
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
		cases = append(cases, TestCase{subSeconds(time.Duration(i)), "Online"})
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			timeago.Reconfigure(timeago.Config{Language: langEn})

			res, err := timeago.Parse(tc.date, "online", "noSuffix")

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("Result must be %s, but got %s instead", tc.res, res)
			}
		})
	}
}
