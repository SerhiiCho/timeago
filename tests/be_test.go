package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseBe(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{utils.SubMinutes(1), "1 хвіліна таму"},
		{utils.SubMinutes(2), "2 хвіліны таму"},
		{utils.SubMinutes(5), "5 хвілін таму"},
		{utils.SubMinutes(9), "9 хвілін таму"},
		{utils.SubMinutes(10), "10 хвілін таму"},
		{utils.SubMinutes(11), "11 хвілін таму"},
		{utils.SubMinutes(20), "20 хвілін таму"},
		{utils.SubMinutes(21), "21 хвіліна таму"},
		{utils.SubMinutes(22), "22 хвіліны таму"},
		{utils.SubMinutes(30), "30 хвілін таму"},
		{utils.SubMinutes(31), "31 хвіліна таму"},
		{utils.SubMinutes(59), "59 хвілін таму"},
		{utils.SubHours(1), "1 гадзіна таму"},
		{utils.SubHours(2), "2 гадзіны таму"},
		{utils.SubHours(9), "9 гадзін таму"},
		{utils.SubHours(10), "10 гадзін таму"},
		{utils.SubHours(11), "11 гадзін таму"},
		{utils.SubHours(20), "20 гадзін таму"},
		{utils.SubHours(21), "21 гадзіна таму"},
		{utils.SubHours(23), "23 гадзіны таму"},
		{utils.SubDays(1), "1 дзень таму"},
		{utils.SubDays(2), "2 дні таму"},
		{utils.SubDays(6), "6 дзён таму"},
		{utils.SubWeeks(1), "1 тыдзень таму"},
		{utils.SubWeeks(2), "2 тыдні таму"},
		{utils.SubWeeks(3), "3 тыдні таму"},
		{utils.SubMonths(1), "1 месяц таму"},
		{utils.SubMonths(2), "2 месяцы таму"},
		{utils.SubMonths(9), "9 месяцаў таму"},
		{utils.SubMonths(11), "11 месяцаў таму"},
		{utils.SubYears(1), "1 год таму"},
		{utils.SubYears(2), "2 гады таму"},
		{utils.SubYears(5), "5 гадоў таму"},
		{utils.SubYears(6), "6 гадоў таму"},
		{utils.SubYears(7), "7 гадоў таму"},
		{utils.SubYears(21), "21 год таму"},
		{utils.SubYears(31), "31 год таму"},
		{utils.SubYears(100), "100 гадоў таму"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangBe})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("Result must be %s, but got %s instead", tc.res, res)
			}
		})
	}
}

func TestParseBeWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{utils.SubSeconds(0), []string{"0 секунд таму", "1 секунда таму"}},
		{utils.SubSeconds(1), []string{"1 секунда таму", "2 секунды таму"}},
		{utils.SubSeconds(2), []string{"2 секунды таму", "3 секунды таму"}},
		{utils.SubSeconds(3), []string{"3 секунды таму", "4 секунды таму"}},
		{utils.SubSeconds(4), []string{"4 секунды таму", "5 секунд таму"}},
		{utils.SubSeconds(9), []string{"9 секунд таму", "10 секунд таму"}},
		{utils.SubSeconds(10), []string{"10 секунд таму", "11 секунд таму"}},
		{utils.SubSeconds(11), []string{"11 секунд таму", "12 секунд таму"}},
		{utils.SubSeconds(29), []string{"29 секунд таму", "30 секунд таму"}},
		{utils.SubSeconds(30), []string{"30 секунд таму", "31 секунда таму"}},
		{utils.SubSeconds(31), []string{"31 секунда таму", "32 секунды таму"}},
		{utils.SubSeconds(59), []string{"59 секунд таму", "1 хвіліна таму"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangBe})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res[0] && res != tc.res[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.res[0], tc.res[1], res)
			}
		})
	}
}
