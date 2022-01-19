package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago"
	. "github.com/SerhiiCho/timeago/utils"
)

func TestParseRu(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{SmallSubTime(-60 * time.Second), "1 минута назад", "ru"},
		{SmallSubTime(-1 * time.Minute), "1 минута назад", "ru"},
		{SmallSubTime(-2 * time.Minute), "2 минуты назад", "ru"},
		{SmallSubTime(-5 * time.Minute), "5 минут назад", "ru"},
		{SmallSubTime(-9 * time.Minute), "9 минут назад", "ru"},
		{SmallSubTime(-10 * time.Minute), "10 минут назад", "ru"},
		{SmallSubTime(-11 * time.Minute), "11 минут назад", "ru"},
		{SmallSubTime(-20 * time.Minute), "20 минут назад", "ru"},
		{SmallSubTime(-21 * time.Minute), "21 минута назад", "ru"},
		{SmallSubTime(-22 * time.Minute), "22 минуты назад", "ru"},
		{SmallSubTime(-30 * time.Minute), "30 минут назад", "ru"},
		{SmallSubTime(-31 * time.Minute), "31 минута назад", "ru"},
		{SmallSubTime(-59 * time.Minute), "59 минут назад", "ru"},
		{SmallSubTime(-60 * time.Minute), "1 час назад", "ru"},
		{SmallSubTime(-1 * time.Hour), "1 час назад", "ru"},
		{SmallSubTime(-2 * time.Hour), "2 часа назад", "ru"},
		{SmallSubTime(-9 * time.Hour), "9 часов назад", "ru"},
		{SmallSubTime(-10 * time.Hour), "10 часов назад", "ru"},
		{SmallSubTime(-11 * time.Hour), "11 часов назад", "ru"},
		{SmallSubTime(-20 * time.Hour), "20 часов назад", "ru"},
		{SmallSubTime(-21 * time.Hour), "21 час назад", "ru"},
		{SmallSubTime(-23 * time.Hour), "23 часа назад", "ru"},
		{SmallSubTime(-24 * time.Hour), "1 день назад", "ru"},
		{SmallSubTime(-30 * time.Hour), "1 день назад", "ru"},
		{SmallSubTime((-24 * 2) * time.Hour), "2 дня назад", "ru"},
		{SmallSubTime((-24 * 6) * time.Hour), "6 дней назад", "ru"},
		{SmallSubTime((-24 * 7) * time.Hour), "1 неделя назад", "ru"},
		{SmallSubTime((-24 * 14) * time.Hour), "2 недели назад", "ru"},
		{SmallSubTime((-24 * 21) * time.Hour), "3 недели назад", "ru"},
		{BigSubTime(0, 1, 1), "1 месяц назад", "ru"},
		{BigSubTime(0, 2, 1), "2 месяца назад", "ru"},
		{BigSubTime(0, 9, 1), "9 месяцев назад", "ru"},
		{BigSubTime(0, 11, 1), "11 месяцев назад", "ru"},
		{BigSubTime(0, 12, 1), "1 год назад", "ru"},
		{BigSubTime(1, 0, 1), "1 год назад", "ru"},
		{BigSubTime(2, 0, 1), "2 года назад", "ru"},
		{BigSubTime(5, 0, 1), "5 лет назад", "ru"},
		{BigSubTime(6, 0, 1), "6 лет назад", "ru"},
		{BigSubTime(7, 0, 1), "7 лет назад", "ru"},
		{BigSubTime(21, 0, 1), "21 год назад", "ru"},
		{BigSubTime(31, 0, 1), "31 год назад", "ru"},
		{BigSubTime(100, 0, 1), "100 лет назад", "ru"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)
			timeago.Set("location", "Europe/Kiev")

			if res := timeago.Parse(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseRuWithOnlineFlag(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{SmallSubTime(time.Second * 2), "В сети", "ru"},
		{SmallSubTime(time.Second), "В сети", "ru"},
		{SmallSubTime(-1 * time.Second), "В сети", "ru"},
		{SmallSubTime(-2 * time.Second), "В сети", "ru"},
		{SmallSubTime(-9 * time.Second), "В сети", "ru"},
		{SmallSubTime(-10 * time.Second), "В сети", "ru"},
		{SmallSubTime(-11 * time.Second), "В сети", "ru"},
		{SmallSubTime(-20 * time.Second), "В сети", "ru"},
		{SmallSubTime(-21 * time.Second), "В сети", "ru"},
		{SmallSubTime(-22 * time.Second), "В сети", "ru"},
		{SmallSubTime(-30 * time.Second), "В сети", "ru"},
		{SmallSubTime(-31 * time.Second), "В сети", "ru"},
		{SmallSubTime(-60 * time.Second), "1 минута назад", "ru"},
		{SmallSubTime(-1 * time.Minute), "1 минута назад", "ru"},
		{SmallSubTime(-2 * time.Minute), "2 минуты назад", "ru"},
		{SmallSubTime(-9 * time.Minute), "9 минут назад", "ru"},
		{SmallSubTime(-10 * time.Minute), "10 минут назад", "ru"},
		{SmallSubTime(-11 * time.Minute), "11 минут назад", "ru"},
		{SmallSubTime(-20 * time.Minute), "20 минут назад", "ru"},
		{SmallSubTime(-21 * time.Minute), "21 минута назад", "ru"},
		{SmallSubTime(-22 * time.Minute), "22 минуты назад", "ru"},
		{SmallSubTime(-30 * time.Minute), "30 минут назад", "ru"},
		{SmallSubTime(-31 * time.Minute), "31 минута назад", "ru"},
		{SmallSubTime(-60 * time.Minute), "1 час назад", "ru"},
		{SmallSubTime(-1 * time.Hour), "1 час назад", "ru"},
		{SmallSubTime(-2 * time.Hour), "2 часа назад", "ru"},
		{SmallSubTime(-9 * time.Hour), "9 часов назад", "ru"},
		{SmallSubTime(-10 * time.Hour), "10 часов назад", "ru"},
		{SmallSubTime(-11 * time.Hour), "11 часов назад", "ru"},
		{SmallSubTime(-20 * time.Hour), "20 часов назад", "ru"},
		{SmallSubTime(-21 * time.Hour), "21 час назад", "ru"},
		{SmallSubTime(-23 * time.Hour), "23 часа назад", "ru"},
		{SmallSubTime(-24 * time.Hour), "1 день назад", "ru"},
		{SmallSubTime(-30 * time.Hour), "1 день назад", "ru"},
		{SmallSubTime((-24 * 2) * time.Hour), "2 дня назад", "ru"},
		{SmallSubTime((-24 * 6) * time.Hour), "6 дней назад", "ru"},
		{SmallSubTime((-24 * 7) * time.Hour), "1 неделя назад", "ru"},
		{SmallSubTime((-24 * 14) * time.Hour), "2 недели назад", "ru"},
		{SmallSubTime((-24 * 21) * time.Hour), "3 недели назад", "ru"},
		{BigSubTime(0, 1, 1), "1 месяц назад", "ru"},
		{BigSubTime(0, 2, 1), "2 месяца назад", "ru"},
		{BigSubTime(0, 9, 1), "9 месяцев назад", "ru"},
		{BigSubTime(0, 11, 1), "11 месяцев назад", "ru"},
		{BigSubTime(0, 12, 1), "1 год назад", "ru"},
		{BigSubTime(1, 0, 1), "1 год назад", "ru"},
		{BigSubTime(2, 0, 1), "2 года назад", "ru"},
		{BigSubTime(5, 0, 1), "5 лет назад", "ru"},
		{BigSubTime(6, 0, 1), "6 лет назад", "ru"},
		{BigSubTime(7, 0, 1), "7 лет назад", "ru"},
		{BigSubTime(21, 0, 1), "21 год назад", "ru"},
		{BigSubTime(31, 0, 1), "31 год назад", "ru"},
		{BigSubTime(100, 0, 1), "100 лет назад", "ru"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)

			if res := timeago.Parse(tc.date, "online"); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseRuWithSeconds(t *testing.T) {
	cases := []struct {
		date   string
		result []string
		lang   string
	}{
		{SmallSubTime(time.Second * 2), []string{"0 секунд назад", "1 секунда назад"}, "ru"},
		{SmallSubTime(time.Second), []string{"0 секунд назад", "1 секунда назад"}, "ru"},
		{SmallSubTime(-1 * time.Second), []string{"1 секунда назад", "2 секунды назад"}, "ru"},
		{SmallSubTime(-2 * time.Second), []string{"2 секунды назад", "3 секунды назад"}, "ru"},
		{SmallSubTime(-3 * time.Second), []string{"3 секунды назад", "4 секунды назад"}, "ru"},
		{SmallSubTime(-4 * time.Second), []string{"4 секунды назад", "5 секунд назад"}, "ru"},
		{SmallSubTime(-9 * time.Second), []string{"9 секунд назад", "10 секунд назад"}, "ru"},
		{SmallSubTime(-10 * time.Second), []string{"10 секунд назад", "11 секунд назад"}, "ru"},
		{SmallSubTime(-11 * time.Second), []string{"11 секунд назад", "12 секунд назад"}, "ru"},
		{SmallSubTime(-29 * time.Second), []string{"29 секунд назад", "30 секунд назад"}, "ru"},
		{SmallSubTime(-30 * time.Second), []string{"30 секунд назад", "31 секунда назад"}, "ru"},
		{SmallSubTime(-31 * time.Second), []string{"31 секунда назад", "32 секунды назад"}, "ru"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)
			timeago.Set("location", "Europe/Kiev")

			if res := timeago.Parse(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
