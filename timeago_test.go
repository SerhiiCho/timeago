package timeago

import (
	"testing"
	"time"
)

func TestGetWords(t *testing.T) {
	cases := []struct {
		timeKind string
		num      int
		result   string
		lang     string
	}{
		// english
		{"days", 11, "11 days ago", "en"},
		{"days", 21, "21 days ago", "en"},
		{"seconds", 30, "30 seconds ago", "en"},
		{"seconds", 31, "31 seconds ago", "en"},
		{"hours", 10, "10 hours ago", "en"},
		{"years", 2, "2 years ago", "en"},
		// russian
		{"hours", 5, "5 часов назад", "ru"},
		{"days", 11, "11 дней назад", "ru"},
		{"years", 21, "21 год назад", "ru"},
		{"minutes", 59, "59 минут назад", "ru"},
	}

	for _, tc := range cases {
		t.Run(tc.result, func(test *testing.T) {
			SetConfig(Config{
				Language: tc.lang,
			})

			if res := getWords(tc.timeKind, tc.num); res != tc.result {
				test.Errorf("Result must be `%s` but got `%s` instead", tc.result, res)
			}
		})
	}
}

func TestParseFunctionCanExceptTimestamp(t *testing.T) {
	cases := []struct {
		timestamp int
		result    string
	}{
		{getTimestampOfPastDate(time.Minute), "1 minute ago"},
		{getTimestampOfPastDate(time.Minute * 5), "5 minutes ago"},
		{getTimestampOfPastDate(time.Hour), "1 hour ago"},
		{getTimestampOfPastDate(time.Hour * 3), "3 hours ago"},
		{getTimestampOfPastDate(time.Hour * 24), "1 day ago"},
		{getTimestampOfPastDate(time.Hour * 48), "2 days ago"},
	}

	SetConfig(Config{
		Language: "en",
	})

	for _, tc := range cases {
		t.Run(tc.result, func(test *testing.T) {
			if res := Parse(tc.timestamp); res != tc.result {
				test.Errorf("Result must be %v, but got %v instead", tc.result, res)
			}
		})
	}
}

func TestParseFunctionCanExceptTimePackage(t *testing.T) {
	cases := []struct {
		time   time.Time
		result string
	}{
		{time.Now().Add(-time.Minute), "1 minute ago"},
		{time.Now().Add(-time.Minute * 2), "2 minutes ago"},
		{time.Now().Add(-time.Minute * 3), "3 minutes ago"},
		{time.Now().Add(-time.Minute * 4), "4 minutes ago"},
		{time.Now().Add(-time.Minute * 5), "5 minutes ago"},
		{time.Now().Add(-time.Minute * 6), "6 minutes ago"},
		{time.Now().Add(-time.Hour * 7), "7 hours ago"},
		{time.Now().Add(-time.Hour * 8), "8 hours ago"},
		{time.Now().Add(-time.Hour * 9), "9 hours ago"},
		{time.Now().Add(-time.Hour * 10), "10 hours ago"},
		{time.Now().Add(-time.Hour * 11), "11 hours ago"},
	}

	SetConfig(Config{
		Language: "en",
	})

	for _, tc := range cases {
		t.Run("Test for date "+tc.time.String(), func(test *testing.T) {
			if res := Parse(tc.time); res != tc.result {
				test.Errorf("Result must be %v, but got %v instead", tc.result, res)
			}
		})
	}
}

func TestParseFuncWillCalculateIntervalToFutureDate(t *testing.T) {
	testCases := []struct {
		time   time.Time
		result string
	}{
		{time.Now().Add(time.Minute * 2), "2 minutes"},
		{time.Now().Add(time.Minute * 5), "5 minutes"},
		{time.Now().Add(time.Minute * 10), "10 minutes"},
		{time.Now().Add(time.Hour), "1 hour"},
		{time.Now().Add(time.Hour * 24), "1 day"},
		{time.Now().Add(time.Hour * 48), "2 days"},
	}

	SetConfig(Config{
		Language: "en",
	})

	for _, tc := range testCases {
		t.Run("Test for date: "+tc.time.String(), func(test *testing.T) {
			if res := Parse(tc.time); res != tc.result {
				test.Errorf("Result must be %v, but got %v instead", tc.result, res)
			}
		})
	}
}

func TestOptionIsEnabled(t *testing.T) {
	t.Run("returns true if option is enabled", func(test *testing.T) {
		globalOptions = []string{"noSuffix"}

		if res := optionIsEnabled("noSuffix"); res == false {
			test.Errorf("Result must be true, but got %v instead", res)
		}

		globalOptions = []string{}
	})

	t.Run("returns false if option is disabled", func(test *testing.T) {
		if res := optionIsEnabled("noSuffix"); res == true {
			test.Errorf("Result must be true, but got %v instead", res)
		}
	})
}
