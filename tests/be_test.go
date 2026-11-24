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
		{subMins(1), "1 хвіліна таму"},
		{subMins(2), "2 хвіліны таму"},
		{subMins(5), "5 хвілін таму"},
		{subMins(9), "9 хвілін таму"},
		{subMins(10), "10 хвілін таму"},
		{subMins(11), "11 хвілін таму"},
		{subMins(20), "20 хвілін таму"},
		{subMins(21), "21 хвіліна таму"},
		{subMins(22), "22 хвіліны таму"},
		{subMins(30), "30 хвілін таму"},
		{subMins(31), "31 хвіліна таму"},
		{subMins(59), "59 хвілін таму"},
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
		{subYears(111), "111 гадоў таму"},
		{subYears(112), "112 гадоў таму"},
		{subYears(113), "113 гадоў таму"},
		{subYears(114), "114 гадоў таму"},
		{subYears(115), "115 гадоў таму"},
		{subYears(117), "117 гадоў таму"},
		{subYears(118), "118 гадоў таму"},
		{subYears(119), "119 гадоў таму"},
		{subYears(120), "120 гадоў таму"},
		{subYears(121), "121 год таму"},
		{subYears(122), "122 гады таму"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangBe})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("want nil error, got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("want %q, got %q instead", tc.res, res)
			}
		})
	}
}

func TestParseBeWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{subSecs(0), []string{"0 секунд таму", "1 секунда таму"}},
		{subSecs(1), []string{"1 секунда таму", "2 секунды таму"}},
		{subSecs(2), []string{"2 секунды таму", "3 секунды таму"}},
		{subSecs(3), []string{"3 секунды таму", "4 секунды таму"}},
		{subSecs(4), []string{"4 секунды таму", "5 секунд таму"}},
		{subSecs(5), []string{"5 секунд таму", "6 секунд таму"}},
		{subSecs(6), []string{"6 секунд таму", "7 секунд таму"}},
		{subSecs(7), []string{"7 секунд таму", "8 секунд таму"}},
		{subSecs(8), []string{"8 секунд таму", "9 секунд таму"}},
		{subSecs(9), []string{"9 секунд таму", "10 секунд таму"}},
		{subSecs(10), []string{"10 секунд таму", "11 секунд таму"}},
		{subSecs(11), []string{"11 секунд таму", "12 секунд таму"}},
		{subSecs(14), []string{"14 секунд таму", "15 секунд таму"}},
		{subSecs(16), []string{"16 секунд таму", "17 секунд таму"}},
		{subSecs(21), []string{"21 секунда таму", "22 секунды таму"}},
		{subSecs(23), []string{"23 секунды таму", "24 секунды таму"}},
		{subSecs(25), []string{"25 секунд таму", "26 секунд таму"}},
		{subSecs(29), []string{"29 секунд таму", "30 секунд таму"}},
		{subSecs(30), []string{"30 секунд таму", "31 секунда таму"}},
		{subSecs(31), []string{"31 секунда таму", "32 секунды таму"}},
		{subSecs(32), []string{"32 секунды таму", "33 секунды таму"}},
		{subSecs(33), []string{"33 секунды таму", "34 секунды таму"}},
		{subSecs(34), []string{"34 секунды таму", "35 секунд таму"}},
		{subSecs(35), []string{"35 секунд таму", "36 секунд таму"}},
		{subSecs(36), []string{"36 секунд таму", "37 секунд таму"}},
		{subSecs(37), []string{"37 секунд таму", "38 секунд таму"}},
		{subSecs(39), []string{"39 секунд таму", "40 секунд таму"}},
		{subSecs(40), []string{"40 секунд таму", "41 секунда таму"}},
		{subSecs(41), []string{"41 секунда таму", "42 секунды таму"}},
		{subSecs(59), []string{"59 секунд таму", "1 хвіліна таму"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangBe})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("want nil error, got %v instead", err)
			}

			if res != tc.res[0] && res != tc.res[1] {
				test.Errorf(
					"want %q or %q, got %q instead",
					tc.res[0],
					tc.res[1],
					res,
				)
			}
		})
	}
}
