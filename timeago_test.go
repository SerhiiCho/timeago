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

func TestParse_InvalidDateType(t *testing.T) {
	Reconfigure(Config{Language: LangEn})

	tc := []any{nil, 1.5, true, []int{1}}

	for _, date := range tc {
		t.Run(strconv.Quote(fmt.Sprintf("%v", date)), func(t *testing.T) {
			res, err := Parse(date)
			if err == nil {
				t.Fatalf("want an error for Parse(%T), got %q instead", date, res)
			}

			if res != "" {
				t.Errorf("want empty result, got %q instead", res)
			}
		})
	}
}
