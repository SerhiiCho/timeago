package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v4"
)

func TestParseUnitBoundaries(t *testing.T) {
	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	cases := []struct {
		name     string
		duration time.Duration
		want     string
	}{
		{name: "59 seconds stays in seconds", duration: 59 * second, want: "59 seconds ago"},
		{name: "60 seconds becomes a minute", duration: minute, want: "1 minute ago"},
		{name: "59 minutes stays in minutes", duration: 59 * minute, want: "59 minutes ago"},
		{name: "60 minutes becomes an hour", duration: hour, want: "1 hour ago"},
		{name: "23 hours stays in hours", duration: 23 * hour, want: "23 hours ago"},
		{name: "24 hours becomes a day", duration: day, want: "1 day ago"},
		{name: "6 days stays in days", duration: 6 * day, want: "6 days ago"},
		{name: "7 days becomes a week", duration: week, want: "1 week ago"},
		{name: "27 days stays in weeks", duration: 27 * day, want: "3 weeks ago"},
		{name: "28 days becomes a month", duration: 4 * week, want: "1 month ago"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ago.Parse(subTime(tc.duration))
			assertResult(t, tc.want, res, err)
		})
	}
}

func TestParseFutureMonthAndYear(t *testing.T) {
	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	cases := []struct {
		name   string
		months int
		want   string
	}{
		{name: "1 month ahead", months: 1, want: "1 month"},
		{name: "11 months ahead", months: 11, want: "11 months"},
		{name: "1 year ahead", months: 12, want: "1 year"},
		{name: "13 months ahead is one year", months: 13, want: "1 year"},
		{name: "2 years ahead", months: 24, want: "2 years"},
		{name: "10 years ahead", months: 12 * 10, want: "10 years"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			date := wantAddMonths(time.Now().Truncate(time.Second), tc.months)
			res, err := ago.Parse(date)
			assertResult(t, tc.want, res, err)
		})
	}
}

func TestParseDateTimeString(t *testing.T) {
	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	now := time.Now().Truncate(time.Second)

	cases := []struct {
		name string
		date time.Time
		want string
	}{
		{
			name: "datetime string in the past",
			date: now.UTC().Add(-5 * minute),
			want: "5 minutes ago",
		},
		{
			name: "datetime string in the future",
			date: now.UTC().Add(5*minute + second),
			want: "5 minutes",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ago.Parse(tc.date.Format(time.DateTime))
			assertResult(t, tc.want, res, err)
		})
	}
}
