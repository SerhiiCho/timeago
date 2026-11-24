package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
)

func TestParseFr(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{subMins(1), "il y a 1 minute"},
		{subMins(2), "il y a 2 minutes"},
		{subMins(5), "il y a 5 minutes"},
		{subMins(9), "il y a 9 minutes"},
		{subMins(10), "il y a 10 minutes"},
		{subMins(11), "il y a 11 minutes"},
		{subMins(20), "il y a 20 minutes"},
		{subMins(21), "il y a 21 minutes"},
		{subMins(22), "il y a 22 minutes"},
		{subMins(30), "il y a 30 minutes"},
		{subMins(31), "il y a 31 minutes"},
		{subMins(59), "il y a 59 minutes"},
		{subHours(1), "il y a 1 heure"},
		{subHours(2), "il y a 2 heures"},
		{subHours(9), "il y a 9 heures"},
		{subHours(10), "il y a 10 heures"},
		{subHours(11), "il y a 11 heures"},
		{subHours(20), "il y a 20 heures"},
		{subHours(21), "il y a 21 heures"},
		{subHours(23), "il y a 23 heures"},
		{subDays(1), "il y a 1 jour"},
		{subDays(2), "il y a 2 jours"},
		{subDays(4), "il y a 4 jours"},
		{subDays(5), "il y a 5 jours"},
		{subDays(6), "il y a 6 jours"},
		{subWeeks(1), "il y a 1 semaine"},
		{subWeeks(2), "il y a 2 semaines"},
		{subWeeks(3), "il y a 3 semaines"},
		{subMonths(1), "il y a 1 mois"},
		{subMonths(2), "il y a 2 mois"},
		{subMonths(9), "il y a 9 mois"},
		{subMonths(11), "il y a 11 mois"},
		{subYears(1), "il y a 1 an"},
		{subYears(2), "il y a 2 ans"},
		{subYears(21), "il y a 21 ans"},
		{subYears(31), "il y a 31 ans"},
		{subYears(100), "il y a 100 ans"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangFr})

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

func TestParseFrWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{subSecs(0), []string{"il y a 0 secondes", "il y a 1 seconde"}},
		{subSecs(1), []string{"il y a 1 seconde", "il y a 2 secondes"}},
		{subSecs(2), []string{"il y a 2 secondes", "il y a 3 secondes"}},
		{subSecs(9), []string{"il y a 9 secondes", "il y a 10 secondes"}},
		{subSecs(10), []string{"il y a 10 secondes", "il y a 11 secondes"}},
		{subSecs(11), []string{"il y a 11 secondes", "il y a 12 secondes"}},
		{subSecs(20), []string{"il y a 20 secondes", "il y a 21 secondes"}},
		{subSecs(21), []string{"il y a 21 secondes", "il y a 22 secondes"}},
		{subSecs(22), []string{"il y a 22 secondes", "il y a 23 secondes"}},
		{subSecs(30), []string{"il y a 30 secondes", "il y a 31 secondes"}},
		{subSecs(59), []string{"il y a 59 secondes", "il y a 1 minute"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangFr})

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
