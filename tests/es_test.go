package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
)

func TestParseEs(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{subMinutes(1), "hace 1 minuto"},
		{subMinutes(2), "hace 2 minutos"},
		{subMinutes(5), "hace 5 minutos"},
		{subMinutes(9), "hace 9 minutos"},
		{subMinutes(10), "hace 10 minutos"},
		{subMinutes(11), "hace 11 minutos"},
		{subMinutes(20), "hace 20 minutos"},
		{subMinutes(21), "hace 21 minutos"},
		{subMinutes(22), "hace 22 minutos"},
		{subMinutes(30), "hace 30 minutos"},
		{subMinutes(31), "hace 31 minutos"},
		{subMinutes(59), "hace 59 minutos"},
		{subHours(1), "hace 1 hora"},
		{subHours(2), "hace 2 horas"},
		{subHours(9), "hace 9 horas"},
		{subHours(10), "hace 10 horas"},
		{subHours(11), "hace 11 horas"},
		{subHours(20), "hace 20 horas"},
		{subHours(21), "hace 21 horas"},
		{subHours(23), "hace 23 horas"},
		{subDays(1), "hace 1 día"},
		{subDays(2), "hace 2 días"},
		{subDays(4), "hace 4 días"},
		{subDays(5), "hace 5 días"},
		{subDays(6), "hace 6 días"},
		{subWeeks(1), "hace 1 semana"},
		{subWeeks(2), "hace 2 semanas"},
		{subWeeks(3), "hace 3 semanas"},
		{subMonths(1), "hace 1 mes"},
		{subMonths(2), "hace 2 meses"},
		{subMonths(9), "hace 9 meses"},
		{subMonths(11), "hace 11 meses"},
		{subYears(1), "hace 1 año"},
		{subYears(2), "hace 2 años"},
		{subYears(21), "hace 21 años"},
		{subYears(31), "hace 31 años"},
		{subYears(100), "hace 100 años"},
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
		{subSeconds(0), []string{"hace 0 segundos", "hace 1 segundo"}},
		{subSeconds(1), []string{"hace 1 segundo", "hace 2 segundos"}},
		{subSeconds(2), []string{"hace 2 segundos", "hace 3 segundos"}},
		{subSeconds(9), []string{"hace 9 segundos", "hace 10 segundos"}},
		{subSeconds(10), []string{"hace 10 segundos", "hace 11 segundos"}},
		{subSeconds(11), []string{"hace 11 segundos", "hace 12 segundos"}},
		{subSeconds(20), []string{"hace 20 segundos", "hace 21 segundos"}},
		{subSeconds(21), []string{"hace 21 segundos", "hace 22 segundos"}},
		{subSeconds(22), []string{"hace 22 segundos", "hace 23 segundos"}},
		{subSeconds(30), []string{"hace 30 segundos", "hace 31 segundos"}},
		{subSeconds(59), []string{"hace 59 segundos", "hace 1 minuto"}},
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
