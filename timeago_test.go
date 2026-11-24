package timeago

import (
	"fmt"
	"strconv"
	"testing"
)

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
