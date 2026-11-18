package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseFr(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{utils.SubMinutes(1), "il y a 1 minute"},
		{utils.SubMinutes(2), "il y a 2 minutes"},
		{utils.SubMinutes(5), "il y a 5 minutes"},
		{utils.SubMinutes(9), "il y a 9 minutes"},
		{utils.SubMinutes(10), "il y a 10 minutes"},
		{utils.SubMinutes(11), "il y a 11 minutes"},
		{utils.SubMinutes(20), "il y a 20 minutes"},
		{utils.SubMinutes(21), "il y a 21 minutes"},
		{utils.SubMinutes(22), "il y a 22 minutes"},
		{utils.SubMinutes(30), "il y a 30 minutes"},
		{utils.SubMinutes(31), "il y a 31 minutes"},
		{utils.SubMinutes(59), "il y a 59 minutes"},
		{utils.SubHours(1), "il y a 1 heure"},
		{utils.SubHours(2), "il y a 2 heures"},
		{utils.SubHours(9), "il y a 9 heures"},
		{utils.SubHours(10), "il y a 10 heures"},
		{utils.SubHours(11), "il y a 11 heures"},
		{utils.SubHours(20), "il y a 20 heures"},
		{utils.SubHours(21), "il y a 21 heures"},
		{utils.SubHours(23), "il y a 23 heures"},
		{utils.SubDays(1), "il y a 1 jour"},
		{utils.SubDays(2), "il y a 2 jours"},
		{utils.SubDays(4), "il y a 4 jours"},
		{utils.SubDays(5), "il y a 5 jours"},
		{utils.SubDays(6), "il y a 6 jours"},
		{utils.SubWeeks(1), "il y a 1 semaine"},
		{utils.SubWeeks(2), "il y a 2 semaines"},
		{utils.SubWeeks(3), "il y a 3 semaines"},
		{utils.SubMonths(1), "il y a 1 mois"},
		{utils.SubMonths(2), "il y a 2 mois"},
		{utils.SubMonths(9), "il y a 9 mois"},
		{utils.SubMonths(11), "il y a 11 mois"},
		{utils.SubYears(1), "il y a 1 an"},
		{utils.SubYears(2), "il y a 2 ans"},
		{utils.SubYears(21), "il y a 21 ans"},
		{utils.SubYears(31), "il y a 31 ans"},
		{utils.SubYears(100), "il y a 100 ans"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangFr})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res {
				test.Errorf("Result must be %q, but got %q instead", tc.res, res)
			}
		})
	}
}

func TestParseFrWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{utils.SubSeconds(0), []string{"il y a 0 secondes", "il y a 1 seconde"}},
		{utils.SubSeconds(1), []string{"il y a 1 seconde", "il y a 2 secondes"}},
		{utils.SubSeconds(2), []string{"il y a 2 secondes", "il y a 3 secondes"}},
		{utils.SubSeconds(9), []string{"il y a 9 secondes", "il y a 10 secondes"}},
		{utils.SubSeconds(10), []string{"il y a 10 secondes", "il y a 11 secondes"}},
		{utils.SubSeconds(11), []string{"il y a 11 secondes", "il y a 12 secondes"}},
		{utils.SubSeconds(20), []string{"il y a 20 secondes", "il y a 21 secondes"}},
		{utils.SubSeconds(21), []string{"il y a 21 secondes", "il y a 22 secondes"}},
		{utils.SubSeconds(22), []string{"il y a 22 secondes", "il y a 23 secondes"}},
		{utils.SubSeconds(30), []string{"il y a 30 secondes", "il y a 31 secondes"}},
		{utils.SubSeconds(59), []string{"il y a 59 secondes", "il y a 1 minute"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangFr})

			res, err := ago.Parse(tc.date)

			if err != nil {
				test.Errorf("Error must be nil, but got %v instead", err)
			}

			if res != tc.res[0] && res != tc.res[1] {
				test.Errorf("Result must be %q or %q, but got %q instead", tc.res[0], tc.res[1], res)
			}
		})
	}
}
