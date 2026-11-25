package tests

import (
	"fmt"
	"math"
	"strconv"
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v4"
)

func TestParseUintOverflow(t *testing.T) {
	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	cases := []struct {
		name string
		date any
	}{
		{
			name: "max int64 + 1",
			date: uint64(math.MaxInt64) + 1,
		},
		{
			name: "max uint + 1",
			date: uint(math.MaxInt) + 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ago.Parse(tc.date)
			if err == nil {
				t.Fatalf("want an error for Parse(%T), got %q instead", tc.date, res)
			}

			if res != "" {
				t.Errorf("want empty result, got %q instead", res)
			}
		})
	}
}

func TestParseInvalidDateString(t *testing.T) {
	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	cases := []struct {
		name string
		date string
	}{
		{
			name: "empty string",
			date: "",
		},
		{
			name: "random text",
			date: "not a date",
		},
		{
			name: "month out of range",
			date: "2026-13-01 00:00:00",
		},
		{
			name: "day out of range",
			date: "2026-02-30 00:00:00",
		},
		{
			name: "wrong format",
			date: "25.11.2026",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ago.Parse(tc.date)
			if err == nil {
				t.Fatalf("want an error for Parse(%q), got %q instead", tc.date, res)
			}

			if res != "" {
				t.Errorf("want empty result, got %q instead", res)
			}
		})
	}
}

func TestParseInvalidDateType(t *testing.T) {
	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	cases := []any{nil, 1.5, true, []int{1}}

	for _, date := range cases {
		t.Run(strconv.Quote(fmt.Sprintf("%v", date)), func(t *testing.T) {
			res, err := ago.Parse(date)
			if err == nil {
				t.Fatalf("want an error for Parse(%T), got %q instead", date, res)
			}

			if res != "" {
				t.Errorf("want empty result, got %q instead", res)
			}
		})
	}
}

func TestParseInvalidLocation(t *testing.T) {
	t.Cleanup(func() {
		ago.Reconfigure(ago.Config{Language: ago.LangEn})
	})

	ago.Reconfigure(ago.Config{Language: ago.LangEn, Location: "Mars/Olympus"})

	cases := []struct {
		name string
		date any
	}{
		{name: "datetime string", date: "2026-11-25 10:00:00"},
		{name: "time.Time value", date: time.Now().Add(-time.Minute)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ago.Parse(tc.date)
			if err == nil {
				t.Fatalf("want an error, got %q instead", res)
			}

			if res != "" {
				t.Errorf("want empty result, got %q instead", res)
			}
		})
	}
}

func TestParseInvalidLanguage(t *testing.T) {
	t.Cleanup(func() {
		ago.Reconfigure(ago.Config{Language: ago.LangEn})
	})

	ago.Reconfigure(ago.Config{Language: "xx"})

	res, err := ago.Parse(time.Now().Add(-time.Minute))
	if err == nil {
		t.Fatalf("want an error, got %q instead", res)
	}

	if res != "" {
		t.Errorf("want empty result, got %q instead", res)
	}
}

func TestParseTimestampStringOverflow(t *testing.T) {
	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	res, err := ago.Parse("9223372036854775808")
	if err == nil {
		t.Fatalf("want an error, got %q instead", res)
	}

	if res != "" {
		t.Errorf("want empty result, got %q instead", res)
	}
}
