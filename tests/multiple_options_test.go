package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
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
		{utils.SubMinutes(1), "1 minute"},
		{utils.SubMinutes(5), "5 minutes"},
		{utils.SubMinutes(10), "10 minutes"},
		{utils.SubMinutes(30), "30 minutes"},
		{utils.SubMinutes(31), "31 minutes"},
		{utils.SubHours(1), "1 hour"},
		{utils.SubHours(11), "11 hours"},
		{utils.SubHours(20), "20 hours"},
		{utils.SubDays(6), "6 days"},
		{utils.SubWeeks(1), "1 week"},
		{utils.SubWeeks(2), "2 weeks"},
		{utils.SubWeeks(3), "3 weeks"},
		{utils.SubMonths(1), "1 month"},
		{utils.SubMonths(6), "6 months"},
		{utils.SubMonths(8), "8 months"},
		{utils.SubMonths(9), "9 months"},
		{utils.SubMonths(11), "11 months"},
		{utils.SubMonths(1), "1 month"},
		{utils.SubYears(2), "2 years"},
		{utils.SubYears(44), "44 years"},
	}

	for i := 0; i < 60; i++ {
		cases = append(cases, TestCase{utils.SubSeconds(time.Duration(i)), "Online"})
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangEn})

			res, err := ago.Parse(tc.date, ago.OptOnline, ago.OptNoSuffix)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("Result must be %s, but got %s instead", tc.res, res)
			}
		})
	}
}
