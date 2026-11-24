package timeago

import (
	"testing"
	"time"
)

func TestCalendarMonthsBetween(t *testing.T) {
	getDate := func(year int, month time.Month, date int) time.Time {
		return time.Date(year, month, date, 12, 0, 0, 0, time.UTC)
	}

	cases := []struct {
		name string
		from time.Time
		to   time.Time
		want int
	}{
		{
			name: "leap day to clamped anniversary",
			from: getDate(2024, time.February, 29),
			to:   getDate(2025, time.February, 28),
			want: 12,
		},
		{
			name: "before leap day anniversary",
			from: getDate(2024, time.February, 29),
			to:   getDate(2025, time.February, 27),
			want: 11,
		},
		{
			name: "month end clamps",
			from: getDate(2025, time.January, 31),
			to:   getDate(2025, time.February, 28),
			want: 1,
		},
		{
			name: "before clamped month end",
			from: getDate(2025, time.January, 31),
			to:   getDate(2025, time.February, 27),
			want: 0,
		},
		{
			name: "3 full months on leap year",
			from: getDate(2025, time.January, 1),
			to:   getDate(2025, time.April, 1),
			want: 3,
		},
		{
			name: "3 full months on non-leap year",
			from: getDate(2023, time.January, 1),
			to:   getDate(2023, time.April, 1),
			want: 3,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := calendarMonthsBetween(tc.from, tc.to); got != tc.want {
				t.Errorf("want %d months, got %d", tc.want, got)
			}
		})
	}
}

func TestParserOptionIsEnabled(t *testing.T) {
	cases := []struct {
		name    string
		options []opt
		want    map[opt]bool
	}{
		{
			name:    "returns true if option is enabled",
			options: []opt{OptNoSuffix},
			want:    map[opt]bool{OptNoSuffix: true},
		},
		{
			name:    "returns true if option is enabled with other option",
			options: []opt{OptNoSuffix, OptUpcoming},
			want:    map[opt]bool{OptUpcoming: true},
		},
		{
			name:    "returns true for the first and the last option",
			options: []opt{OptUpcoming, OptOnline, OptNoSuffix},
			want:    map[opt]bool{OptUpcoming: true, OptNoSuffix: true},
		},
		{
			name:    "returns true for duplicated option",
			options: []opt{OptNoSuffix, OptNoSuffix},
			want:    map[opt]bool{OptNoSuffix: true},
		},
		{
			name:    "returns true for every option when all are enabled",
			options: []opt{OptUpcoming, OptOnline, OptJustNow, OptNoSuffix},
			want: map[opt]bool{
				OptUpcoming: true,
				OptOnline:   true,
				OptJustNow:  true,
				OptNoSuffix: true,
			},
		},
		{
			name:    "returns false if option is disabled",
			options: nil,
			want:    map[opt]bool{OptNoSuffix: false},
		},
		{
			name:    "returns false for empty options slice",
			options: []opt{},
			want:    map[opt]bool{OptNoSuffix: false},
		},
		{
			name:    "returns false if option is not in the list",
			options: []opt{OptNoSuffix, OptUpcoming},
			want:    map[opt]bool{OptOnline: false},
		},
		{
			name:    "returns false for unknown option",
			options: []opt{OptNoSuffix},
			want:    map[opt]bool{opt("unknown"): false},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := &parser{options: tc.options}

			for option, expect := range tc.want {
				if res := p.optionIsEnabled(option); res != expect {
					t.Errorf("want optionIsEnabled(%q) = %v, got %v instead", option, expect, res)
				}
			}
		})
	}
}
