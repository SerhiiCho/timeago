package timeago

import (
	"testing"
	"time"
)

func TestGetOption(test *testing.T) {
	test.Run("case has online option", func(t *testing.T) {
		option, hasOption := getOption("2017-02-01 00:00:00|online")

		if !hasOption || option != "online" {
			t.Errorf("Result of getOption func must return `online` string and `true`, but `%s` string and `%v` returned returned", option, hasOption)
		}
	})

	test.Run("case has random option", func(t *testing.T) {
		option, hasOption := getOption("2017-02-01 00:00:00|random")

		if !hasOption || option != "random" {
			t.Errorf("Result of getOption func must return `random` string and `true`, but `%s` string and `%v` returned returned", option, hasOption)
		}
	})

	test.Run("case don't have option", func(t *testing.T) {
		option, hasOption := getOption("2017-02-01 00:00:00")

		if hasOption || option != "" {
			t.Errorf("Result of getOption func must return empty string and `false`, but `%s` string and `%v` returned returned", option, hasOption)
		}
	})
}

func TestGetTimeTranslations(t *testing.T) {
	t.Parallel()

	t.Run("returns 7 items", func(test *testing.T) {
		if res := len(getTimeTranslations()); res != 7 {
			test.Errorf("Must return 7 items but got %d", res)
		}
	})

	t.Run("every slice has 3 items", func(test *testing.T) {
		for _, slice := range getTimeTranslations() {
			if res := len(slice); res != 3 {
				test.Errorf("Slice Must return 3 items but got %d", res)
			}
		}
	})
}

func TestGetWords(t *testing.T) {
	cases := []struct {
		timeKind string
		num      int
		result   string
		lang     string
	}{
		// english
		{"seconds", 1, "1 second ago", "en"},
		{"seconds", 2, "2 seconds ago", "en"},
		{"days", 11, "11 days ago", "en"},
		{"days", 21, "21 day ago", "en"},
		{"seconds", 30, "30 seconds ago", "en"},
		{"seconds", 31, "31 second ago", "en"},
		{"hours", 10, "10 hours ago", "en"},
		{"years", 2, "2 years ago", "en"},
		// russian
		{"seconds", 1, "1 секунда назад", "ru"},
		{"minutes", 2, "2 минуты назад", "ru"},
		{"seconds", 3, "3 секунды назад", "ru"},
		{"seconds", 4, "4 секунды назад", "ru"},
		{"hours", 5, "5 часов назад", "ru"},
		{"days", 11, "11 дней назад", "ru"},
		{"years", 21, "21 год назад", "ru"},
		{"minutes", 59, "59 минут назад", "ru"},
	}

	for _, tc := range cases {
		t.Run(tc.result, func(test *testing.T) {
			SetLang(tc.lang)

			if res := getWords(tc.timeKind, tc.num); res != tc.result {
				test.Errorf("Result must be `%s` but got `%s` instead", tc.result, res)
			}
		})
	}
}

func TestGetLastNumber(t *testing.T) {
	cases := []struct {
		number int
		result int
		name   string
	}{
		{0, 0, "must return 0"},
		{1, 1, "must return 1"},
		{9, 9, "must return 9"},
		{20, 0, "must return 0"},
		{253, 3, "must return 3"},
		{23423252, 2, "must return 2"},
		{223424342325, 5, "must return 5"},
		{23423521562348, 8, "must return 8"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(test *testing.T) {
			if res := getLastNumber(tc.number); res != tc.result {
				test.Errorf("Result must be %d, but got %d instead", tc.result, res)
			}
		})
	}
}

func TestTake(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		{time.Now().UTC().Add(-10 * time.Second).Format("2006-01-02 15:04:05"), "10 seconds ago", "en"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			SetLang(tc.lang)

			if res := Take(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}
