package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago"
)

func TestTakeUk(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{smallSubTime(-60 * time.Second), "1 хвилина назад", "uk"},
		{smallSubTime(-1 * time.Minute), "1 хвилина назад", "uk"},
		{smallSubTime(-2 * time.Minute), "2 хвилини назад", "uk"},
		{smallSubTime(-5 * time.Minute), "5 хвилин назад", "uk"},
		{smallSubTime(-9 * time.Minute), "9 хвилин назад", "uk"},
		{smallSubTime(-10 * time.Minute), "10 хвилин назад", "uk"},
		{smallSubTime(-11 * time.Minute), "11 хвилин назад", "uk"},
		{smallSubTime(-20 * time.Minute), "20 хвилин назад", "uk"},
		{smallSubTime(-21 * time.Minute), "21 хвилина назад", "uk"},
		{smallSubTime(-22 * time.Minute), "22 хвилини назад", "uk"},
		{smallSubTime(-30 * time.Minute), "30 хвилин назад", "uk"},
		{smallSubTime(-31 * time.Minute), "31 хвилина назад", "uk"},
		{smallSubTime(-59 * time.Minute), "59 хвилин назад", "uk"},
		{smallSubTime(-60 * time.Minute), "1 година назад", "uk"},
		{smallSubTime(-1 * time.Hour), "1 година назад", "uk"},
		{smallSubTime(-2 * time.Hour), "2 години назад", "uk"},
		{smallSubTime(-9 * time.Hour), "9 годин назад", "uk"},
		{smallSubTime(-10 * time.Hour), "10 годин назад", "uk"},
		{smallSubTime(-11 * time.Hour), "11 годин назад", "uk"},
		{smallSubTime(-20 * time.Hour), "20 годин назад", "uk"},
		{smallSubTime(-21 * time.Hour), "21 година назад", "uk"},
		{smallSubTime(-23 * time.Hour), "23 години назад", "uk"},
		{smallSubTime(-24 * time.Hour), "1 день назад", "uk"},
		{smallSubTime(-30 * time.Hour), "1 день назад", "uk"},
		{smallSubTime((-24 * 2) * time.Hour), "2 дня назад", "uk"},
		{smallSubTime((-24 * 6) * time.Hour), "6 днів назад", "uk"},
		{smallSubTime((-24 * 7) * time.Hour), "1 тиждень назад", "uk"},
		{smallSubTime((-24 * 14) * time.Hour), "2 тижні назад", "uk"},
		{smallSubTime((-24 * 21) * time.Hour), "3 тижні назад", "uk"},
		{bigSubTime(0, 1, 1), "1 місяць назад", "uk"},
		{bigSubTime(0, 2, 1), "2 місяці назад", "uk"},
		{bigSubTime(0, 9, 1), "9 місяців назад", "uk"},
		{bigSubTime(0, 11, 1), "11 місяців назад", "uk"},
		{bigSubTime(0, 12, 1), "1 рік назад", "uk"},
		{bigSubTime(1, 0, 1), "1 рік назад", "uk"},
		{bigSubTime(2, 0, 1), "2 року назад", "uk"},
		{bigSubTime(5, 0, 1), "5 років назад", "uk"},
		{bigSubTime(6, 0, 1), "6 років назад", "uk"},
		{bigSubTime(7, 0, 1), "7 років назад", "uk"},
		{bigSubTime(21, 0, 1), "21 рік назад", "uk"},
		{bigSubTime(31, 0, 1), "31 рік назад", "uk"},
		{bigSubTime(100, 0, 1), "100 років назад", "uk"},
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

func TestTakeUkWithOnlineFlag(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{smallSubTime(time.Second * 2), "В мережі", "uk"},
		{smallSubTime(time.Second), "В мережі", "uk"},
		{smallSubTime(-1 * time.Second), "В мережі", "uk"},
		{smallSubTime(-2 * time.Second), "В мережі", "uk"},
		{smallSubTime(-9 * time.Second), "В мережі", "uk"},
		{smallSubTime(-10 * time.Second), "В мережі", "uk"},
		{smallSubTime(-11 * time.Second), "В мережі", "uk"},
		{smallSubTime(-20 * time.Second), "В мережі", "uk"},
		{smallSubTime(-21 * time.Second), "В мережі", "uk"},
		{smallSubTime(-22 * time.Second), "В мережі", "uk"},
		{smallSubTime(-30 * time.Second), "В мережі", "uk"},
		{smallSubTime(-31 * time.Second), "В мережі", "uk"},
		{smallSubTime(-60 * time.Second), "1 хвилина назад", "uk"},
		{smallSubTime(-1 * time.Minute), "1 хвилина назад", "uk"},
		{smallSubTime(-2 * time.Minute), "2 хвилини назад", "uk"},
		{smallSubTime(-9 * time.Minute), "9 хвилин назад", "uk"},
		{smallSubTime(-10 * time.Minute), "10 хвилин назад", "uk"},
		{smallSubTime(-11 * time.Minute), "11 хвилин назад", "uk"},
		{smallSubTime(-20 * time.Minute), "20 хвилин назад", "uk"},
		{smallSubTime(-21 * time.Minute), "21 хвилина назад", "uk"},
		{smallSubTime(-22 * time.Minute), "22 хвилини назад", "uk"},
		{smallSubTime(-30 * time.Minute), "30 хвилин назад", "uk"},
		{smallSubTime(-31 * time.Minute), "31 хвилина назад", "uk"},
		{smallSubTime(-60 * time.Minute), "1 година назад", "uk"},
		{smallSubTime(-1 * time.Hour), "1 година назад", "uk"},
		{smallSubTime(-2 * time.Hour), "2 години назад", "uk"},
		{smallSubTime(-9 * time.Hour), "9 годин назад", "uk"},
		{smallSubTime(-10 * time.Hour), "10 годин назад", "uk"},
		{smallSubTime(-11 * time.Hour), "11 годин назад", "uk"},
		{smallSubTime(-20 * time.Hour), "20 годин назад", "uk"},
		{smallSubTime(-21 * time.Hour), "21 година назад", "uk"},
		{smallSubTime(-23 * time.Hour), "23 години назад", "uk"},
		{smallSubTime(-24 * time.Hour), "1 день назад", "uk"},
		{smallSubTime(-30 * time.Hour), "1 день назад", "uk"},
		{smallSubTime((-24 * 2) * time.Hour), "2 дня назад", "uk"},
		{smallSubTime((-24 * 6) * time.Hour), "6 днів назад", "uk"},
		{smallSubTime((-24 * 7) * time.Hour), "1 тиждень назад", "uk"},
		{smallSubTime((-24 * 14) * time.Hour), "2 тижні назад", "uk"},
		{smallSubTime((-24 * 21) * time.Hour), "3 тижні назад", "uk"},
		{bigSubTime(0, 1, 1), "1 місяць назад", "uk"},
		{bigSubTime(0, 2, 1), "2 місяці назад", "uk"},
		{bigSubTime(0, 9, 1), "9 місяців назад", "uk"},
		{bigSubTime(0, 11, 1), "11 місяців назад", "uk"},
		{bigSubTime(0, 12, 1), "1 рік назад", "uk"},
		{bigSubTime(1, 0, 1), "1 рік назад", "uk"},
		{bigSubTime(2, 0, 1), "2 року назад", "uk"},
		{bigSubTime(5, 0, 1), "5 років назад", "uk"},
		{bigSubTime(6, 0, 1), "6 років назад", "uk"},
		{bigSubTime(7, 0, 1), "7 років назад", "uk"},
		{bigSubTime(21, 0, 1), "21 рік назад", "uk"},
		{bigSubTime(31, 0, 1), "31 рік назад", "uk"},
		{bigSubTime(100, 0, 1), "100 років назад", "uk"},
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

func TestTakeUkWithSeconds(t *testing.T) {
	cases := []struct {
		date   string
		result []string
		lang   string
	}{
		{smallSubTime(time.Second * 2), []string{"0 секунд назад", "1 секунда назад"}, "uk"},
		{smallSubTime(time.Second), []string{"0 секунд назад", "1 секунда назад"}, "uk"},
		{smallSubTime(-1 * time.Second), []string{"1 секунда назад", "2 секунди назад"}, "uk"},
		{smallSubTime(-2 * time.Second), []string{"2 секунди назад", "3 секунди назад"}, "uk"},
		{smallSubTime(-3 * time.Second), []string{"3 секунди назад", "4 секунди назад"}, "uk"},
		{smallSubTime(-4 * time.Second), []string{"4 секунди назад", "5 секунд назад"}, "uk"},
		{smallSubTime(-9 * time.Second), []string{"9 секунд назад", "10 секунд назад"}, "uk"},
		{smallSubTime(-10 * time.Second), []string{"10 секунд назад", "11 секунд назад"}, "uk"},
		{smallSubTime(-11 * time.Second), []string{"11 секунд назад", "12 секунд назад"}, "uk"},
		{smallSubTime(-29 * time.Second), []string{"29 секунд назад", "30 секунд назад"}, "uk"},
		{smallSubTime(-30 * time.Second), []string{"30 секунд назад", "31 секунда назад"}, "uk"},
		{smallSubTime(-31 * time.Second), []string{"31 секунда назад", "32 секунди назад"}, "uk"},
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
