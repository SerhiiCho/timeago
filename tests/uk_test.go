package tests

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago"
	. "github.com/SerhiiCho/timeago/utils"
)

func TestConvUk(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{SmallSubTime(-60 * time.Second), "1 хвилина назад", "uk"},
		{SmallSubTime(-1 * time.Minute), "1 хвилина назад", "uk"},
		{SmallSubTime(-2 * time.Minute), "2 хвилини назад", "uk"},
		{SmallSubTime(-5 * time.Minute), "5 хвилин назад", "uk"},
		{SmallSubTime(-9 * time.Minute), "9 хвилин назад", "uk"},
		{SmallSubTime(-10 * time.Minute), "10 хвилин назад", "uk"},
		{SmallSubTime(-11 * time.Minute), "11 хвилин назад", "uk"},
		{SmallSubTime(-20 * time.Minute), "20 хвилин назад", "uk"},
		{SmallSubTime(-21 * time.Minute), "21 хвилина назад", "uk"},
		{SmallSubTime(-22 * time.Minute), "22 хвилини назад", "uk"},
		{SmallSubTime(-30 * time.Minute), "30 хвилин назад", "uk"},
		{SmallSubTime(-31 * time.Minute), "31 хвилина назад", "uk"},
		{SmallSubTime(-59 * time.Minute), "59 хвилин назад", "uk"},
		{SmallSubTime(-60 * time.Minute), "1 година назад", "uk"},
		{SmallSubTime(-1 * time.Hour), "1 година назад", "uk"},
		{SmallSubTime(-2 * time.Hour), "2 години назад", "uk"},
		{SmallSubTime(-9 * time.Hour), "9 годин назад", "uk"},
		{SmallSubTime(-10 * time.Hour), "10 годин назад", "uk"},
		{SmallSubTime(-11 * time.Hour), "11 годин назад", "uk"},
		{SmallSubTime(-20 * time.Hour), "20 годин назад", "uk"},
		{SmallSubTime(-21 * time.Hour), "21 година назад", "uk"},
		{SmallSubTime(-23 * time.Hour), "23 години назад", "uk"},
		{SmallSubTime(-24 * time.Hour), "1 день назад", "uk"},
		{SmallSubTime(-30 * time.Hour), "1 день назад", "uk"},
		{SmallSubTime((-24 * 2) * time.Hour), "2 дня назад", "uk"},
		{SmallSubTime((-24 * 6) * time.Hour), "6 днів назад", "uk"},
		{SmallSubTime((-24 * 7) * time.Hour), "1 тиждень назад", "uk"},
		{SmallSubTime((-24 * 14) * time.Hour), "2 тижні назад", "uk"},
		{SmallSubTime((-24 * 21) * time.Hour), "3 тижні назад", "uk"},
		{BigSubTime(0, 1, 1), "1 місяць назад", "uk"},
		{BigSubTime(0, 2, 1), "2 місяці назад", "uk"},
		{BigSubTime(0, 9, 1), "9 місяців назад", "uk"},
		{BigSubTime(0, 11, 1), "11 місяців назад", "uk"},
		{BigSubTime(0, 12, 1), "1 рік назад", "uk"},
		{BigSubTime(1, 0, 1), "1 рік назад", "uk"},
		{BigSubTime(2, 0, 1), "2 року назад", "uk"},
		{BigSubTime(5, 0, 1), "5 років назад", "uk"},
		{BigSubTime(6, 0, 1), "6 років назад", "uk"},
		{BigSubTime(7, 0, 1), "7 років назад", "uk"},
		{BigSubTime(21, 0, 1), "21 рік назад", "uk"},
		{BigSubTime(31, 0, 1), "31 рік назад", "uk"},
		{BigSubTime(100, 0, 1), "100 років назад", "uk"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)
			timeago.Set("location", "Europe/Kiev")

			if res := timeago.Conv(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestConvUkWithOnlineFlag(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{SmallSubTime(time.Second * 2), "В мережі", "uk"},
		{SmallSubTime(time.Second), "В мережі", "uk"},
		{SmallSubTime(-1 * time.Second), "В мережі", "uk"},
		{SmallSubTime(-2 * time.Second), "В мережі", "uk"},
		{SmallSubTime(-9 * time.Second), "В мережі", "uk"},
		{SmallSubTime(-10 * time.Second), "В мережі", "uk"},
		{SmallSubTime(-11 * time.Second), "В мережі", "uk"},
		{SmallSubTime(-20 * time.Second), "В мережі", "uk"},
		{SmallSubTime(-21 * time.Second), "В мережі", "uk"},
		{SmallSubTime(-22 * time.Second), "В мережі", "uk"},
		{SmallSubTime(-30 * time.Second), "В мережі", "uk"},
		{SmallSubTime(-31 * time.Second), "В мережі", "uk"},
		{SmallSubTime(-60 * time.Second), "1 хвилина назад", "uk"},
		{SmallSubTime(-1 * time.Minute), "1 хвилина назад", "uk"},
		{SmallSubTime(-2 * time.Minute), "2 хвилини назад", "uk"},
		{SmallSubTime(-9 * time.Minute), "9 хвилин назад", "uk"},
		{SmallSubTime(-10 * time.Minute), "10 хвилин назад", "uk"},
		{SmallSubTime(-11 * time.Minute), "11 хвилин назад", "uk"},
		{SmallSubTime(-20 * time.Minute), "20 хвилин назад", "uk"},
		{SmallSubTime(-21 * time.Minute), "21 хвилина назад", "uk"},
		{SmallSubTime(-22 * time.Minute), "22 хвилини назад", "uk"},
		{SmallSubTime(-30 * time.Minute), "30 хвилин назад", "uk"},
		{SmallSubTime(-31 * time.Minute), "31 хвилина назад", "uk"},
		{SmallSubTime(-60 * time.Minute), "1 година назад", "uk"},
		{SmallSubTime(-1 * time.Hour), "1 година назад", "uk"},
		{SmallSubTime(-2 * time.Hour), "2 години назад", "uk"},
		{SmallSubTime(-9 * time.Hour), "9 годин назад", "uk"},
		{SmallSubTime(-10 * time.Hour), "10 годин назад", "uk"},
		{SmallSubTime(-11 * time.Hour), "11 годин назад", "uk"},
		{SmallSubTime(-20 * time.Hour), "20 годин назад", "uk"},
		{SmallSubTime(-21 * time.Hour), "21 година назад", "uk"},
		{SmallSubTime(-23 * time.Hour), "23 години назад", "uk"},
		{SmallSubTime(-24 * time.Hour), "1 день назад", "uk"},
		{SmallSubTime(-30 * time.Hour), "1 день назад", "uk"},
		{SmallSubTime((-24 * 2) * time.Hour), "2 дня назад", "uk"},
		{SmallSubTime((-24 * 6) * time.Hour), "6 днів назад", "uk"},
		{SmallSubTime((-24 * 7) * time.Hour), "1 тиждень назад", "uk"},
		{SmallSubTime((-24 * 14) * time.Hour), "2 тижні назад", "uk"},
		{SmallSubTime((-24 * 21) * time.Hour), "3 тижні назад", "uk"},
		{BigSubTime(0, 1, 1), "1 місяць назад", "uk"},
		{BigSubTime(0, 2, 1), "2 місяці назад", "uk"},
		{BigSubTime(0, 9, 1), "9 місяців назад", "uk"},
		{BigSubTime(0, 11, 1), "11 місяців назад", "uk"},
		{BigSubTime(0, 12, 1), "1 рік назад", "uk"},
		{BigSubTime(1, 0, 1), "1 рік назад", "uk"},
		{BigSubTime(2, 0, 1), "2 року назад", "uk"},
		{BigSubTime(5, 0, 1), "5 років назад", "uk"},
		{BigSubTime(6, 0, 1), "6 років назад", "uk"},
		{BigSubTime(7, 0, 1), "7 років назад", "uk"},
		{BigSubTime(21, 0, 1), "21 рік назад", "uk"},
		{BigSubTime(31, 0, 1), "31 рік назад", "uk"},
		{BigSubTime(100, 0, 1), "100 років назад", "uk"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)

			if res := timeago.Conv(tc.date, "online"); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestConvUkWithSeconds(t *testing.T) {
	cases := []struct {
		date   string
		result []string
		lang   string
	}{
		{SmallSubTime(time.Second * 2), []string{"0 секунд назад", "1 секунда назад"}, "uk"},
		{SmallSubTime(time.Second), []string{"0 секунд назад", "1 секунда назад"}, "uk"},
		{SmallSubTime(-1 * time.Second), []string{"1 секунда назад", "2 секунди назад"}, "uk"},
		{SmallSubTime(-2 * time.Second), []string{"2 секунди назад", "3 секунди назад"}, "uk"},
		{SmallSubTime(-3 * time.Second), []string{"3 секунди назад", "4 секунди назад"}, "uk"},
		{SmallSubTime(-4 * time.Second), []string{"4 секунди назад", "5 секунд назад"}, "uk"},
		{SmallSubTime(-9 * time.Second), []string{"9 секунд назад", "10 секунд назад"}, "uk"},
		{SmallSubTime(-10 * time.Second), []string{"10 секунд назад", "11 секунд назад"}, "uk"},
		{SmallSubTime(-11 * time.Second), []string{"11 секунд назад", "12 секунд назад"}, "uk"},
		{SmallSubTime(-29 * time.Second), []string{"29 секунд назад", "30 секунд назад"}, "uk"},
		{SmallSubTime(-30 * time.Second), []string{"30 секунд назад", "31 секунда назад"}, "uk"},
		{SmallSubTime(-31 * time.Second), []string{"31 секунда назад", "32 секунди назад"}, "uk"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			timeago.Set("language", tc.lang)
			timeago.Set("location", "Europe/Kiev")

			if res := timeago.Conv(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
