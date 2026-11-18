package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseEs(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{utils.SubMinutes(1), "hace 1 minuto"},
		{utils.SubMinutes(2), "hace 2 minutos"},
		{utils.SubMinutes(5), "hace 5 minutos"},
		{utils.SubMinutes(9), "hace 9 minutos"},
		{utils.SubMinutes(10), "hace 10 minutos"},
		{utils.SubMinutes(11), "hace 11 minutos"},
		{utils.SubMinutes(20), "hace 20 minutos"},
		{utils.SubMinutes(21), "hace 21 minutos"},
		{utils.SubMinutes(22), "hace 22 minutos"},
		{utils.SubMinutes(30), "hace 30 minutos"},
		{utils.SubMinutes(31), "hace 31 minutos"},
		{utils.SubMinutes(59), "hace 59 minutos"},
		{utils.SubHours(1), "hace 1 hora"},
		{utils.SubHours(2), "hace 2 horas"},
		{utils.SubHours(9), "hace 9 horas"},
		{utils.SubHours(10), "hace 10 horas"},
		{utils.SubHours(11), "hace 11 horas"},
		{utils.SubHours(20), "hace 20 horas"},
		{utils.SubHours(21), "hace 21 horas"},
		{utils.SubHours(23), "hace 23 horas"},
		{utils.SubDays(1), "hace 1 día"},
		{utils.SubDays(2), "hace 2 días"},
		{utils.SubDays(4), "hace 4 días"},
		{utils.SubDays(5), "hace 5 días"},
		{utils.SubDays(6), "hace 6 días"},
		{utils.SubWeeks(1), "hace 1 semana"},
		{utils.SubWeeks(2), "hace 2 semanas"},
		{utils.SubWeeks(3), "hace 3 semanas"},
		{utils.SubMonths(1), "hace 1 mes"},
		{utils.SubMonths(2), "hace 2 meses"},
		{utils.SubMonths(9), "hace 9 meses"},
		{utils.SubMonths(11), "hace 11 meses"},
		{utils.SubYears(1), "hace 1 año"},
		{utils.SubYears(2), "hace 2 años"},
		{utils.SubYears(21), "hace 21 años"},
		{utils.SubYears(31), "hace 31 años"},
		{utils.SubYears(100), "hace 100 años"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangEs})

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

func TestParseEsWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{utils.SubSeconds(0), []string{"hace 0 segundos", "hace 1 segundo"}},
		{utils.SubSeconds(1), []string{"hace 1 segundo", "hace 2 segundos"}},
		{utils.SubSeconds(2), []string{"hace 2 segundos", "hace 3 segundos"}},
		{utils.SubSeconds(9), []string{"hace 9 segundos", "hace 10 segundos"}},
		{utils.SubSeconds(10), []string{"hace 10 segundos", "hace 11 segundos"}},
		{utils.SubSeconds(11), []string{"hace 11 segundos", "hace 12 segundos"}},
		{utils.SubSeconds(20), []string{"hace 20 segundos", "hace 21 segundos"}},
		{utils.SubSeconds(21), []string{"hace 21 segundos", "hace 22 segundos"}},
		{utils.SubSeconds(22), []string{"hace 22 segundos", "hace 23 segundos"}},
		{utils.SubSeconds(30), []string{"hace 30 segundos", "hace 31 segundos"}},
		{utils.SubSeconds(59), []string{"hace 59 segundos", "hace 1 minuto"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangEs})

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
