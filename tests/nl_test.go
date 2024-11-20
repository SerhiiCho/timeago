package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago/v3"
)

const langNl = "nl"

func TestParseNl(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{subMinutes(1), "1 minuut geleden"},
		{subMinutes(2), "2 minuten geleden"},
		{subMinutes(5), "5 minuten geleden"},
		{subMinutes(9), "9 minuten geleden"},
		{subMinutes(10), "10 minuten geleden"},
		{subMinutes(11), "11 minuten geleden"},
		{subMinutes(20), "20 minuten geleden"},
		{subMinutes(21), "21 minuten geleden"},
		{subMinutes(22), "22 minuten geleden"},
		{subMinutes(30), "30 minuten geleden"},
		{subMinutes(31), "31 minuten geleden"},
		{subMinutes(59), "59 minuten geleden"},
		{subMinutes(60), "1 uur geleden"},
		{subHours(1), "1 uur geleden"},
		{subHours(2), "2 uur geleden"},
		{subHours(9), "9 uur geleden"},
		{subHours(10), "10 uur geleden"},
		{subHours(11), "11 uur geleden"},
		{subHours(20), "20 uur geleden"},
		{subHours(21), "21 uur geleden"},
		{subHours(23), "23 uur geleden"},
		{subDays(1), "1 dag geleden"},
		{subDays(2), "2 dagen geleden"},
		{subDays(6), "6 dagen geleden"},
		{subWeeks(1), "1 week geleden"},
		{subWeeks(2), "2 weken geleden"},
		{subWeeks(3), "3 weken geleden"},
		{subMonths(1), "1 maand geleden"},
		{subMonths(2), "2 maanden geleden"},
		{subMonths(3), "3 maanden geleden"},
		{subMonths(4), "4 maanden geleden"},
		{subMonths(5), "5 maanden geleden"},
		{subMonths(6), "6 maanden geleden"},
		{subMonths(7), "7 maanden geleden"},
		{subMonths(8), "8 maanden geleden"},
		{subMonths(9), "9 maanden geleden"},
		{subMonths(10), "10 maanden geleden"},
		{subMonths(11), "11 maanden geleden"},
		{subYears(1), "1 jaar geleden"},
		{subYears(2), "2 jaar geleden"},
		{subYears(21), "21 jaar geleden"},
		{subYears(31), "31 jaar geleden"},
		{subYears(100), "100 jaar geleden"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			timeago.Reconfigure(timeago.Config{Language: langNl})

			res, err := timeago.Parse(tc.date)

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
		{subSeconds(0), []string{"0 seconden geleden", "1 seconde geleden"}},
		{subSeconds(1), []string{"1 seconde geleden", "2 seconden geleden"}},
		{subSeconds(2), []string{"2 seconden geleden", "3 seconden geleden"}},
		{subSeconds(9), []string{"9 seconden geleden", "10 seconden geleden"}},
		{subSeconds(10), []string{"10 seconden geleden", "11 seconden geleden"}},
		{subSeconds(11), []string{"11 seconden geleden", "12 seconden geleden"}},
		{subSeconds(20), []string{"20 seconden geleden", "21 seconden geleden"}},
		{subSeconds(21), []string{"21 seconden geleden", "22 seconden geleden"}},
		{subSeconds(22), []string{"22 seconden geleden", "23 seconden geleden"}},
		{subSeconds(30), []string{"30 seconden geleden", "31 seconden geleden"}},
		{subSeconds(31), []string{"31 seconden geleden", "32 seconden geleden"}},
		{subSeconds(59), []string{"59 seconden geleden", "1 minuten geleden"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			timeago.Reconfigure(timeago.Config{Language: langNl})

			res, err := timeago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res[0] && res != tc.res[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.res[0], tc.res[1], res)
			}
		})
	}
}
