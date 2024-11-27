package timeago

import (
	"testing"
	"time"
)

func TestParseFunctionCanExceptTimestamp(t *testing.T) {
	cases := []struct {
		date int
		res  string
	}{
		{timestampFromPastDate(time.Minute), "1 minute ago"},
		{timestampFromPastDate(time.Minute * 5), "5 minutes ago"},
		{timestampFromPastDate(time.Hour), "1 hour ago"},
		{timestampFromPastDate(time.Hour * 3), "3 hours ago"},
		{timestampFromPastDate(time.Hour * 5), "5 hours ago"},
		{timestampFromPastDate(time.Hour * 24), "1 day ago"},
		{timestampFromPastDate(time.Hour * 24 * 2), "2 days ago"},
		{timestampFromPastDate(time.Hour * 24 * 3), "3 days ago"},
		{timestampFromPastDate(time.Hour * 24 * 4), "4 days ago"},
		{timestampFromPastDate(time.Hour * 24 * 5), "5 days ago"},
		{timestampFromPastDate(time.Hour * 24 * 6), "6 days ago"},
		{timestampFromPastDate(time.Hour * 24 * 7), "1 week ago"},
	}

	Reconfigure(Config{Language: LangEn})

	for _, tc := range cases {
		t.Run(tc.res, func(t *testing.T) {
			res, err := Parse(tc.date)

			if err != nil {
				t.Errorf("Error must be nil, but got %q instead", err)
			}

			if res != tc.res {
				t.Errorf("Result must be %q, but got %q instead", tc.res, res)
			}
		})
	}
}

func TestParseFunctionCanExceptTimePackage(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
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

	Reconfigure(Config{Language: LangEn})

	for _, tc := range cases {
		t.Run("Test for date "+tc.date.String(), func(t *testing.T) {
			res, err := Parse(tc.date)

			if err != nil {
				t.Errorf("Error must be nil, but got %q instead", err)
			}

			if res != tc.res {
				t.Errorf("Result must be %q, but got %q instead", tc.res, res)
			}
		})
	}
}

func TestParseFuncWillCalculateIntervalToFutureDate(t *testing.T) {
	testCases := []struct {
		date time.Time
		res  string
	}{
		{time.Now().Add(time.Minute * 2), "2 minutes"},
		{time.Now().Add(time.Minute * 5), "5 minutes"},
		{time.Now().Add(time.Minute * 10), "10 minutes"},
		{time.Now().Add(time.Hour), "1 hour"},
		{time.Now().Add(time.Hour * 24), "1 day"},
		{time.Now().Add(time.Hour * 48), "2 days"},
	}

	Reconfigure(Config{Language: "en"})

	for _, tc := range testCases {
		t.Run("Test for date: "+tc.date.String(), func(t *testing.T) {
			res, err := Parse(tc.date)

			if err != nil {
				t.Errorf("Error must be nil, but got %q instead", err)
			}

			if res != tc.res {
				t.Errorf("Result must be %q, but got %q instead", tc.res, res)
			}
		})
	}
}
