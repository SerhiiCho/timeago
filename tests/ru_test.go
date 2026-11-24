package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
)

func TestParseRu(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{subMins(1), "1 минута назад"},
		{subMins(2), "2 минуты назад"},
		{subMins(5), "5 минут назад"},
		{subMins(9), "9 минут назад"},
		{subMins(10), "10 минут назад"},
		{subMins(11), "11 минут назад"},
		{subMins(20), "20 минут назад"},
		{subMins(21), "21 минута назад"},
		{subMins(22), "22 минуты назад"},
		{subMins(30), "30 минут назад"},
		{subMins(31), "31 минута назад"},
		{subMins(59), "59 минут назад"},
		{subHours(1), "1 час назад"},
		{subHours(2), "2 часа назад"},
		{subHours(9), "9 часов назад"},
		{subHours(10), "10 часов назад"},
		{subHours(11), "11 часов назад"},
		{subHours(20), "20 часов назад"},
		{subHours(21), "21 час назад"},
		{subHours(23), "23 часа назад"},
		{subDays(1), "1 день назад"},
		{subDays(2), "2 дня назад"},
		{subDays(6), "6 дней назад"},
		{subWeeks(1), "1 неделя назад"},
		{subWeeks(2), "2 недели назад"},
		{subWeeks(3), "3 недели назад"},
		{subMonths(1), "1 месяц назад"},
		{subMonths(2), "2 месяца назад"},
		{subMonths(9), "9 месяцев назад"},
		{subMonths(11), "11 месяцев назад"},
		{subYears(1), "1 год назад"},
		{subYears(2), "2 года назад"},
		{subYears(5), "5 лет назад"},
		{subYears(6), "6 лет назад"},
		{subYears(7), "7 лет назад"},
		{subYears(21), "21 год назад"},
		{subYears(31), "31 год назад"},
		{subYears(100), "100 лет назад"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangRu})

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

func TestParseRuWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{subSecs(0), []string{"0 секунд назад", "1 секунда назад"}},
		{subSecs(1), []string{"1 секунда назад", "2 секунды назад"}},
		{subSecs(2), []string{"2 секунды назад", "3 секунды назад"}},
		{subSecs(3), []string{"3 секунды назад", "4 секунды назад"}},
		{subSecs(4), []string{"4 секунды назад", "5 секунд назад"}},
		{subSecs(5), []string{"5 секунд назад", "6 секунд назад"}},
		{subSecs(6), []string{"6 секунд назад", "7 секунд назад"}},
		{subSecs(7), []string{"7 секунд назад", "8 секунд назад"}},
		{subSecs(8), []string{"8 секунд назад", "9 секунд назад"}},
		{subSecs(9), []string{"9 секунд назад", "10 секунд назад"}},
		{subSecs(10), []string{"10 секунд назад", "11 секунд назад"}},
		{subSecs(11), []string{"11 секунд назад", "12 секунд назад"}},
		{subSecs(14), []string{"14 секунд назад", "15 секунд назад"}},
		{subSecs(16), []string{"16 секунд назад", "17 секунд назад"}},
		{subSecs(21), []string{"21 секунда назад", "22 секунды назад"}},
		{subSecs(23), []string{"23 секунды назад", "24 секунды назад"}},
		{subSecs(25), []string{"25 секунд назад", "26 секунд назад"}},
		{subSecs(29), []string{"29 секунд назад", "30 секунд назад"}},
		{subSecs(30), []string{"30 секунд назад", "31 секунда назад"}},
		{subSecs(31), []string{"31 секунда назад", "32 секунды назад"}},
		{subSecs(32), []string{"32 секунды назад", "33 секунды назад"}},
		{subSecs(33), []string{"33 секунды назад", "34 секунды назад"}},
		{subSecs(34), []string{"34 секунды назад", "35 секунд назад"}},
		{subSecs(35), []string{"35 секунд назад", "36 секунд назад"}},
		{subSecs(36), []string{"36 секунд назад", "37 секунд назад"}},
		{subSecs(37), []string{"37 секунд назад", "38 секунд назад"}},
		{subSecs(39), []string{"39 секунд назад", "40 секунд назад"}},
		{subSecs(40), []string{"40 секунд назад", "41 секунда назад"}},
		{subSecs(41), []string{"41 секунда назад", "42 секунды назад"}},
		{subSecs(59), []string{"59 секунд назад", "1 минута назад"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangRu})

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
