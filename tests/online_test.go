package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago/v3"
)

func TestParseWithOnlineFlag(t *testing.T) {
	type TestCase struct {
		date time.Time
		res  string
	}

	cases := []TestCase{
		{subMinutes(1), "1 minute ago"},
		{subMinutes(2), "2 minutes ago"},
		{subMinutes(9), "9 minutes ago"},
		{subMinutes(10), "10 minutes ago"},
		{subMinutes(11), "11 minutes ago"},
		{subMinutes(20), "20 minutes ago"},
		{subMinutes(21), "21 minutes ago"},
		{subMinutes(22), "22 minutes ago"},
		{subMinutes(30), "30 minutes ago"},
		{subMinutes(31), "31 minutes ago"},
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
		{subDays(6), "6 days ago"},
		{subWeeks(1), "1 week ago"},
		{subWeeks(2), "2 weeks ago"},
		{subWeeks(3), "3 weeks ago"},
		{subMonths(1), "1 month ago"},
		{subMonths(2), "2 months ago"},
		{subMonths(3), "3 months ago"},
		{subMonths(4), "4 months ago"},
		{subMonths(5), "5 months ago"},
		{subMonths(6), "6 months ago"},
		{subMonths(7), "7 months ago"},
		{subMonths(8), "8 months ago"},
		{subMonths(9), "9 months ago"},
		{subMonths(10), "10 months ago"},
		{subMonths(11), "11 months ago"},
		{subMonths(12), "1 year ago"},
		{subYears(1), "1 year ago"},
		{subYears(2), "2 years ago"},
		{subYears(21), "21 years ago"},
		{subYears(31), "31 years ago"},
		{subYears(100), "100 years ago"},
	}

	for i := 0; i < 60; i++ {
		cases = append(cases, TestCase{subSeconds(time.Duration(i)), "Online"})
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			timeago.Reconfigure(timeago.Config{Language: langEn})

			res, err := timeago.Parse(tc.date, timeago.OptOnline)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("Result must be %s, but got %s instead", tc.res, res)
			}
		})
	}
}
