package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago"
)

func TestTakeRu(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{smallSubTime(-60 * time.Second), "1 минута назад", "ru"},
		{smallSubTime(-1 * time.Minute), "1 минута назад", "ru"},
		{smallSubTime(-2 * time.Minute), "2 минуты назад", "ru"},
		{smallSubTime(-5 * time.Minute), "5 минут назад", "ru"},
		{smallSubTime(-9 * time.Minute), "9 минут назад", "ru"},
		{smallSubTime(-10 * time.Minute), "10 минут назад", "ru"},
		{smallSubTime(-11 * time.Minute), "11 минут назад", "ru"},
		{smallSubTime(-20 * time.Minute), "20 минут назад", "ru"},
		{smallSubTime(-21 * time.Minute), "21 минута назад", "ru"},
		{smallSubTime(-22 * time.Minute), "22 минуты назад", "ru"},
		{smallSubTime(-30 * time.Minute), "30 минут назад", "ru"},
		{smallSubTime(-31 * time.Minute), "31 минута назад", "ru"},
		{smallSubTime(-59 * time.Minute), "59 минут назад", "ru"},
		{smallSubTime(-60 * time.Minute), "1 час назад", "ru"},
		{smallSubTime(-1 * time.Hour), "1 час назад", "ru"},
		{smallSubTime(-2 * time.Hour), "2 часа назад", "ru"},
		{smallSubTime(-9 * time.Hour), "9 часов назад", "ru"},
		{smallSubTime(-10 * time.Hour), "10 часов назад", "ru"},
		{smallSubTime(-11 * time.Hour), "11 часов назад", "ru"},
		{smallSubTime(-20 * time.Hour), "20 часов назад", "ru"},
		{smallSubTime(-21 * time.Hour), "21 час назад", "ru"},
		{smallSubTime(-23 * time.Hour), "23 часа назад", "ru"},
		{smallSubTime(-24 * time.Hour), "1 день назад", "ru"},
		{smallSubTime(-30 * time.Hour), "1 день назад", "ru"},
		{smallSubTime((-24 * 2) * time.Hour), "2 дня назад", "ru"},
		{smallSubTime((-24 * 6) * time.Hour), "6 дней назад", "ru"},
		{smallSubTime((-24 * 7) * time.Hour), "1 неделя назад", "ru"},
		{smallSubTime((-24 * 14) * time.Hour), "2 недели назад", "ru"},
		{smallSubTime((-24 * 21) * time.Hour), "3 недели назад", "ru"},
		{bigSubTime(0, 1, 1), "1 месяц назад", "ru"},
		{bigSubTime(0, 2, 1), "2 месяца назад", "ru"},
		{bigSubTime(0, 9, 1), "9 месяцев назад", "ru"},
		{bigSubTime(0, 11, 1), "11 месяцев назад", "ru"},
		{bigSubTime(0, 12, 1), "1 год назад", "ru"},
		{bigSubTime(1, 0, 1), "1 год назад", "ru"},
		{bigSubTime(2, 0, 1), "2 года назад", "ru"},
		{bigSubTime(5, 0, 1), "5 лет назад", "ru"},
		{bigSubTime(6, 0, 1), "6 лет назад", "ru"},
		{bigSubTime(7, 0, 1), "7 лет назад", "ru"},
		{bigSubTime(21, 0, 1), "21 год назад", "ru"},
		{bigSubTime(31, 0, 1), "31 год назад", "ru"},
		{bigSubTime(100, 0, 1), "100 лет назад", "ru"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)
			timeago.Set("location", "Europe/Kiev")

			if res := timeago.Take(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestTakeRuWithOnlineFlag(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{smallSubTime(time.Second * 2), "В сети", "ru"},
		{smallSubTime(time.Second), "В сети", "ru"},
		{smallSubTime(-1 * time.Second), "В сети", "ru"},
		{smallSubTime(-2 * time.Second), "В сети", "ru"},
		{smallSubTime(-9 * time.Second), "В сети", "ru"},
		{smallSubTime(-10 * time.Second), "В сети", "ru"},
		{smallSubTime(-11 * time.Second), "В сети", "ru"},
		{smallSubTime(-20 * time.Second), "В сети", "ru"},
		{smallSubTime(-21 * time.Second), "В сети", "ru"},
		{smallSubTime(-22 * time.Second), "В сети", "ru"},
		{smallSubTime(-30 * time.Second), "В сети", "ru"},
		{smallSubTime(-31 * time.Second), "В сети", "ru"},
		{smallSubTime(-60 * time.Second), "1 минута назад", "ru"},
		{smallSubTime(-1 * time.Minute), "1 минута назад", "ru"},
		{smallSubTime(-2 * time.Minute), "2 минуты назад", "ru"},
		{smallSubTime(-9 * time.Minute), "9 минут назад", "ru"},
		{smallSubTime(-10 * time.Minute), "10 минут назад", "ru"},
		{smallSubTime(-11 * time.Minute), "11 минут назад", "ru"},
		{smallSubTime(-20 * time.Minute), "20 минут назад", "ru"},
		{smallSubTime(-21 * time.Minute), "21 минута назад", "ru"},
		{smallSubTime(-22 * time.Minute), "22 минуты назад", "ru"},
		{smallSubTime(-30 * time.Minute), "30 минут назад", "ru"},
		{smallSubTime(-31 * time.Minute), "31 минута назад", "ru"},
		{smallSubTime(-60 * time.Minute), "1 час назад", "ru"},
		{smallSubTime(-1 * time.Hour), "1 час назад", "ru"},
		{smallSubTime(-2 * time.Hour), "2 часа назад", "ru"},
		{smallSubTime(-9 * time.Hour), "9 часов назад", "ru"},
		{smallSubTime(-10 * time.Hour), "10 часов назад", "ru"},
		{smallSubTime(-11 * time.Hour), "11 часов назад", "ru"},
		{smallSubTime(-20 * time.Hour), "20 часов назад", "ru"},
		{smallSubTime(-21 * time.Hour), "21 час назад", "ru"},
		{smallSubTime(-23 * time.Hour), "23 часа назад", "ru"},
		{smallSubTime(-24 * time.Hour), "1 день назад", "ru"},
		{smallSubTime(-30 * time.Hour), "1 день назад", "ru"},
		{smallSubTime((-24 * 2) * time.Hour), "2 дня назад", "ru"},
		{smallSubTime((-24 * 6) * time.Hour), "6 дней назад", "ru"},
		{smallSubTime((-24 * 7) * time.Hour), "1 неделя назад", "ru"},
		{smallSubTime((-24 * 14) * time.Hour), "2 недели назад", "ru"},
		{smallSubTime((-24 * 21) * time.Hour), "3 недели назад", "ru"},
		{bigSubTime(0, 1, 1), "1 месяц назад", "ru"},
		{bigSubTime(0, 2, 1), "2 месяца назад", "ru"},
		{bigSubTime(0, 9, 1), "9 месяцев назад", "ru"},
		{bigSubTime(0, 11, 1), "11 месяцев назад", "ru"},
		{bigSubTime(0, 12, 1), "1 год назад", "ru"},
		{bigSubTime(1, 0, 1), "1 год назад", "ru"},
		{bigSubTime(2, 0, 1), "2 года назад", "ru"},
		{bigSubTime(5, 0, 1), "5 лет назад", "ru"},
		{bigSubTime(6, 0, 1), "6 лет назад", "ru"},
		{bigSubTime(7, 0, 1), "7 лет назад", "ru"},
		{bigSubTime(21, 0, 1), "21 год назад", "ru"},
		{bigSubTime(31, 0, 1), "31 год назад", "ru"},
		{bigSubTime(100, 0, 1), "100 лет назад", "ru"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date+"|online", func(test *testing.T) {
			timeago.Set("language", tc.lang)

			if res := timeago.Take(tc.date + "|online"); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestTakeRuWithSeconds(t *testing.T) {
	cases := []struct {
		date   string
		result []string
		lang   string
	}{
		{smallSubTime(time.Second * 2), []string{"0 секунд назад", "1 секунда назад"}, "ru"},
		{smallSubTime(time.Second), []string{"0 секунд назад", "1 секунда назад"}, "ru"},
		{smallSubTime(-1 * time.Second), []string{"1 секунда назад", "2 секунды назад"}, "ru"},
		{smallSubTime(-2 * time.Second), []string{"2 секунды назад", "3 секунды назад"}, "ru"},
		{smallSubTime(-3 * time.Second), []string{"3 секунды назад", "4 секунды назад"}, "ru"},
		{smallSubTime(-4 * time.Second), []string{"4 секунды назад", "5 секунд назад"}, "ru"},
		{smallSubTime(-9 * time.Second), []string{"9 секунд назад", "10 секунд назад"}, "ru"},
		{smallSubTime(-10 * time.Second), []string{"10 секунд назад", "11 секунд назад"}, "ru"},
		{smallSubTime(-11 * time.Second), []string{"11 секунд назад", "12 секунд назад"}, "ru"},
		{smallSubTime(-29 * time.Second), []string{"29 секунд назад", "30 секунд назад"}, "ru"},
		{smallSubTime(-30 * time.Second), []string{"30 секунд назад", "31 секунда назад"}, "ru"},
		{smallSubTime(-31 * time.Second), []string{"31 секунда назад", "32 секунды назад"}, "ru"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)
			timeago.Set("location", "Europe/Kiev")

			if res := timeago.Take(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
