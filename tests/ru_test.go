package tests

import (
	"testing"
	"time"

	. "github.com/SerhiiCho/timeago/v2"
)

const langRu = "ru"

func TestParseRu(t *testing.T) {
	cases := []struct {
		date   time.Time
		result string
	}{
		{subMinutes(1), "1 минута назад"},
		{subMinutes(2), "2 минуты назад"},
		{subMinutes(5), "5 минут назад"},
		{subMinutes(9), "9 минут назад"},
		{subMinutes(10), "10 минут назад"},
		{subMinutes(11), "11 минут назад"},
		{subMinutes(20), "20 минут назад"},
		{subMinutes(21), "21 минута назад"},
		{subMinutes(22), "22 минуты назад"},
		{subMinutes(30), "30 минут назад"},
		{subMinutes(31), "31 минута назад"},
		{subMinutes(59), "59 минут назад"},
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
			SetConfig(Config{
				Language: langRu,
			})

			if res := Parse(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseRuWithSeconds(t *testing.T) {
	cases := []struct {
		date   time.Time
		result []string
	}{
		{subSeconds(0), []string{"0 секунд назад", "1 секунда назад"}},
		{subSeconds(1), []string{"1 секунда назад", "2 секунды назад"}},
		{subSeconds(2), []string{"2 секунды назад", "3 секунды назад"}},
		{subSeconds(3), []string{"3 секунды назад", "4 секунды назад"}},
		{subSeconds(4), []string{"4 секунды назад", "5 секунд назад"}},
		{subSeconds(9), []string{"9 секунд назад", "10 секунд назад"}},
		{subSeconds(10), []string{"10 секунд назад", "11 секунд назад"}},
		{subSeconds(11), []string{"11 секунд назад", "12 секунд назад"}},
		{subSeconds(29), []string{"29 секунд назад", "30 секунд назад"}},
		{subSeconds(30), []string{"30 секунд назад", "31 секунда назад"}},
		{subSeconds(31), []string{"31 секунда назад", "32 секунды назад"}},
		{subSeconds(59), []string{"59 секунд назад", "1 минута назад"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			SetConfig(Config{
				Language: langRu,
			})

			if res := Parse(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
