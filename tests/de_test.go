package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseDe(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{utils.SubMinutes(1), "Vor 1 Minute"},
		{utils.SubMinutes(2), "Vor 2 Minuten"},
		{utils.SubMinutes(5), "Vor 5 Minuten"},
		{utils.SubMinutes(9), "Vor 9 Minuten"},
		{utils.SubMinutes(10), "Vor 10 Minuten"},
		{utils.SubMinutes(11), "Vor 11 Minuten"},
		{utils.SubMinutes(20), "Vor 20 Minuten"},
		{utils.SubMinutes(21), "Vor 21 Minuten"},
		{utils.SubMinutes(22), "Vor 22 Minuten"},
		{utils.SubMinutes(30), "Vor 30 Minuten"},
		{utils.SubMinutes(31), "Vor 31 Minuten"},
		{utils.SubMinutes(59), "Vor 59 Minuten"},
		{utils.SubMinutes(60), "Vor 1 Stunde"},
		{utils.SubHours(1), "Vor 1 Stunde"},
		{utils.SubHours(2), "Vor 2 Stunden"},
		{utils.SubHours(9), "Vor 9 Stunden"},
		{utils.SubHours(10), "Vor 10 Stunden"},
		{utils.SubHours(11), "Vor 11 Stunden"},
		{utils.SubHours(20), "Vor 20 Stunden"},
		{utils.SubHours(21), "Vor 21 Stunden"},
		{utils.SubHours(23), "Vor 23 Stunden"},
		{utils.SubDays(1), "Vor 1 Tag"},
		{utils.SubDays(2), "Vor 2 Tagen"},
		{utils.SubDays(6), "Vor 6 Tagen"},
		{utils.SubWeeks(1), "Vor 1 Woche"},
		{utils.SubWeeks(2), "Vor 2 Wochen"},
		{utils.SubWeeks(3), "Vor 3 Wochen"},
		{utils.SubMonths(1), "Vor 1 Monat"},
		{utils.SubMonths(2), "Vor 2 Monaten"},
		{utils.SubMonths(3), "Vor 3 Monaten"},
		{utils.SubMonths(4), "Vor 4 Monaten"},
		{utils.SubMonths(5), "Vor 5 Monaten"},
		{utils.SubMonths(6), "Vor 6 Monaten"},
		{utils.SubMonths(7), "Vor 7 Monaten"},
		{utils.SubMonths(8), "Vor 8 Monaten"},
		{utils.SubMonths(9), "Vor 9 Monaten"},
		{utils.SubMonths(10), "Vor 10 Monaten"},
		{utils.SubMonths(11), "Vor 11 Monaten"},
		{utils.SubYears(1), "Vor 1 Jahr"},
		{utils.SubYears(2), "Vor 2 Jahren"},
		{utils.SubYears(21), "Vor 21 Jahren"},
		{utils.SubYears(31), "Vor 31 Jahren"},
		{utils.SubYears(100), "Vor 100 Jahren"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangDe})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("Result must be %s, but got %s instead", tc.res, res)
			}
		})
	}
}

func TestParseDeWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{utils.SubSeconds(0), []string{"Vor 0 Sekunden", "Vor 1 Sekunde"}},
		{utils.SubSeconds(1), []string{"Vor 1 Sekunde", "Vor 2 Sekunden"}},
		{utils.SubSeconds(2), []string{"Vor 2 Sekunden", "Vor 3 Sekunden"}},
		{utils.SubSeconds(9), []string{"Vor 9 Sekunden", "Vor 10 Sekunden"}},
		{utils.SubSeconds(10), []string{"Vor 10 Sekunden", "Vor 11 Sekunden"}},
		{utils.SubSeconds(11), []string{"Vor 11 Sekunden", "Vor 12 Sekunden"}},
		{utils.SubSeconds(20), []string{"Vor 20 Sekunden", "Vor 21 Sekunden"}},
		{utils.SubSeconds(21), []string{"Vor 21 Sekunden", "Vor 22 Sekunden"}},
		{utils.SubSeconds(22), []string{"Vor 22 Sekunden", "Vor 23 Sekunden"}},
		{utils.SubSeconds(30), []string{"Vor 30 Sekunden", "Vor 31 Sekunden"}},
		{utils.SubSeconds(31), []string{"Vor 31 Sekunden", "Vor 32 Sekunden"}},
		{utils.SubSeconds(59), []string{"Vor 59 Sekunden", "Vor 1 Minute"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangDe})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res[0] && res != tc.res[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.res[0], tc.res[1], res)
			}
		})
	}
}
