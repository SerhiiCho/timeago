package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseRu(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{utils.SubMinutes(1), "1 минута назад"},
		{utils.SubMinutes(2), "2 минуты назад"},
		{utils.SubMinutes(5), "5 минут назад"},
		{utils.SubMinutes(9), "9 минут назад"},
		{utils.SubMinutes(10), "10 минут назад"},
		{utils.SubMinutes(11), "11 минут назад"},
		{utils.SubMinutes(20), "20 минут назад"},
		{utils.SubMinutes(21), "21 минута назад"},
		{utils.SubMinutes(22), "22 минуты назад"},
		{utils.SubMinutes(30), "30 минут назад"},
		{utils.SubMinutes(31), "31 минута назад"},
		{utils.SubMinutes(59), "59 минут назад"},
		{utils.SubHours(1), "1 час назад"},
		{utils.SubHours(2), "2 часа назад"},
		{utils.SubHours(9), "9 часов назад"},
		{utils.SubHours(10), "10 часов назад"},
		{utils.SubHours(11), "11 часов назад"},
		{utils.SubHours(20), "20 часов назад"},
		{utils.SubHours(21), "21 час назад"},
		{utils.SubHours(23), "23 часа назад"},
		{utils.SubDays(1), "1 день назад"},
		{utils.SubDays(2), "2 дня назад"},
		{utils.SubDays(6), "6 дней назад"},
		{utils.SubWeeks(1), "1 неделя назад"},
		{utils.SubWeeks(2), "2 недели назад"},
		{utils.SubWeeks(3), "3 недели назад"},
		{utils.SubMonths(1), "1 месяц назад"},
		{utils.SubMonths(2), "2 месяца назад"},
		{utils.SubMonths(9), "9 месяцев назад"},
		{utils.SubMonths(11), "11 месяцев назад"},
		{utils.SubYears(1), "1 год назад"},
		{utils.SubYears(2), "2 года назад"},
		{utils.SubYears(5), "5 лет назад"},
		{utils.SubYears(6), "6 лет назад"},
		{utils.SubYears(7), "7 лет назад"},
		{utils.SubYears(21), "21 год назад"},
		{utils.SubYears(31), "31 год назад"},
		{utils.SubYears(100), "100 лет назад"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangRu})

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

func TestParseRuWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{utils.SubSeconds(0), []string{"0 секунд назад", "1 секунда назад"}},
		{utils.SubSeconds(1), []string{"1 секунда назад", "2 секунды назад"}},
		{utils.SubSeconds(2), []string{"2 секунды назад", "3 секунды назад"}},
		{utils.SubSeconds(3), []string{"3 секунды назад", "4 секунды назад"}},
		{utils.SubSeconds(4), []string{"4 секунды назад", "5 секунд назад"}},
		{utils.SubSeconds(5), []string{"5 секунд назад", "6 секунд назад"}},
		{utils.SubSeconds(6), []string{"6 секунд назад", "7 секунд назад"}},
		{utils.SubSeconds(7), []string{"7 секунд назад", "8 секунд назад"}},
		{utils.SubSeconds(8), []string{"8 секунд назад", "9 секунд назад"}},
		{utils.SubSeconds(9), []string{"9 секунд назад", "10 секунд назад"}},
		{utils.SubSeconds(10), []string{"10 секунд назад", "11 секунд назад"}},
		{utils.SubSeconds(11), []string{"11 секунд назад", "12 секунд назад"}},
		{utils.SubSeconds(14), []string{"14 секунд назад", "15 секунд назад"}},
		{utils.SubSeconds(16), []string{"16 секунд назад", "17 секунд назад"}},
		{utils.SubSeconds(21), []string{"21 секунда назад", "22 секунды назад"}},
		{utils.SubSeconds(23), []string{"23 секунды назад", "24 секунды назад"}},
		{utils.SubSeconds(25), []string{"25 секунд назад", "26 секунд назад"}},
		{utils.SubSeconds(29), []string{"29 секунд назад", "30 секунд назад"}},
		{utils.SubSeconds(30), []string{"30 секунд назад", "31 секунда назад"}},
		{utils.SubSeconds(31), []string{"31 секунда назад", "32 секунды назад"}},
		{utils.SubSeconds(32), []string{"32 секунды назад", "33 секунды назад"}},
		{utils.SubSeconds(33), []string{"33 секунды назад", "34 секунды назад"}},
		{utils.SubSeconds(34), []string{"34 секунды назад", "35 секунд назад"}},
		{utils.SubSeconds(35), []string{"35 секунд назад", "36 секунд назад"}},
		{utils.SubSeconds(36), []string{"36 секунд назад", "37 секунд назад"}},
		{utils.SubSeconds(37), []string{"37 секунд назад", "38 секунд назад"}},
		{utils.SubSeconds(39), []string{"39 секунд назад", "40 секунд назад"}},
		{utils.SubSeconds(40), []string{"40 секунд назад", "41 секунда назад"}},
		{utils.SubSeconds(41), []string{"41 секунда назад", "42 секунды назад"}},
		{utils.SubSeconds(59), []string{"59 секунд назад", "1 минута назад"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangRu})

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
