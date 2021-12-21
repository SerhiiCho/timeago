package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago"
)

func TestTakeEn(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{smallSubTime(-60 * time.Second), "1 minute ago", "en"},
		{smallSubTime(-1 * time.Minute), "1 minute ago", "en"},
		{smallSubTime(-2 * time.Minute), "2 minutes ago", "en"},
		{smallSubTime(-5 * time.Minute), "5 minutes ago", "en"},
		{smallSubTime(-9 * time.Minute), "9 minutes ago", "en"},
		{smallSubTime(-10 * time.Minute), "10 minutes ago", "en"},
		{smallSubTime(-11 * time.Minute), "11 minutes ago", "en"},
		{smallSubTime(-20 * time.Minute), "20 minutes ago", "en"},
		{smallSubTime(-21 * time.Minute), "21 minutes ago", "en"},
		{smallSubTime(-22 * time.Minute), "22 minutes ago", "en"},
		{smallSubTime(-30 * time.Minute), "30 minutes ago", "en"},
		{smallSubTime(-31 * time.Minute), "31 minutes ago", "en"},
		{smallSubTime(-59 * time.Minute), "59 minutes ago", "en"},
		{smallSubTime(-60 * time.Minute), "1 hour ago", "en"},
		{smallSubTime(-1 * time.Hour), "1 hour ago", "en"},
		{smallSubTime(-2 * time.Hour), "2 hours ago", "en"},
		{smallSubTime(-9 * time.Hour), "9 hours ago", "en"},
		{smallSubTime(-10 * time.Hour), "10 hours ago", "en"},
		{smallSubTime(-11 * time.Hour), "11 hours ago", "en"},
		{smallSubTime(-20 * time.Hour), "20 hours ago", "en"},
		{smallSubTime(-21 * time.Hour), "21 hours ago", "en"},
		{smallSubTime(-23 * time.Hour), "23 hours ago", "en"},
		{smallSubTime(-24 * time.Hour), "1 day ago", "en"},
		{smallSubTime(-30 * time.Hour), "1 day ago", "en"},
		{smallSubTime((-24 * 2) * time.Hour), "2 days ago", "en"},
		{smallSubTime((-24 * 6) * time.Hour), "6 days ago", "en"},
		{smallSubTime((-24 * 7) * time.Hour), "1 week ago", "en"},
		{smallSubTime((-24 * 14) * time.Hour), "2 weeks ago", "en"},
		{smallSubTime((-24 * 21) * time.Hour), "3 weeks ago", "en"},
		{bigSubTime(0, 1, 1), "1 month ago", "en"},
		{bigSubTime(0, 2, 1), "2 months ago", "en"},
		{bigSubTime(0, 9, 1), "9 months ago", "en"},
		{bigSubTime(0, 11, 1), "11 months ago", "en"},
		{bigSubTime(0, 12, 1), "1 year ago", "en"},
		{bigSubTime(1, 0, 1), "1 year ago", "en"},
		{bigSubTime(2, 0, 1), "2 years ago", "en"},
		{bigSubTime(21, 0, 1), "21 years ago", "en"},
		{bigSubTime(31, 0, 1), "31 years ago", "en"},
		{bigSubTime(100, 0, 1), "100 years ago", "en"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)
			timeago.Set("location", "Europe/Kiev")

			if res := timeago.Take(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestTakeEnWithOnlineFlag(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{smallSubTime(time.Second * 2), "Online", "en"},
		{smallSubTime(time.Second), "Online", "en"},
		{smallSubTime(-1 * time.Second), "Online", "en"},
		{smallSubTime(-2 * time.Second), "Online", "en"},
		{smallSubTime(-9 * time.Second), "Online", "en"},
		{smallSubTime(-10 * time.Second), "Online", "en"},
		{smallSubTime(-11 * time.Second), "Online", "en"},
		{smallSubTime(-20 * time.Second), "Online", "en"},
		{smallSubTime(-21 * time.Second), "Online", "en"},
		{smallSubTime(-22 * time.Second), "Online", "en"},
		{smallSubTime(-30 * time.Second), "Online", "en"},
		{smallSubTime(-31 * time.Second), "Online", "en"},
		{smallSubTime(-60 * time.Second), "1 minute ago", "en"},
		{smallSubTime(-1 * time.Minute), "1 minute ago", "en"},
		{smallSubTime(-2 * time.Minute), "2 minutes ago", "en"},
		{smallSubTime(-9 * time.Minute), "9 minutes ago", "en"},
		{smallSubTime(-10 * time.Minute), "10 minutes ago", "en"},
		{smallSubTime(-11 * time.Minute), "11 minutes ago", "en"},
		{smallSubTime(-20 * time.Minute), "20 minutes ago", "en"},
		{smallSubTime(-21 * time.Minute), "21 minutes ago", "en"},
		{smallSubTime(-22 * time.Minute), "22 minutes ago", "en"},
		{smallSubTime(-30 * time.Minute), "30 minutes ago", "en"},
		{smallSubTime(-31 * time.Minute), "31 minutes ago", "en"},
		{smallSubTime(-60 * time.Minute), "1 hour ago", "en"},
		{smallSubTime(-1 * time.Hour), "1 hour ago", "en"},
		{smallSubTime(-2 * time.Hour), "2 hours ago", "en"},
		{smallSubTime(-9 * time.Hour), "9 hours ago", "en"},
		{smallSubTime(-10 * time.Hour), "10 hours ago", "en"},
		{smallSubTime(-11 * time.Hour), "11 hours ago", "en"},
		{smallSubTime(-20 * time.Hour), "20 hours ago", "en"},
		{smallSubTime(-21 * time.Hour), "21 hours ago", "en"},
		{smallSubTime(-23 * time.Hour), "23 hours ago", "en"},
		{smallSubTime(-24 * time.Hour), "1 day ago", "en"},
		{smallSubTime(-30 * time.Hour), "1 day ago", "en"},
		{smallSubTime((-24 * 2) * time.Hour), "2 days ago", "en"},
		{smallSubTime((-24 * 6) * time.Hour), "6 days ago", "en"},
		{smallSubTime((-24 * 7) * time.Hour), "1 week ago", "en"},
		{smallSubTime((-24 * 14) * time.Hour), "2 weeks ago", "en"},
		{smallSubTime((-24 * 21) * time.Hour), "3 weeks ago", "en"},
		{bigSubTime(0, 1, 1), "1 month ago", "en"},
		{bigSubTime(0, 2, 1), "2 months ago", "en"},
		{bigSubTime(0, 9, 1), "9 months ago", "en"},
		{bigSubTime(0, 11, 1), "11 months ago", "en"},
		{bigSubTime(0, 12, 1), "1 year ago", "en"},
		{bigSubTime(1, 0, 1), "1 year ago", "en"},
		{bigSubTime(2, 0, 1), "2 years ago", "en"},
		{bigSubTime(21, 0, 1), "21 years ago", "en"},
		{bigSubTime(31, 0, 1), "31 years ago", "en"},
		{bigSubTime(100, 0, 1), "100 years ago", "en"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date+"|online", func(test *testing.T) {
			timeago.Set("language", tc.lang)

			if res := timeago.Take(tc.date + "|online"); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestTakeEnWithSeconds(t *testing.T) {
	cases := []struct {
		date   string
		result []string
		lang   string
	}{
		{smallSubTime(time.Second * 2), []string{"0 seconds ago", "1 second ago"}, "en"},
		{smallSubTime(time.Second), []string{"0 seconds ago", "1 second ago"}, "en"},
		{smallSubTime(-1 * time.Second), []string{"1 second ago", "2 seconds ago"}, "en"},
		{smallSubTime(-2 * time.Second), []string{"2 seconds ago", "3 seconds ago"}, "en"},
		{smallSubTime(-9 * time.Second), []string{"9 seconds ago", "10 seconds ago"}, "en"},
		{smallSubTime(-10 * time.Second), []string{"10 seconds ago", "11 seconds ago"}, "en"},
		{smallSubTime(-11 * time.Second), []string{"11 seconds ago", "12 seconds ago"}, "en"},
		{smallSubTime(-20 * time.Second), []string{"20 seconds ago", "21 seconds ago"}, "en"},
		{smallSubTime(-21 * time.Second), []string{"21 seconds ago", "22 seconds ago"}, "en"},
		{smallSubTime(-22 * time.Second), []string{"22 seconds ago", "23 seconds ago"}, "en"},
		{smallSubTime(-30 * time.Second), []string{"30 seconds ago", "31 seconds ago"}, "en"},
		{smallSubTime(-31 * time.Second), []string{"31 seconds ago", "32 seconds ago"}, "en"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)
			timeago.Set("location", "Europe/Kiev")

			if res := timeago.Take(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
