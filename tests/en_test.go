package tests

import (
	"testing"
	"time"

	. "github.com/SerhiiCho/timeago"
	. "github.com/SerhiiCho/timeago/utils"
)

func TestParseEn(t *testing.T) {
	cases := []struct {
		date   string
		result string
	}{
		{SmallSubTime(-60 * time.Second), "1 minute ago"},
		{SmallSubTime(-1 * time.Minute), "1 minute ago"},
		{SmallSubTime(-2 * time.Minute), "2 minutes ago"},
		{SmallSubTime(-5 * time.Minute), "5 minutes ago"},
		{SmallSubTime(-9 * time.Minute), "9 minutes ago"},
		{SmallSubTime(-10 * time.Minute), "10 minutes ago"},
		{SmallSubTime(-11 * time.Minute), "11 minutes ago"},
		{SmallSubTime(-20 * time.Minute), "20 minutes ago"},
		{SmallSubTime(-21 * time.Minute), "21 minutes ago"},
		{SmallSubTime(-22 * time.Minute), "22 minutes ago"},
		{SmallSubTime(-30 * time.Minute), "30 minutes ago"},
		{SmallSubTime(-31 * time.Minute), "31 minutes ago"},
		{SmallSubTime(-59 * time.Minute), "59 minutes ago"},
		{SmallSubTime(-60 * time.Minute), "1 hour ago"},
		{SmallSubTime(-1 * time.Hour), "1 hour ago"},
		{SmallSubTime(-2 * time.Hour), "2 hours ago"},
		{SmallSubTime(-9 * time.Hour), "9 hours ago"},
		{SmallSubTime(-10 * time.Hour), "10 hours ago"},
		{SmallSubTime(-11 * time.Hour), "11 hours ago"},
		{SmallSubTime(-20 * time.Hour), "20 hours ago"},
		{SmallSubTime(-21 * time.Hour), "21 hours ago"},
		{SmallSubTime(-23 * time.Hour), "23 hours ago"},
		{SmallSubTime(-24 * time.Hour), "1 day ago"},
		{SmallSubTime(-30 * time.Hour), "1 day ago"},
		{SmallSubTime((-24 * 2) * time.Hour), "2 days ago"},
		{SmallSubTime((-24 * 6) * time.Hour), "6 days ago"},
		{SmallSubTime((-24 * 7) * time.Hour), "1 week ago"},
		{SmallSubTime((-24 * 14) * time.Hour), "2 weeks ago"},
		{SmallSubTime((-24 * 21) * time.Hour), "3 weeks ago"},
		{BigSubTime(0, 1, 1), "1 month ago"},
		{BigSubTime(0, 2, 1), "2 months ago"},
		{BigSubTime(0, 9, 1), "9 months ago"},
		{BigSubTime(0, 11, 1), "11 months ago"},
		{BigSubTime(0, 12, 1), "1 year ago"},
		{BigSubTime(1, 0, 1), "1 year ago"},
		{BigSubTime(2, 0, 1), "2 years ago"},
		{BigSubTime(21, 0, 1), "21 years ago"},
		{BigSubTime(31, 0, 1), "31 years ago"},
		{BigSubTime(100, 0, 1), "100 years ago"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			SetConfig(Config{
				Language: "en",
				Location: "Europe/Kiev",
			})

			if res := Parse(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseEnWithOnlineFlag(t *testing.T) {
	cases := []struct {
		date   string
		result string
	}{
		{SmallSubTime(time.Second * 2), "Online"},
		{SmallSubTime(time.Second), "Online"},
		{SmallSubTime(-1 * time.Second), "Online"},
		{SmallSubTime(-2 * time.Second), "Online"},
		{SmallSubTime(-9 * time.Second), "Online"},
		{SmallSubTime(-10 * time.Second), "Online"},
		{SmallSubTime(-11 * time.Second), "Online"},
		{SmallSubTime(-20 * time.Second), "Online"},
		{SmallSubTime(-21 * time.Second), "Online"},
		{SmallSubTime(-22 * time.Second), "Online"},
		{SmallSubTime(-30 * time.Second), "Online"},
		{SmallSubTime(-31 * time.Second), "Online"},
		{SmallSubTime(-60 * time.Second), "1 minute ago"},
		{SmallSubTime(-1 * time.Minute), "1 minute ago"},
		{SmallSubTime(-2 * time.Minute), "2 minutes ago"},
		{SmallSubTime(-9 * time.Minute), "9 minutes ago"},
		{SmallSubTime(-10 * time.Minute), "10 minutes ago"},
		{SmallSubTime(-11 * time.Minute), "11 minutes ago"},
		{SmallSubTime(-20 * time.Minute), "20 minutes ago"},
		{SmallSubTime(-21 * time.Minute), "21 minutes ago"},
		{SmallSubTime(-22 * time.Minute), "22 minutes ago"},
		{SmallSubTime(-30 * time.Minute), "30 minutes ago"},
		{SmallSubTime(-31 * time.Minute), "31 minutes ago"},
		{SmallSubTime(-60 * time.Minute), "1 hour ago"},
		{SmallSubTime(-1 * time.Hour), "1 hour ago"},
		{SmallSubTime(-2 * time.Hour), "2 hours ago"},
		{SmallSubTime(-9 * time.Hour), "9 hours ago"},
		{SmallSubTime(-10 * time.Hour), "10 hours ago"},
		{SmallSubTime(-11 * time.Hour), "11 hours ago"},
		{SmallSubTime(-20 * time.Hour), "20 hours ago"},
		{SmallSubTime(-21 * time.Hour), "21 hours ago"},
		{SmallSubTime(-23 * time.Hour), "23 hours ago"},
		{SmallSubTime(-24 * time.Hour), "1 day ago"},
		{SmallSubTime(-30 * time.Hour), "1 day ago"},
		{SmallSubTime((-24 * 2) * time.Hour), "2 days ago"},
		{SmallSubTime((-24 * 6) * time.Hour), "6 days ago"},
		{SmallSubTime((-24 * 7) * time.Hour), "1 week ago"},
		{SmallSubTime((-24 * 14) * time.Hour), "2 weeks ago"},
		{SmallSubTime((-24 * 21) * time.Hour), "3 weeks ago"},
		{BigSubTime(0, 1, 1), "1 month ago"},
		{BigSubTime(0, 2, 1), "2 months ago"},
		{BigSubTime(0, 9, 1), "9 months ago"},
		{BigSubTime(0, 11, 1), "11 months ago"},
		{BigSubTime(0, 12, 1), "1 year ago"},
		{BigSubTime(1, 0, 1), "1 year ago"},
		{BigSubTime(2, 0, 1), "2 years ago"},
		{BigSubTime(21, 0, 1), "21 years ago"},
		{BigSubTime(31, 0, 1), "31 years ago"},
		{BigSubTime(100, 0, 1), "100 years ago"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {

			SetConfig(Config{
				Language: "en",
			})

			if res := Parse(tc.date, "online"); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseEnWithSeconds(t *testing.T) {
	cases := []struct {
		date   string
		result []string
	}{
		{SmallSubTime(time.Second * 2), []string{"0 seconds ago", "1 second ago"}},
		{SmallSubTime(time.Second), []string{"0 seconds ago", "1 second ago"}},
		{SmallSubTime(-1 * time.Second), []string{"1 second ago", "2 seconds ago"}},
		{SmallSubTime(-2 * time.Second), []string{"2 seconds ago", "3 seconds ago"}},
		{SmallSubTime(-9 * time.Second), []string{"9 seconds ago", "10 seconds ago"}},
		{SmallSubTime(-10 * time.Second), []string{"10 seconds ago", "11 seconds ago"}},
		{SmallSubTime(-11 * time.Second), []string{"11 seconds ago", "12 seconds ago"}},
		{SmallSubTime(-20 * time.Second), []string{"20 seconds ago", "21 seconds ago"}},
		{SmallSubTime(-21 * time.Second), []string{"21 seconds ago", "22 seconds ago"}},
		{SmallSubTime(-22 * time.Second), []string{"22 seconds ago", "23 seconds ago"}},
		{SmallSubTime(-30 * time.Second), []string{"30 seconds ago", "31 seconds ago"}},
		{SmallSubTime(-31 * time.Second), []string{"31 seconds ago", "32 seconds ago"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			SetConfig(Config{
				Language: "en",
				Location: "Europe/Kiev",
			})

			if res := Parse(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
