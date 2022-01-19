package tests

import (
	"testing"
	"time"

	. "github.com/SerhiiCho/timeago"
	. "github.com/SerhiiCho/timeago/utils"
)

func TestParseRu(t *testing.T) {
	cases := []struct {
		date   string
		result string
	}{
		{SmallSubTime(-60 * time.Second), "1 минута назад"},
		{SmallSubTime(-1 * time.Minute), "1 минута назад"},
		{SmallSubTime(-2 * time.Minute), "2 минуты назад"},
		{SmallSubTime(-5 * time.Minute), "5 минут назад"},
		{SmallSubTime(-9 * time.Minute), "9 минут назад"},
		{SmallSubTime(-10 * time.Minute), "10 минут назад"},
		{SmallSubTime(-11 * time.Minute), "11 минут назад"},
		{SmallSubTime(-20 * time.Minute), "20 минут назад"},
		{SmallSubTime(-21 * time.Minute), "21 минута назад"},
		{SmallSubTime(-22 * time.Minute), "22 минуты назад"},
		{SmallSubTime(-30 * time.Minute), "30 минут назад"},
		{SmallSubTime(-31 * time.Minute), "31 минута назад"},
		{SmallSubTime(-59 * time.Minute), "59 минут назад"},
		{SmallSubTime(-60 * time.Minute), "1 час назад"},
		{SmallSubTime(-1 * time.Hour), "1 час назад"},
		{SmallSubTime(-2 * time.Hour), "2 часа назад"},
		{SmallSubTime(-9 * time.Hour), "9 часов назад"},
		{SmallSubTime(-10 * time.Hour), "10 часов назад"},
		{SmallSubTime(-11 * time.Hour), "11 часов назад"},
		{SmallSubTime(-20 * time.Hour), "20 часов назад"},
		{SmallSubTime(-21 * time.Hour), "21 час назад"},
		{SmallSubTime(-23 * time.Hour), "23 часа назад"},
		{SmallSubTime(-24 * time.Hour), "1 день назад"},
		{SmallSubTime(-30 * time.Hour), "1 день назад"},
		{SmallSubTime((-24 * 2) * time.Hour), "2 дня назад"},
		{SmallSubTime((-24 * 6) * time.Hour), "6 дней назад"},
		{SmallSubTime((-24 * 7) * time.Hour), "1 неделя назад"},
		{SmallSubTime((-24 * 14) * time.Hour), "2 недели назад"},
		{SmallSubTime((-24 * 21) * time.Hour), "3 недели назад"},
		{BigSubTime(0, 1, 1), "1 месяц назад"},
		{BigSubTime(0, 2, 1), "2 месяца назад"},
		{BigSubTime(0, 9, 1), "9 месяцев назад"},
		{BigSubTime(0, 11, 1), "11 месяцев назад"},
		{BigSubTime(0, 12, 1), "1 год назад"},
		{BigSubTime(1, 0, 1), "1 год назад"},
		{BigSubTime(2, 0, 1), "2 года назад"},
		{BigSubTime(5, 0, 1), "5 лет назад"},
		{BigSubTime(6, 0, 1), "6 лет назад"},
		{BigSubTime(7, 0, 1), "7 лет назад"},
		{BigSubTime(21, 0, 1), "21 год назад"},
		{BigSubTime(31, 0, 1), "31 год назад"},
		{BigSubTime(100, 0, 1), "100 лет назад"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			SetConfig(Config{
				Language: "ru",
				Location: "Europe/Kiev",
			})

			if res := Parse(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseRuWithOnlineFlag(t *testing.T) {
	cases := []struct {
		date   string
		result string
	}{
		{SmallSubTime(time.Second * 2), "В сети"},
		{SmallSubTime(time.Second), "В сети"},
		{SmallSubTime(-1 * time.Second), "В сети"},
		{SmallSubTime(-2 * time.Second), "В сети"},
		{SmallSubTime(-9 * time.Second), "В сети"},
		{SmallSubTime(-10 * time.Second), "В сети"},
		{SmallSubTime(-11 * time.Second), "В сети"},
		{SmallSubTime(-20 * time.Second), "В сети"},
		{SmallSubTime(-21 * time.Second), "В сети"},
		{SmallSubTime(-22 * time.Second), "В сети"},
		{SmallSubTime(-30 * time.Second), "В сети"},
		{SmallSubTime(-31 * time.Second), "В сети"},
		{SmallSubTime(-60 * time.Second), "1 минута назад"},
		{SmallSubTime(-1 * time.Minute), "1 минута назад"},
		{SmallSubTime(-2 * time.Minute), "2 минуты назад"},
		{SmallSubTime(-9 * time.Minute), "9 минут назад"},
		{SmallSubTime(-10 * time.Minute), "10 минут назад"},
		{SmallSubTime(-11 * time.Minute), "11 минут назад"},
		{SmallSubTime(-20 * time.Minute), "20 минут назад"},
		{SmallSubTime(-21 * time.Minute), "21 минута назад"},
		{SmallSubTime(-22 * time.Minute), "22 минуты назад"},
		{SmallSubTime(-30 * time.Minute), "30 минут назад"},
		{SmallSubTime(-31 * time.Minute), "31 минута назад"},
		{SmallSubTime(-60 * time.Minute), "1 час назад"},
		{SmallSubTime(-1 * time.Hour), "1 час назад"},
		{SmallSubTime(-2 * time.Hour), "2 часа назад"},
		{SmallSubTime(-9 * time.Hour), "9 часов назад"},
		{SmallSubTime(-10 * time.Hour), "10 часов назад"},
		{SmallSubTime(-11 * time.Hour), "11 часов назад"},
		{SmallSubTime(-20 * time.Hour), "20 часов назад"},
		{SmallSubTime(-21 * time.Hour), "21 час назад"},
		{SmallSubTime(-23 * time.Hour), "23 часа назад"},
		{SmallSubTime(-24 * time.Hour), "1 день назад"},
		{SmallSubTime(-30 * time.Hour), "1 день назад"},
		{SmallSubTime((-24 * 2) * time.Hour), "2 дня назад"},
		{SmallSubTime((-24 * 6) * time.Hour), "6 дней назад"},
		{SmallSubTime((-24 * 7) * time.Hour), "1 неделя назад"},
		{SmallSubTime((-24 * 14) * time.Hour), "2 недели назад"},
		{SmallSubTime((-24 * 21) * time.Hour), "3 недели назад"},
		{BigSubTime(0, 1, 1), "1 месяц назад"},
		{BigSubTime(0, 2, 1), "2 месяца назад"},
		{BigSubTime(0, 9, 1), "9 месяцев назад"},
		{BigSubTime(0, 11, 1), "11 месяцев назад"},
		{BigSubTime(0, 12, 1), "1 год назад"},
		{BigSubTime(1, 0, 1), "1 год назад"},
		{BigSubTime(2, 0, 1), "2 года назад"},
		{BigSubTime(5, 0, 1), "5 лет назад"},
		{BigSubTime(6, 0, 1), "6 лет назад"},
		{BigSubTime(7, 0, 1), "7 лет назад"},
		{BigSubTime(21, 0, 1), "21 год назад"},
		{BigSubTime(31, 0, 1), "31 год назад"},
		{BigSubTime(100, 0, 1), "100 лет назад"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {

			SetConfig(Config{
				Language: "ru",
			})

			if res := Parse(tc.date, "online"); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseRuWithSeconds(t *testing.T) {
	cases := []struct {
		date   string
		result []string
	}{
		{SmallSubTime(time.Second * 2), []string{"0 секунд назад", "1 секунда назад"}},
		{SmallSubTime(time.Second), []string{"0 секунд назад", "1 секунда назад"}},
		{SmallSubTime(-1 * time.Second), []string{"1 секунда назад", "2 секунды назад"}},
		{SmallSubTime(-2 * time.Second), []string{"2 секунды назад", "3 секунды назад"}},
		{SmallSubTime(-3 * time.Second), []string{"3 секунды назад", "4 секунды назад"}},
		{SmallSubTime(-4 * time.Second), []string{"4 секунды назад", "5 секунд назад"}},
		{SmallSubTime(-9 * time.Second), []string{"9 секунд назад", "10 секунд назад"}},
		{SmallSubTime(-10 * time.Second), []string{"10 секунд назад", "11 секунд назад"}},
		{SmallSubTime(-11 * time.Second), []string{"11 секунд назад", "12 секунд назад"}},
		{SmallSubTime(-29 * time.Second), []string{"29 секунд назад", "30 секунд назад"}},
		{SmallSubTime(-30 * time.Second), []string{"30 секунд назад", "31 секунда назад"}},
		{SmallSubTime(-31 * time.Second), []string{"31 секунда назад", "32 секунды назад"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			SetConfig(Config{
				Language: "ru",
				Location: "Europe/Kiev",
			})

			if res := Parse(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
