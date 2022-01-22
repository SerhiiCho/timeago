package tests

import (
	"testing"
	"time"

	. "github.com/SerhiiCho/timeago"
)

func TestParseNl(t *testing.T) {
	cases := []struct {
		date   string
		result string
	}{
		{smallSubTime(-60 * time.Second), "1 minuut geleden"},
		{smallSubTime(-1 * time.Minute), "1 minuut geleden"},
		{smallSubTime(-2 * time.Minute), "2 minuten geleden"},
		{smallSubTime(-5 * time.Minute), "5 minuten geleden"},
		{smallSubTime(-9 * time.Minute), "9 minuten geleden"},
		{smallSubTime(-10 * time.Minute), "10 minuten geleden"},
		{smallSubTime(-11 * time.Minute), "11 minuten geleden"},
		{smallSubTime(-20 * time.Minute), "20 minuten geleden"},
		{smallSubTime(-21 * time.Minute), "21 minuten geleden"},
		{smallSubTime(-22 * time.Minute), "22 minuten geleden"},
		{smallSubTime(-30 * time.Minute), "30 minuten geleden"},
		{smallSubTime(-31 * time.Minute), "31 minuten geleden"},
		{smallSubTime(-59 * time.Minute), "59 minuten geleden"},
		{smallSubTime(-60 * time.Minute), "1 uur geleden"},
		{smallSubTime(-1 * time.Hour), "1 uur geleden"},
		{smallSubTime(-2 * time.Hour), "2 uur geleden"},
		{smallSubTime(-9 * time.Hour), "9 uur geleden"},
		{smallSubTime(-10 * time.Hour), "10 uur geleden"},
		{smallSubTime(-11 * time.Hour), "11 uur geleden"},
		{smallSubTime(-20 * time.Hour), "20 uur geleden"},
		{smallSubTime(-21 * time.Hour), "21 uur geleden"},
		{smallSubTime(-23 * time.Hour), "23 uur geleden"},
		{smallSubTime(-24 * time.Hour), "1 dag geleden"},
		{smallSubTime(-30 * time.Hour), "1 dag geleden"},
		{smallSubTime((-24 * 2) * time.Hour), "2 dagen geleden"},
		{smallSubTime((-24 * 6) * time.Hour), "6 dagen geleden"},
		{smallSubTime((-24 * 7) * time.Hour), "1 week geleden"},
		{smallSubTime((-24 * 14) * time.Hour), "2 weken geleden"},
		{smallSubTime((-24 * 21) * time.Hour), "3 weken geleden"},
		{bigSubTime(0, 1, 1), "1 maand geleden"},
		{bigSubTime(0, 2, 1), "2 maanden geleden"},
		{bigSubTime(0, 9, 1), "9 maanden geleden"},
		{bigSubTime(0, 11, 1), "11 maanden geleden"},
		{bigSubTime(0, 12, 1), "1 jaar geleden"},
		{bigSubTime(1, 0, 1), "1 jaar geleden"},
		{bigSubTime(2, 0, 1), "2 jaar geleden"},
		{bigSubTime(21, 0, 1), "21 jaar geleden"},
		{bigSubTime(31, 0, 1), "31 jaar geleden"},
		{bigSubTime(100, 0, 1), "100 jaar geleden"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			SetConfig(Config{
				Language: "nl",
			})

			if res := Parse(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseNlWithOnlineFlag(t *testing.T) {
	cases := []struct {
		date   string
		result string
	}{
		{smallSubTime(time.Second * 2), "Online"},
		{smallSubTime(time.Second), "Online"},
		{smallSubTime(-1 * time.Second), "Online"},
		{smallSubTime(-2 * time.Second), "Online"},
		{smallSubTime(-9 * time.Second), "Online"},
		{smallSubTime(-10 * time.Second), "Online"},
		{smallSubTime(-11 * time.Second), "Online"},
		{smallSubTime(-20 * time.Second), "Online"},
		{smallSubTime(-21 * time.Second), "Online"},
		{smallSubTime(-22 * time.Second), "Online"},
		{smallSubTime(-30 * time.Second), "Online"},
		{smallSubTime(-31 * time.Second), "Online"},
		{smallSubTime(-60 * time.Second), "1 minuut geleden"},
		{smallSubTime(-1 * time.Minute), "1 minuut geleden"},
		{smallSubTime(-2 * time.Minute), "2 minuten geleden"},
		{smallSubTime(-9 * time.Minute), "9 minuten geleden"},
		{smallSubTime(-10 * time.Minute), "10 minuten geleden"},
		{smallSubTime(-11 * time.Minute), "11 minuten geleden"},
		{smallSubTime(-20 * time.Minute), "20 minuten geleden"},
		{smallSubTime(-21 * time.Minute), "21 minuten geleden"},
		{smallSubTime(-22 * time.Minute), "22 minuten geleden"},
		{smallSubTime(-30 * time.Minute), "30 minuten geleden"},
		{smallSubTime(-31 * time.Minute), "31 minuten geleden"},
		{smallSubTime(-60 * time.Minute), "1 uur geleden"},
		{smallSubTime(-1 * time.Hour), "1 uur geleden"},
		{smallSubTime(-2 * time.Hour), "2 uur geleden"},
		{smallSubTime(-9 * time.Hour), "9 uur geleden"},
		{smallSubTime(-10 * time.Hour), "10 uur geleden"},
		{smallSubTime(-11 * time.Hour), "11 uur geleden"},
		{smallSubTime(-20 * time.Hour), "20 uur geleden"},
		{smallSubTime(-21 * time.Hour), "21 uur geleden"},
		{smallSubTime(-23 * time.Hour), "23 uur geleden"},
		{smallSubTime(-24 * time.Hour), "1 dag geleden"},
		{smallSubTime(-30 * time.Hour), "1 dag geleden"},
		{smallSubTime((-24 * 2) * time.Hour), "2 dagen geleden"},
		{smallSubTime((-24 * 6) * time.Hour), "6 dagen geleden"},
		{smallSubTime((-24 * 7) * time.Hour), "1 week geleden"},
		{smallSubTime((-24 * 14) * time.Hour), "2 weken geleden"},
		{smallSubTime((-24 * 21) * time.Hour), "3 weken geleden"},
		{bigSubTime(0, 1, 1), "1 maand geleden"},
		{bigSubTime(0, 2, 1), "2 maanden geleden"},
		{bigSubTime(0, 9, 1), "9 maanden geleden"},
		{bigSubTime(0, 11, 1), "11 maanden geleden"},
		{bigSubTime(0, 12, 1), "1 jaar geleden"},
		{bigSubTime(1, 0, 1), "1 jaar geleden"},
		{bigSubTime(2, 0, 1), "2 jaar geleden"},
		{bigSubTime(21, 0, 1), "21 jaar geleden"},
		{bigSubTime(31, 0, 1), "31 jaar geleden"},
		{bigSubTime(100, 0, 1), "100 jaar geleden"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {

			SetConfig(Config{
				Language: "nl",
			})

			if res := Parse(tc.date, "online"); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseNlWithSeconds(t *testing.T) {
	cases := []struct {
		date   string
		result []string
	}{
		{smallSubTime(time.Second * 2), []string{"0 seconden geleden", "1 seconde geleden"}},
		{smallSubTime(time.Second), []string{"0 seconden geleden", "1 seconde geleden"}},
		{smallSubTime(-1 * time.Second), []string{"1 seconde geleden", "2 seconden geleden"}},
		{smallSubTime(-2 * time.Second), []string{"2 seconden geleden", "3 seconden geleden"}},
		{smallSubTime(-9 * time.Second), []string{"9 seconden geleden", "10 seconden geleden"}},
		{smallSubTime(-10 * time.Second), []string{"10 seconden geleden", "11 seconden geleden"}},
		{smallSubTime(-11 * time.Second), []string{"11 seconden geleden", "12 seconden geleden"}},
		{smallSubTime(-20 * time.Second), []string{"20 seconden geleden", "21 seconden geleden"}},
		{smallSubTime(-21 * time.Second), []string{"21 seconden geleden", "22 seconden geleden"}},
		{smallSubTime(-22 * time.Second), []string{"22 seconden geleden", "23 seconden geleden"}},
		{smallSubTime(-30 * time.Second), []string{"30 seconden geleden", "31 seconden geleden"}},
		{smallSubTime(-31 * time.Second), []string{"31 seconden geleden", "32 seconden geleden"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			SetConfig(Config{
				Language: "nl",
			})

			if res := Parse(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
