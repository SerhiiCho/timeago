package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
)

func TestParseBe(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{subMinutes(1), "1 хвіліна таму"},
		{subMinutes(2), "2 хвіліны таму"},
		{subMinutes(5), "5 хвілін таму"},
		{subMinutes(9), "9 хвілін таму"},
		{subMinutes(10), "10 хвілін таму"},
		{subMinutes(11), "11 хвілін таму"},
		{subMinutes(20), "20 хвілін таму"},
		{subMinutes(21), "21 хвіліна таму"},
		{subMinutes(22), "22 хвіліны таму"},
		{subMinutes(30), "30 хвілін таму"},
		{subMinutes(31), "31 хвіліна таму"},
		{subMinutes(59), "59 хвілін таму"},
		{subHours(1), "1 гадзіна таму"},
		{subHours(2), "2 гадзіны таму"},
		{subHours(9), "9 гадзін таму"},
		{subHours(10), "10 гадзін таму"},
		{subHours(11), "11 гадзін таму"},
		{subHours(20), "20 гадзін таму"},
		{subHours(21), "21 гадзіна таму"},
		{subHours(23), "23 гадзіны таму"},
		{subDays(1), "1 дзень таму"},
		{subDays(2), "2 дні таму"},
		{subDays(6), "6 дзён таму"},
		{subWeeks(1), "1 тыдзень таму"},
		{subWeeks(2), "2 тыдні таму"},
		{subWeeks(3), "3 тыдні таму"},
		{subMonths(1), "1 месяц таму"},
		{subMonths(2), "2 месяцы таму"},
		{subMonths(9), "9 месяцаў таму"},
		{subMonths(11), "11 месяцаў таму"},
		{subYears(1), "1 год таму"},
		{subYears(2), "2 гады таму"},
		{subYears(5), "5 гадоў таму"},
		{subYears(6), "6 гадоў таму"},
		{subYears(7), "7 гадоў таму"},
		{subYears(21), "21 год таму"},
		{subYears(31), "31 год таму"},
		{subYears(100), "100 гадоў таму"},
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
		{subSeconds(0), []string{"0 секунд таму", "1 секунда таму"}},
		{subSeconds(1), []string{"1 секунда таму", "2 секунды таму"}},
		{subSeconds(2), []string{"2 секунды таму", "3 секунды таму"}},
		{subSeconds(3), []string{"3 секунды таму", "4 секунды таму"}},
		{subSeconds(4), []string{"4 секунды таму", "5 секунд таму"}},
		{subSeconds(9), []string{"9 секунд таму", "10 секунд таму"}},
		{subSeconds(10), []string{"10 секунд таму", "11 секунд таму"}},
		{subSeconds(11), []string{"11 секунд таму", "12 секунд таму"}},
		{subSeconds(29), []string{"29 секунд таму", "30 секунд таму"}},
		{subSeconds(30), []string{"30 секунд таму", "31 секунда таму"}},
		{subSeconds(31), []string{"31 секунда таму", "32 секунды таму"}},
		{subSeconds(59), []string{"59 секунд таму", "1 хвіліна таму"}},
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
