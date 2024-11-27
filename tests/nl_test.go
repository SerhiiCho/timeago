package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseNl(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{utils.SubMinutes(1), "1 minuut geleden"},
		{utils.SubMinutes(2), "2 minuten geleden"},
		{utils.SubMinutes(5), "5 minuten geleden"},
		{utils.SubMinutes(9), "9 minuten geleden"},
		{utils.SubMinutes(10), "10 minuten geleden"},
		{utils.SubMinutes(11), "11 minuten geleden"},
		{utils.SubMinutes(20), "20 minuten geleden"},
		{utils.SubMinutes(21), "21 minuten geleden"},
		{utils.SubMinutes(22), "22 minuten geleden"},
		{utils.SubMinutes(30), "30 minuten geleden"},
		{utils.SubMinutes(31), "31 minuten geleden"},
		{utils.SubMinutes(59), "59 minuten geleden"},
		{utils.SubMinutes(60), "1 uur geleden"},
		{utils.SubHours(1), "1 uur geleden"},
		{utils.SubHours(2), "2 uur geleden"},
		{utils.SubHours(9), "9 uur geleden"},
		{utils.SubHours(10), "10 uur geleden"},
		{utils.SubHours(11), "11 uur geleden"},
		{utils.SubHours(20), "20 uur geleden"},
		{utils.SubHours(21), "21 uur geleden"},
		{utils.SubHours(23), "23 uur geleden"},
		{utils.SubDays(1), "1 dag geleden"},
		{utils.SubDays(2), "2 dagen geleden"},
		{utils.SubDays(6), "6 dagen geleden"},
		{utils.SubWeeks(1), "1 week geleden"},
		{utils.SubWeeks(2), "2 weken geleden"},
		{utils.SubWeeks(3), "3 weken geleden"},
		{utils.SubMonths(1), "1 maand geleden"},
		{utils.SubMonths(2), "2 maanden geleden"},
		{utils.SubMonths(3), "3 maanden geleden"},
		{utils.SubMonths(4), "4 maanden geleden"},
		{utils.SubMonths(5), "5 maanden geleden"},
		{utils.SubMonths(6), "6 maanden geleden"},
		{utils.SubMonths(7), "7 maanden geleden"},
		{utils.SubMonths(8), "8 maanden geleden"},
		{utils.SubMonths(9), "9 maanden geleden"},
		{utils.SubMonths(10), "10 maanden geleden"},
		{utils.SubMonths(11), "11 maanden geleden"},
		{utils.SubYears(1), "1 jaar geleden"},
		{utils.SubYears(2), "2 jaar geleden"},
		{utils.SubYears(21), "21 jaar geleden"},
		{utils.SubYears(31), "31 jaar geleden"},
		{utils.SubYears(100), "100 jaar geleden"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangNl})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("Result must be %s, but got %s instead", tc.res, res)
			}
		})
	}
}

func TestParseNlWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{utils.SubSeconds(0), []string{"0 seconden geleden", "1 seconde geleden"}},
		{utils.SubSeconds(1), []string{"1 seconde geleden", "2 seconden geleden"}},
		{utils.SubSeconds(2), []string{"2 seconden geleden", "3 seconden geleden"}},
		{utils.SubSeconds(9), []string{"9 seconden geleden", "10 seconden geleden"}},
		{utils.SubSeconds(10), []string{"10 seconden geleden", "11 seconden geleden"}},
		{utils.SubSeconds(11), []string{"11 seconden geleden", "12 seconden geleden"}},
		{utils.SubSeconds(20), []string{"20 seconden geleden", "21 seconden geleden"}},
		{utils.SubSeconds(21), []string{"21 seconden geleden", "22 seconden geleden"}},
		{utils.SubSeconds(22), []string{"22 seconden geleden", "23 seconden geleden"}},
		{utils.SubSeconds(30), []string{"30 seconden geleden", "31 seconden geleden"}},
		{utils.SubSeconds(31), []string{"31 seconden geleden", "32 seconden geleden"}},
		{utils.SubSeconds(59), []string{"59 seconden geleden", "1 minuten geleden"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangNl})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res[0] && res != tc.res[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.res[0], tc.res[1], res)
			}
		})
	}
}
