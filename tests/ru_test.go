package tests

import (
	"testing"
	"time"

	. "github.com/SerhiiCho/timeago"
)

func TestParseRu(t *testing.T) {
	cases := []struct {
		date   string
		result string
	}{
		{smallSubTime(-60 * time.Second), "1 минута назад"},
		{smallSubTime(-1 * time.Minute), "1 минута назад"},
		{smallSubTime(-2 * time.Minute), "2 минуты назад"},
		{smallSubTime(-5 * time.Minute), "5 минут назад"},
		{smallSubTime(-9 * time.Minute), "9 минут назад"},
		{smallSubTime(-10 * time.Minute), "10 минут назад"},
		{smallSubTime(-11 * time.Minute), "11 минут назад"},
		{smallSubTime(-20 * time.Minute), "20 минут назад"},
		{smallSubTime(-21 * time.Minute), "21 минута назад"},
		{smallSubTime(-22 * time.Minute), "22 минуты назад"},
		{smallSubTime(-30 * time.Minute), "30 минут назад"},
		{smallSubTime(-31 * time.Minute), "31 минута назад"},
		{smallSubTime(-59 * time.Minute), "59 минут назад"},
		{smallSubTime(-60 * time.Minute), "1 час назад"},
		{smallSubTime(-1 * time.Hour), "1 час назад"},
		{smallSubTime(-2 * time.Hour), "2 часа назад"},
		{smallSubTime(-9 * time.Hour), "9 часов назад"},
		{smallSubTime(-10 * time.Hour), "10 часов назад"},
		{smallSubTime(-11 * time.Hour), "11 часов назад"},
		{smallSubTime(-20 * time.Hour), "20 часов назад"},
		{smallSubTime(-21 * time.Hour), "21 час назад"},
		{smallSubTime(-23 * time.Hour), "23 часа назад"},
		{smallSubTime(-24 * time.Hour), "1 день назад"},
		{smallSubTime(-30 * time.Hour), "1 день назад"},
		{smallSubTime((-24 * 2) * time.Hour), "2 дня назад"},
		{smallSubTime((-24 * 6) * time.Hour), "6 дней назад"},
		{smallSubTime((-24 * 7) * time.Hour), "1 неделя назад"},
		{smallSubTime((-24 * 14) * time.Hour), "2 недели назад"},
		{smallSubTime((-24 * 21) * time.Hour), "3 недели назад"},
		{bigSubTime(0, 1, 1), "1 месяц назад"},
		{bigSubTime(0, 2, 1), "2 месяца назад"},
		{bigSubTime(0, 9, 1), "9 месяцев назад"},
		{bigSubTime(0, 11, 1), "11 месяцев назад"},
		{bigSubTime(0, 12, 1), "1 год назад"},
		{bigSubTime(1, 0, 1), "1 год назад"},
		{bigSubTime(2, 0, 1), "2 года назад"},
		{bigSubTime(5, 0, 1), "5 лет назад"},
		{bigSubTime(6, 0, 1), "6 лет назад"},
		{bigSubTime(7, 0, 1), "7 лет назад"},
		{bigSubTime(21, 0, 1), "21 год назад"},
		{bigSubTime(31, 0, 1), "31 год назад"},
		{bigSubTime(100, 0, 1), "100 лет назад"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			SetConfig(Config{
				Language: "ru",
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
		{smallSubTime(time.Second * 2), "В сети"},
		{smallSubTime(time.Second), "В сети"},
		{smallSubTime(-1 * time.Second), "В сети"},
		{smallSubTime(-2 * time.Second), "В сети"},
		{smallSubTime(-9 * time.Second), "В сети"},
		{smallSubTime(-10 * time.Second), "В сети"},
		{smallSubTime(-11 * time.Second), "В сети"},
		{smallSubTime(-20 * time.Second), "В сети"},
		{smallSubTime(-21 * time.Second), "В сети"},
		{smallSubTime(-22 * time.Second), "В сети"},
		{smallSubTime(-30 * time.Second), "В сети"},
		{smallSubTime(-31 * time.Second), "В сети"},
		{smallSubTime(-60 * time.Second), "1 минута назад"},
		{smallSubTime(-1 * time.Minute), "1 минута назад"},
		{smallSubTime(-2 * time.Minute), "2 минуты назад"},
		{smallSubTime(-9 * time.Minute), "9 минут назад"},
		{smallSubTime(-10 * time.Minute), "10 минут назад"},
		{smallSubTime(-11 * time.Minute), "11 минут назад"},
		{smallSubTime(-20 * time.Minute), "20 минут назад"},
		{smallSubTime(-21 * time.Minute), "21 минута назад"},
		{smallSubTime(-22 * time.Minute), "22 минуты назад"},
		{smallSubTime(-30 * time.Minute), "30 минут назад"},
		{smallSubTime(-31 * time.Minute), "31 минута назад"},
		{smallSubTime(-60 * time.Minute), "1 час назад"},
		{smallSubTime(-1 * time.Hour), "1 час назад"},
		{smallSubTime(-2 * time.Hour), "2 часа назад"},
		{smallSubTime(-9 * time.Hour), "9 часов назад"},
		{smallSubTime(-10 * time.Hour), "10 часов назад"},
		{smallSubTime(-11 * time.Hour), "11 часов назад"},
		{smallSubTime(-20 * time.Hour), "20 часов назад"},
		{smallSubTime(-21 * time.Hour), "21 час назад"},
		{smallSubTime(-23 * time.Hour), "23 часа назад"},
		{smallSubTime(-24 * time.Hour), "1 день назад"},
		{smallSubTime(-30 * time.Hour), "1 день назад"},
		{smallSubTime((-24 * 2) * time.Hour), "2 дня назад"},
		{smallSubTime((-24 * 6) * time.Hour), "6 дней назад"},
		{smallSubTime((-24 * 7) * time.Hour), "1 неделя назад"},
		{smallSubTime((-24 * 14) * time.Hour), "2 недели назад"},
		{smallSubTime((-24 * 21) * time.Hour), "3 недели назад"},
		{bigSubTime(0, 1, 1), "1 месяц назад"},
		{bigSubTime(0, 2, 1), "2 месяца назад"},
		{bigSubTime(0, 9, 1), "9 месяцев назад"},
		{bigSubTime(0, 11, 1), "11 месяцев назад"},
		{bigSubTime(0, 12, 1), "1 год назад"},
		{bigSubTime(1, 0, 1), "1 год назад"},
		{bigSubTime(2, 0, 1), "2 года назад"},
		{bigSubTime(5, 0, 1), "5 лет назад"},
		{bigSubTime(6, 0, 1), "6 лет назад"},
		{bigSubTime(7, 0, 1), "7 лет назад"},
		{bigSubTime(21, 0, 1), "21 год назад"},
		{bigSubTime(31, 0, 1), "31 год назад"},
		{bigSubTime(100, 0, 1), "100 лет назад"},
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
		{smallSubTime(0), []string{"0 секунд назад", "1 секунда назад"}},
		{smallSubTime(time.Second * 2), []string{"0 секунд назад", "1 секунда назад"}},
		{smallSubTime(time.Second), []string{"0 секунд назад", "1 секунда назад"}},
		{smallSubTime(-1 * time.Second), []string{"1 секунда назад", "2 секунды назад"}},
		{smallSubTime(-2 * time.Second), []string{"2 секунды назад", "3 секунды назад"}},
		{smallSubTime(-3 * time.Second), []string{"3 секунды назад", "4 секунды назад"}},
		{smallSubTime(-4 * time.Second), []string{"4 секунды назад", "5 секунд назад"}},
		{smallSubTime(-9 * time.Second), []string{"9 секунд назад", "10 секунд назад"}},
		{smallSubTime(-10 * time.Second), []string{"10 секунд назад", "11 секунд назад"}},
		{smallSubTime(-11 * time.Second), []string{"11 секунд назад", "12 секунд назад"}},
		{smallSubTime(-29 * time.Second), []string{"29 секунд назад", "30 секунд назад"}},
		{smallSubTime(-30 * time.Second), []string{"30 секунд назад", "31 секунда назад"}},
		{smallSubTime(-31 * time.Second), []string{"31 секунда назад", "32 секунды назад"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			SetConfig(Config{
				Language: "ru",
			})

			if res := Parse(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
