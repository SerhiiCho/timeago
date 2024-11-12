package timeago

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago/v3/config"
	"github.com/SerhiiCho/timeago/v3/option"
)

func TestParseFunctionCanExceptTimestamp(t *testing.T) {
	cases := []struct {
		timestamp int
		result    string
	}{
		{getTimestampOfPastDate(time.Minute), "1 minute ago"},
		// {getTimestampOfPastDate(time.Minute * 5), "5 minutes ago"},
		// {getTimestampOfPastDate(time.Hour), "1 hour ago"},
		// {getTimestampOfPastDate(time.Hour * 3), "3 hours ago"},
		// {getTimestampOfPastDate(time.Hour * 5), "5 hours ago"},
		// {getTimestampOfPastDate(time.Hour * 24), "1 day ago"},
		// {getTimestampOfPastDate(time.Hour * 24 * 2), "2 days ago"},
		// {getTimestampOfPastDate(time.Hour * 24 * 3), "3 days ago"},
		// {getTimestampOfPastDate(time.Hour * 24 * 4), "4 days ago"},
		// {getTimestampOfPastDate(time.Hour * 24 * 5), "5 days ago"},
		// {getTimestampOfPastDate(time.Hour * 24 * 6), "6 days ago"},
		// {getTimestampOfPastDate(time.Hour * 24 * 7), "1 week ago"},
	}

	Configure(&config.Config{Language: "en"})

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

	Configure(&config.Config{Language: "en"})

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

	Configure(&config.Config{Language: "en"})

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
		options = []option.Option{"noSuffix"}

		if res := optionIsEnabled("noSuffix"); res == false {
			test.Error("Result must be true, but got false instead")
		}

		options = []option.Option{}
	})

	t.Run("returns true if option is enabled with other option", func(test *testing.T) {
		options = []option.Option{"noSuffix", "upcoming"}

		if res := optionIsEnabled("upcoming"); res == false {
			test.Error("Result must be true, but got false instead")
		}

		options = []option.Option{}
	})

	t.Run("returns false if option is disabled", func(test *testing.T) {
		if res := optionIsEnabled("noSuffix"); res == true {
			test.Error("Result must be true, but got false instead")
		}
	})
}

func getTimestampOfPastDate(subDuration time.Duration) int {
	return int(time.Now().Add(-subDuration).UnixNano() / 1000000000)
}
