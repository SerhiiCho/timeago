package timeago

import (
	"testing"
	"time"
)

// test helper
func smallSubTime(d time.Duration) string {
	return time.Now().Add(d).Format("2006-01-02 15:04:05")
}

// test helper
func bigSubTime(years int, months int, days int) string {
	return time.Now().AddDate(-years, -months, -days).Format("2006-01-02 15:04:05")
}

func TestGetOption(test *testing.T) {
	cases := []struct {
		name            string
		date            string
		optionMustBe    string
		hasOptionMustBe bool
	}{
		{"case has online option", "2017-02-01 00:00:00|online", "online", true},
		{"case has random option", "2017-02-01 00:00:00|random", "random", true},
		{"case has online option", "2017-02-01 00:00:00|korotchaeva", "korotchaeva", true},
		{"case don't have option", "2017-02-01 00:00:00", "", false},
	}

	for _, tc := range cases {
		test.Run(tc.name, func(t *testing.T) {
			option, hasOption := getOption(&tc.date)

			if hasOption != tc.hasOptionMustBe || option != tc.optionMustBe {
				t.Errorf("Result of getOption func must return `online` string and `true`, but `%s` string and `%v` returned returned", tc.optionMustBe, tc.hasOptionMustBe)
			}
		})
	}
}

