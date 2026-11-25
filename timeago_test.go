package timeago

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestParse_UintOverflow(t *testing.T) {
	Reconfigure(Config{Language: LangEn})

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
			res, err := Parse(tc.date)
			if err == nil {
				t.Fatalf("want an error for Parse(%T), got %q instead", tc.date, res)
			}

			if res != "" {
				t.Errorf("want empty result, got %q instead", res)
			}
		})
	}
}

func TestParse_InvalidDateString(t *testing.T) {
	Reconfigure(Config{Language: LangEn})

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
			date, err := Parse(tc.date)
			if err == nil {
				t.Fatalf("want an error for Parse(%q), got %q instead", tc.date, date)
			}

			if date != "" {
				t.Errorf("want empty result, got %q instead", date)
			}
		})
	}
}

func TestParse_InvalidDateType(t *testing.T) {
	Reconfigure(Config{Language: LangEn})

	cases := []any{nil, 1.5, true, []int{1}}

	for _, tc := range cases {
		t.Run(strconv.Quote(fmt.Sprintf("%v", tc)), func(t *testing.T) {
			date, err := Parse(tc)
			if err == nil {
				t.Fatalf("want an error for Parse(%T), got %q instead", tc, date)
			}

			if date != "" {
				t.Errorf("want empty result, got %q instead", date)
			}
		})
	}
}
