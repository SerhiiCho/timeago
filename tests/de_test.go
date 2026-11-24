package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
)

func TestParseDe(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{subMins(1), "Vor 1 Minute"},
		{subMins(2), "Vor 2 Minuten"},
		{subMins(5), "Vor 5 Minuten"},
		{subMins(9), "Vor 9 Minuten"},
		{subMins(10), "Vor 10 Minuten"},
		{subMins(11), "Vor 11 Minuten"},
		{subMins(20), "Vor 20 Minuten"},
		{subMins(21), "Vor 21 Minuten"},
		{subMins(22), "Vor 22 Minuten"},
		{subMins(30), "Vor 30 Minuten"},
		{subMins(31), "Vor 31 Minuten"},
		{subMins(59), "Vor 59 Minuten"},
		{subMins(60), "Vor 1 Stunde"},
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
			ago.Reconfigure(ago.Config{Language: ago.LangDe})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("want nil error, got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("want %q, got %q instead", tc.res, res)
			}
		})
	}
}

func TestParseDeWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{subSecs(0), []string{"Vor 0 Sekunden", "Vor 1 Sekunde"}},
		{subSecs(1), []string{"Vor 1 Sekunde", "Vor 2 Sekunden"}},
		{subSecs(2), []string{"Vor 2 Sekunden", "Vor 3 Sekunden"}},
		{subSecs(9), []string{"Vor 9 Sekunden", "Vor 10 Sekunden"}},
		{subSecs(10), []string{"Vor 10 Sekunden", "Vor 11 Sekunden"}},
		{subSecs(11), []string{"Vor 11 Sekunden", "Vor 12 Sekunden"}},
		{subSecs(20), []string{"Vor 20 Sekunden", "Vor 21 Sekunden"}},
		{subSecs(21), []string{"Vor 21 Sekunden", "Vor 22 Sekunden"}},
		{subSecs(22), []string{"Vor 22 Sekunden", "Vor 23 Sekunden"}},
		{subSecs(30), []string{"Vor 30 Sekunden", "Vor 31 Sekunden"}},
		{subSecs(31), []string{"Vor 31 Sekunden", "Vor 32 Sekunden"}},
		{subSecs(59), []string{"Vor 59 Sekunden", "Vor 1 Minute"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangDe})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("want nil error, got %v instead", err)
			}

			if res != tc.res[0] && res != tc.res[1] {
				test.Errorf(
					"want %q or %q, got %q instead",
					tc.res[0],
					tc.res[1],
					res,
				)
			}
		})
	}
}
