package tests

import (
	"testing"
	"time"

	. "github.com/SerhiiCho/timeago"
)

func TestParseUk(t *testing.T) {
	cases := []struct {
		date   string
		result string
	}{
		{smallSubTime(-60 * time.Second), "1 хвилина назад"},
		{smallSubTime(-1 * time.Minute), "1 хвилина назад"},
		{smallSubTime(-2 * time.Minute), "2 хвилини назад"},
		{smallSubTime(-5 * time.Minute), "5 хвилин назад"},
		{smallSubTime(-9 * time.Minute), "9 хвилин назад"},
		{smallSubTime(-10 * time.Minute), "10 хвилин назад"},
		{smallSubTime(-11 * time.Minute), "11 хвилин назад"},
		{smallSubTime(-20 * time.Minute), "20 хвилин назад"},
		{smallSubTime(-21 * time.Minute), "21 хвилина назад"},
		{smallSubTime(-22 * time.Minute), "22 хвилини назад"},
		{smallSubTime(-30 * time.Minute), "30 хвилин назад"},
		{smallSubTime(-31 * time.Minute), "31 хвилина назад"},
		{smallSubTime(-59 * time.Minute), "59 хвилин назад"},
		{smallSubTime(-60 * time.Minute), "1 година назад"},
		{smallSubTime(-1 * time.Hour), "1 година назад"},
		{smallSubTime(-2 * time.Hour), "2 години назад"},
		{smallSubTime(-9 * time.Hour), "9 годин назад"},
		{smallSubTime(-10 * time.Hour), "10 годин назад"},
		{smallSubTime(-11 * time.Hour), "11 годин назад"},
		{smallSubTime(-20 * time.Hour), "20 годин назад"},
		{smallSubTime(-21 * time.Hour), "21 година назад"},
		{smallSubTime(-23 * time.Hour), "23 години назад"},
		{smallSubTime(-24 * time.Hour), "1 день назад"},
		{smallSubTime(-30 * time.Hour), "1 день назад"},
		{smallSubTime((-24 * 2) * time.Hour), "2 дня назад"},
		{smallSubTime((-24 * 6) * time.Hour), "6 днів назад"},
		{smallSubTime((-24 * 7) * time.Hour), "1 тиждень назад"},
		{smallSubTime((-24 * 14) * time.Hour), "2 тижні назад"},
		{smallSubTime((-24 * 21) * time.Hour), "3 тижні назад"},
		{bigSubTime(0, 1, 1), "1 місяць назад"},
		{bigSubTime(0, 2, 1), "2 місяці назад"},
		{bigSubTime(0, 9, 1), "9 місяців назад"},
		{bigSubTime(0, 11, 1), "11 місяців назад"},
		{bigSubTime(0, 12, 1), "1 рік назад"},
		{bigSubTime(1, 0, 1), "1 рік назад"},
		{bigSubTime(2, 0, 1), "2 року назад"},
		{bigSubTime(5, 0, 1), "5 років назад"},
		{bigSubTime(6, 0, 1), "6 років назад"},
		{bigSubTime(7, 0, 1), "7 років назад"},
		{bigSubTime(21, 0, 1), "21 рік назад"},
		{bigSubTime(31, 0, 1), "31 рік назад"},
		{bigSubTime(100, 0, 1), "100 років назад"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			SetConfig(Config{
				Language: "uk",
				Location: "Europe/Kiev",
			})

			if res := Parse(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseUkWithOnlineFlag(t *testing.T) {
	cases := []struct {
		date   string
		result string
	}{
		{smallSubTime(time.Second * 2), "В мережі"},
		{smallSubTime(time.Second), "В мережі"},
		{smallSubTime(-1 * time.Second), "В мережі"},
		{smallSubTime(-2 * time.Second), "В мережі"},
		{smallSubTime(-9 * time.Second), "В мережі"},
		{smallSubTime(-10 * time.Second), "В мережі"},
		{smallSubTime(-11 * time.Second), "В мережі"},
		{smallSubTime(-20 * time.Second), "В мережі"},
		{smallSubTime(-21 * time.Second), "В мережі"},
		{smallSubTime(-22 * time.Second), "В мережі"},
		{smallSubTime(-30 * time.Second), "В мережі"},
		{smallSubTime(-31 * time.Second), "В мережі"},
		{smallSubTime(-60 * time.Second), "1 хвилина назад"},
		{smallSubTime(-1 * time.Minute), "1 хвилина назад"},
		{smallSubTime(-2 * time.Minute), "2 хвилини назад"},
		{smallSubTime(-9 * time.Minute), "9 хвилин назад"},
		{smallSubTime(-10 * time.Minute), "10 хвилин назад"},
		{smallSubTime(-11 * time.Minute), "11 хвилин назад"},
		{smallSubTime(-20 * time.Minute), "20 хвилин назад"},
		{smallSubTime(-21 * time.Minute), "21 хвилина назад"},
		{smallSubTime(-22 * time.Minute), "22 хвилини назад"},
		{smallSubTime(-30 * time.Minute), "30 хвилин назад"},
		{smallSubTime(-31 * time.Minute), "31 хвилина назад"},
		{smallSubTime(-60 * time.Minute), "1 година назад"},
		{smallSubTime(-1 * time.Hour), "1 година назад"},
		{smallSubTime(-2 * time.Hour), "2 години назад"},
		{smallSubTime(-9 * time.Hour), "9 годин назад"},
		{smallSubTime(-10 * time.Hour), "10 годин назад"},
		{smallSubTime(-11 * time.Hour), "11 годин назад"},
		{smallSubTime(-20 * time.Hour), "20 годин назад"},
		{smallSubTime(-21 * time.Hour), "21 година назад"},
		{smallSubTime(-23 * time.Hour), "23 години назад"},
		{smallSubTime(-24 * time.Hour), "1 день назад"},
		{smallSubTime(-30 * time.Hour), "1 день назад"},
		{smallSubTime((-24 * 2) * time.Hour), "2 дня назад"},
		{smallSubTime((-24 * 6) * time.Hour), "6 днів назад"},
		{smallSubTime((-24 * 7) * time.Hour), "1 тиждень назад"},
		{smallSubTime((-24 * 14) * time.Hour), "2 тижні назад"},
		{smallSubTime((-24 * 21) * time.Hour), "3 тижні назад"},
		{bigSubTime(0, 1, 1), "1 місяць назад"},
		{bigSubTime(0, 2, 1), "2 місяці назад"},
		{bigSubTime(0, 9, 1), "9 місяців назад"},
		{bigSubTime(0, 11, 1), "11 місяців назад"},
		{bigSubTime(0, 12, 1), "1 рік назад"},
		{bigSubTime(1, 0, 1), "1 рік назад"},
		{bigSubTime(2, 0, 1), "2 року назад"},
		{bigSubTime(5, 0, 1), "5 років назад"},
		{bigSubTime(6, 0, 1), "6 років назад"},
		{bigSubTime(7, 0, 1), "7 років назад"},
		{bigSubTime(21, 0, 1), "21 рік назад"},
		{bigSubTime(31, 0, 1), "31 рік назад"},
		{bigSubTime(100, 0, 1), "100 років назад"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {

			SetConfig(Config{
				Language: "uk",
			})

			if res := Parse(tc.date, "online"); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseUkWithSeconds(t *testing.T) {
	cases := []struct {
		date   string
		result []string
	}{
		{smallSubTime(time.Second * 2), []string{"0 секунд назад", "1 секунда назад"}},
		{smallSubTime(time.Second), []string{"0 секунд назад", "1 секунда назад"}},
		{smallSubTime(-1 * time.Second), []string{"1 секунда назад", "2 секунди назад"}},
		{smallSubTime(-2 * time.Second), []string{"2 секунди назад", "3 секунди назад"}},
		{smallSubTime(-3 * time.Second), []string{"3 секунди назад", "4 секунди назад"}},
		{smallSubTime(-4 * time.Second), []string{"4 секунди назад", "5 секунд назад"}},
		{smallSubTime(-9 * time.Second), []string{"9 секунд назад", "10 секунд назад"}},
		{smallSubTime(-10 * time.Second), []string{"10 секунд назад", "11 секунд назад"}},
		{smallSubTime(-11 * time.Second), []string{"11 секунд назад", "12 секунд назад"}},
		{smallSubTime(-29 * time.Second), []string{"29 секунд назад", "30 секунд назад"}},
		{smallSubTime(-30 * time.Second), []string{"30 секунд назад", "31 секунда назад"}},
		{smallSubTime(-31 * time.Second), []string{"31 секунда назад", "32 секунди назад"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			SetConfig(Config{
				Language: "uk",
				Location: "Europe/Kiev",
			})

			if res := Parse(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
