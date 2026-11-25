package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v4"
)

func TestParseDateTimeStringWithLocation(t *testing.T) {
	t.Setenv("TZ", "UTC")

	t.Cleanup(func() {
		ago.Reconfigure(ago.Config{Language: ago.LangEn})
	})

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("could not load location: %v", err)
	}

	ago.Reconfigure(ago.Config{Language: ago.LangEn, Location: loc.String()})

	date := time.Now().Truncate(time.Second).Add(-5 * minute).In(loc)
	res, err := ago.Parse(date.Format(time.DateTime))
	assertResult(t, "5 minutes ago", res, err)
}

func TestParseSameInstantInDifferentLocations(t *testing.T) {
	t.Setenv("TZ", "UTC")

	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("could not load location: %v", err)
	}

	instant := time.Now().Truncate(time.Second).Add(-5 * minute)

	cases := []struct {
		name string
		date time.Time
	}{
		{name: "UTC time", date: instant.UTC()},
		{name: "Tokyo time", date: instant.In(loc)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ago.Parse(tc.date)
			assertResult(t, "5 minutes ago", res, err)
		})
	}
}
