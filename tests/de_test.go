package tests

import (
	"testing"
	"time"

	. "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/config"
)

const langDe = "de"

func TestParseDe(t *testing.T) {
	cases := []struct {
		date   time.Time
		result string
	}{
		{subMinutes(1), "Vor 1 Minute"},
		{subMinutes(2), "Vor 2 Minuten"},
		{subMinutes(5), "Vor 5 Minuten"},
		{subMinutes(9), "Vor 9 Minuten"},
		{subMinutes(10), "Vor 10 Minuten"},
		{subMinutes(11), "Vor 11 Minuten"},
		{subMinutes(20), "Vor 20 Minuten"},
		{subMinutes(21), "Vor 21 Minuten"},
		{subMinutes(22), "Vor 22 Minuten"},
		{subMinutes(30), "Vor 30 Minuten"},
		{subMinutes(31), "Vor 31 Minuten"},
		{subMinutes(59), "Vor 59 Minuten"},
		{subMinutes(60), "Vor 1 Stunde"},
		{subHours(1), "Vor 1 Stunde"},
		{subHours(2), "Vor 2 Stunden"},
		{subHours(9), "Vor 9 Stunden"},
		{subHours(10), "Vor 10 Stunden"},
		{subHours(11), "Vor 11 Stunden"},
		{subHours(20), "Vor 20 Stunden"},
		{subHours(21), "Vor 21 Stunden"},
		{subHours(23), "Vor 23 Stunden"},
		{subDays(1), "Vor 1 Tag"},
		{subDays(2), "Vor 2 Tagen"},
		{subDays(6), "Vor 6 Tagen"},
		{subWeeks(1), "Vor 1 Woche"},
		{subWeeks(2), "Vor 2 Wochen"},
		{subWeeks(3), "Vor 3 Wochen"},
		{subMonths(1), "Vor 1 Monat"},
		{subMonths(2), "Vor 2 Monaten"},
		{subMonths(3), "Vor 3 Monaten"},
		{subMonths(4), "Vor 4 Monaten"},
		{subMonths(5), "Vor 5 Monaten"},
		{subMonths(6), "Vor 6 Monaten"},
		{subMonths(7), "Vor 7 Monaten"},
		{subMonths(8), "Vor 8 Monaten"},
		{subMonths(9), "Vor 9 Monaten"},
		{subMonths(10), "Vor 10 Monaten"},
		{subMonths(11), "Vor 11 Monaten"},
		{subYears(1), "Vor 1 Jahr"},
		{subYears(2), "Vor 2 Jahren"},
		{subYears(21), "Vor 21 Jahren"},
		{subYears(31), "Vor 31 Jahren"},
		{subYears(100), "Vor 100 Jahren"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			Configure(&config.Config{
				Language: langDe,
			})

			if res := Parse(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseDeWithSeconds(t *testing.T) {
	cases := []struct {
		date   time.Time
		result []string
	}{
		{subSeconds(0), []string{"Vor 0 Sekunden", "Vor 1 Sekunde"}},
		{subSeconds(1), []string{"Vor 1 Sekunde", "Vor 2 Sekunden"}},
		{subSeconds(2), []string{"Vor 2 Sekunden", "Vor 3 Sekunden"}},
		{subSeconds(9), []string{"Vor 9 Sekunden", "Vor 10 Sekunden"}},
		{subSeconds(10), []string{"Vor 10 Sekunden", "Vor 11 Sekunden"}},
		{subSeconds(11), []string{"Vor 11 Sekunden", "Vor 12 Sekunden"}},
		{subSeconds(20), []string{"Vor 20 Sekunden", "Vor 21 Sekunden"}},
		{subSeconds(21), []string{"Vor 21 Sekunden", "Vor 22 Sekunden"}},
		{subSeconds(22), []string{"Vor 22 Sekunden", "Vor 23 Sekunden"}},
		{subSeconds(30), []string{"Vor 30 Sekunden", "Vor 31 Sekunden"}},
		{subSeconds(31), []string{"Vor 31 Sekunden", "Vor 32 Sekunden"}},
		{subSeconds(59), []string{"Vor 59 Sekunden", "Vor 1 Minute"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			Configure(&config.Config{
				Language: langDe,
			})

			if res := Parse(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