func TestGetWords(t *testing.T) {
	cases := []struct {
		timeKind string
		num      int
		result   string
		lang     string
	}{
		// english
		{"days", 11, "11 days ago", "en"},
		{"days", 21, "21 days ago", "en"},
		{"seconds", 30, "30 seconds ago", "en"},
		{"seconds", 31, "31 seconds ago", "en"},
		{"hours", 10, "10 hours ago", "en"},
		{"years", 2, "2 years ago", "en"},
		// russian
		{"hours", 5, "5 часов назад", "ru"},
		{"days", 11, "11 дней назад", "ru"},
		{"years", 21, "21 год назад", "ru"},
		{"minutes", 59, "59 минут назад", "ru"},
	}

	for _, tc := range cases {
		t.Run(tc.result, func(test *testing.T) {
			Set("language", tc.lang)
			Set("location", "Europe/Kiev")

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
		// english
		{smallSubTime(-60 * time.Second), "1 minute ago", "en"},
		{smallSubTime(-1 * time.Minute), "1 minute ago", "en"},
		{smallSubTime(-2 * time.Minute), "2 minutes ago", "en"},
		{smallSubTime(-9 * time.Minute), "9 minutes ago", "en"},
		{smallSubTime(-10 * time.Minute), "10 minutes ago", "en"},
		{smallSubTime(-11 * time.Minute), "11 minutes ago", "en"},
		{smallSubTime(-20 * time.Minute), "20 minutes ago", "en"},
		{smallSubTime(-21 * time.Minute), "21 minutes ago", "en"},
		{smallSubTime(-22 * time.Minute), "22 minutes ago", "en"},
		{smallSubTime(-30 * time.Minute), "30 minutes ago", "en"},
		{smallSubTime(-31 * time.Minute), "31 minutes ago", "en"},
		{smallSubTime(-59 * time.Minute), "59 minutes ago", "en"},
		{smallSubTime(-60 * time.Minute), "1 hour ago", "en"},
		{smallSubTime(-1 * time.Hour), "1 hour ago", "en"},
		{smallSubTime(-2 * time.Hour), "2 hours ago", "en"},
		{smallSubTime(-9 * time.Hour), "9 hours ago", "en"},
		{smallSubTime(-10 * time.Hour), "10 hours ago", "en"},
		{smallSubTime(-11 * time.Hour), "11 hours ago", "en"},
		{smallSubTime(-20 * time.Hour), "20 hours ago", "en"},
		{smallSubTime(-21 * time.Hour), "21 hours ago", "en"},
		{smallSubTime(-23 * time.Hour), "23 hours ago", "en"},
		{smallSubTime(-24 * time.Hour), "1 day ago", "en"},
		{smallSubTime(-30 * time.Hour), "1 day ago", "en"},
		{smallSubTime((-24 * 2) * time.Hour), "2 days ago", "en"},
		{smallSubTime((-24 * 6) * time.Hour), "6 days ago", "en"},
		{smallSubTime((-24 * 7) * time.Hour), "1 week ago", "en"},
		{smallSubTime((-24 * 14) * time.Hour), "2 weeks ago", "en"},
		{smallSubTime((-24 * 21) * time.Hour), "3 weeks ago", "en"},
		{bigSubTime(0, 1, 1), "1 month ago", "en"},
		{bigSubTime(0, 2, 1), "2 months ago", "en"},
		{bigSubTime(0, 9, 1), "9 months ago", "en"},
		{bigSubTime(0, 11, 1), "11 months ago", "en"},
		{bigSubTime(0, 12, 1), "1 year ago", "en"},
		{bigSubTime(1, 0, 1), "1 year ago", "en"},
		{bigSubTime(2, 0, 1), "2 years ago", "en"},
		{bigSubTime(21, 0, 1), "21 years ago", "en"},
		{bigSubTime(31, 0, 1), "31 years ago", "en"},
		{bigSubTime(100, 0, 1), "100 years ago", "en"},
		// russian
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
			Set("language", tc.lang)
			Set("location", "Europe/Kiev")

			if res := Take(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestTake_with_online_flag(t *testing.T) {
	cases := []struct {
		date   string
		result string
		lang   string
	}{
		// english
		{smallSubTime(time.Second * 2), "Online", "en"},
		{smallSubTime(time.Second), "Online", "en"},
		{smallSubTime(-1 * time.Second), "Online", "en"},
		{smallSubTime(-2 * time.Second), "Online", "en"},
		{smallSubTime(-9 * time.Second), "Online", "en"},
		{smallSubTime(-10 * time.Second), "Online", "en"},
		{smallSubTime(-11 * time.Second), "Online", "en"},
		{smallSubTime(-20 * time.Second), "Online", "en"},
		{smallSubTime(-21 * time.Second), "Online", "en"},
		{smallSubTime(-22 * time.Second), "Online", "en"},
		{smallSubTime(-30 * time.Second), "Online", "en"},
		{smallSubTime(-31 * time.Second), "Online", "en"},
		{smallSubTime(-59 * time.Second), "Online", "en"},
		{smallSubTime(-60 * time.Second), "1 minute ago", "en"},
		{smallSubTime(-1 * time.Minute), "1 minute ago", "en"},
		{smallSubTime(-2 * time.Minute), "2 minutes ago", "en"},
		{smallSubTime(-9 * time.Minute), "9 minutes ago", "en"},
		{smallSubTime(-10 * time.Minute), "10 minutes ago", "en"},
		{smallSubTime(-11 * time.Minute), "11 minutes ago", "en"},
		{smallSubTime(-20 * time.Minute), "20 minutes ago", "en"},
		{smallSubTime(-21 * time.Minute), "21 minutes ago", "en"},
		{smallSubTime(-22 * time.Minute), "22 minutes ago", "en"},
		{smallSubTime(-30 * time.Minute), "30 minutes ago", "en"},
		{smallSubTime(-31 * time.Minute), "31 minutes ago", "en"},
		{smallSubTime(-59 * time.Minute), "59 minutes ago", "en"},
		{smallSubTime(-60 * time.Minute), "1 hour ago", "en"},
		{smallSubTime(-1 * time.Hour), "1 hour ago", "en"},
		{smallSubTime(-2 * time.Hour), "2 hours ago", "en"},
		{smallSubTime(-9 * time.Hour), "9 hours ago", "en"},
		{smallSubTime(-10 * time.Hour), "10 hours ago", "en"},
		{smallSubTime(-11 * time.Hour), "11 hours ago", "en"},
		{smallSubTime(-20 * time.Hour), "20 hours ago", "en"},
		{smallSubTime(-21 * time.Hour), "21 hours ago", "en"},
		{smallSubTime(-23 * time.Hour), "23 hours ago", "en"},
		{smallSubTime(-24 * time.Hour), "1 day ago", "en"},
		{smallSubTime(-30 * time.Hour), "1 day ago", "en"},
		{smallSubTime((-24 * 2) * time.Hour), "2 days ago", "en"},
		{smallSubTime((-24 * 6) * time.Hour), "6 days ago", "en"},
		{smallSubTime((-24 * 7) * time.Hour), "1 week ago", "en"},
		{smallSubTime((-24 * 14) * time.Hour), "2 weeks ago", "en"},
		{smallSubTime((-24 * 21) * time.Hour), "3 weeks ago", "en"},
		{bigSubTime(0, 1, 1), "1 month ago", "en"},
		{bigSubTime(0, 2, 1), "2 months ago", "en"},
		{bigSubTime(0, 9, 1), "9 months ago", "en"},
		{bigSubTime(0, 11, 1), "11 months ago", "en"},
		{bigSubTime(0, 12, 1), "1 year ago", "en"},
		{bigSubTime(1, 0, 1), "1 year ago", "en"},
		{bigSubTime(2, 0, 1), "2 years ago", "en"},
		{bigSubTime(21, 0, 1), "21 years ago", "en"},
		{bigSubTime(31, 0, 1), "31 years ago", "en"},
		{bigSubTime(100, 0, 1), "100 years ago", "en"},
		// russian
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
		{smallSubTime(-59 * time.Second), "В сети", "ru"},
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
		t.Run("result for "+tc.date+"|online", func(test *testing.T) {
			Set("language", tc.lang)

			if res := Take(tc.date + "|online"); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestTakeWithSeconds(t *testing.T) {
	cases := []struct {
		date   string
		result []string
		lang   string
	}{
		// english
		{smallSubTime(time.Second * 2), []string{"0 seconds ago", "1 second ago"}, "en"},
		{smallSubTime(time.Second), []string{"0 seconds ago", "1 second ago"}, "en"},
		{smallSubTime(-1 * time.Second), []string{"1 second ago", "2 seconds ago"}, "en"},
		{smallSubTime(-2 * time.Second), []string{"2 seconds ago", "3 seconds ago"}, "en"},
		{smallSubTime(-9 * time.Second), []string{"9 seconds ago", "10 seconds ago"}, "en"},
		{smallSubTime(-10 * time.Second), []string{"10 seconds ago", "11 seconds ago"}, "en"},
		{smallSubTime(-11 * time.Second), []string{"11 seconds ago", "12 seconds ago"}, "en"},
		{smallSubTime(-20 * time.Second), []string{"20 seconds ago", "21 seconds ago"}, "en"},
		{smallSubTime(-21 * time.Second), []string{"21 seconds ago", "22 seconds ago"}, "en"},
		{smallSubTime(-22 * time.Second), []string{"22 seconds ago", "23 seconds ago"}, "en"},
		{smallSubTime(-30 * time.Second), []string{"30 seconds ago", "31 seconds ago"}, "en"},
		{smallSubTime(-31 * time.Second), []string{"31 seconds ago", "32 seconds ago"}, "en"},
		{smallSubTime(-59 * time.Second), []string{"59 seconds ago", "1 minute ago"}, "en"},
		// russian
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
		{smallSubTime(-31 * time.Second), []string{"31 секунда назад", "32 секунда назад"}, "ru"},
		{smallSubTime(-59 * time.Second), []string{"59 секунд назад", "1 минута назад"}, "ru"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date, func(test *testing.T) {
			Set("language", tc.lang)
			Set("location", "Europe/Kiev")

			if res := Take(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
