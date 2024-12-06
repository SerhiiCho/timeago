package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseUk(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{utils.SubSeconds(60), "1 хвилина тому"},
		{utils.SubMinutes(1), "1 хвилина тому"},
		{utils.SubMinutes(2), "2 хвилини тому"},
		{utils.SubMinutes(5), "5 хвилин тому"},
		{utils.SubMinutes(9), "9 хвилин тому"},
		{utils.SubMinutes(10), "10 хвилин тому"},
		{utils.SubMinutes(11), "11 хвилин тому"},
		{utils.SubMinutes(20), "20 хвилин тому"},
		{utils.SubMinutes(21), "21 хвилина тому"},
		{utils.SubMinutes(22), "22 хвилини тому"},
		{utils.SubMinutes(30), "30 хвилин тому"},
		{utils.SubMinutes(31), "31 хвилина тому"},
		{utils.SubMinutes(59), "59 хвилин тому"},
		{utils.SubMinutes(60), "1 година тому"},
		{utils.SubHours(1), "1 година тому"},
		{utils.SubHours(2), "2 години тому"},
		{utils.SubHours(9), "9 годин тому"},
		{utils.SubHours(10), "10 годин тому"},
		{utils.SubHours(11), "11 годин тому"},
		{utils.SubHours(20), "20 годин тому"},
		{utils.SubHours(21), "21 година тому"},
		{utils.SubHours(23), "23 години тому"},
		{utils.SubHours(24), "1 день тому"},
		{utils.SubDays(1), "1 день тому"},
		{utils.SubDays(2), "2 дні тому"},
		{utils.SubDays(3), "3 дні тому"},
		{utils.SubDays(4), "4 дні тому"},
		{utils.SubDays(5), "5 днів тому"},
		{utils.SubDays(6), "6 днів тому"},
		{utils.SubDays(7), "1 тиждень тому"},
		{utils.SubWeeks(1), "1 тиждень тому"},
		{utils.SubWeeks(2), "2 тижні тому"},
		{utils.SubWeeks(3), "3 тижні тому"},
		{utils.SubMonths(1), "1 місяць тому"},
		{utils.SubMonths(2), "2 місяці тому"},
		{utils.SubMonths(3), "3 місяці тому"},
		{utils.SubMonths(4), "4 місяці тому"},
		{utils.SubMonths(5), "5 місяців тому"},
		{utils.SubMonths(6), "6 місяців тому"},
		{utils.SubMonths(7), "7 місяців тому"},
		{utils.SubMonths(8), "8 місяців тому"},
		{utils.SubMonths(9), "9 місяців тому"},
		{utils.SubMonths(10), "10 місяців тому"},
		{utils.SubMonths(11), "11 місяців тому"},
		{utils.SubMonths(12), "1 рік тому"},
		{utils.SubYears(1), "1 рік тому"},
		{utils.SubYears(2), "2 роки тому"},
		{utils.SubYears(5), "5 років тому"},
		{utils.SubYears(6), "6 років тому"},
		{utils.SubYears(7), "7 років тому"},
		{utils.SubYears(21), "21 рік тому"},
		{utils.SubYears(31), "31 рік тому"},
		{utils.SubYears(100), "100 років тому"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangUk})

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

func TestParseUkWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{utils.SubSeconds(0), []string{"0 секунд тому", "1 секунда тому"}},
		{utils.SubSeconds(1), []string{"1 секунда тому", "2 секунди тому"}},
		{utils.SubSeconds(2), []string{"2 секунди тому", "3 секунди тому"}},
		{utils.SubSeconds(3), []string{"3 секунди тому", "4 секунди тому"}},
		{utils.SubSeconds(4), []string{"4 секунди тому", "5 секунд тому"}},
		{utils.SubSeconds(5), []string{"5 секунд тому", "6 секунд тому"}},
		{utils.SubSeconds(6), []string{"6 секунд тому", "7 секунд тому"}},
		{utils.SubSeconds(7), []string{"7 секунд тому", "8 секунд тому"}},
		{utils.SubSeconds(8), []string{"8 секунд тому", "9 секунд тому"}},
		{utils.SubSeconds(9), []string{"9 секунд тому", "10 секунд тому"}},
		{utils.SubSeconds(10), []string{"10 секунд тому", "11 секунд тому"}},
		{utils.SubSeconds(11), []string{"11 секунд тому", "12 секунд тому"}},
		{utils.SubSeconds(14), []string{"14 секунд тому", "15 секунд тому"}},
		{utils.SubSeconds(16), []string{"16 секунд тому", "17 секунд тому"}},
		{utils.SubSeconds(21), []string{"21 секунда тому", "22 секунди тому"}},
		{utils.SubSeconds(23), []string{"23 секунди тому", "24 секунди тому"}},
		{utils.SubSeconds(25), []string{"25 секунд тому", "26 секунд тому"}},
		{utils.SubSeconds(29), []string{"29 секунд тому", "30 секунд тому"}},
		{utils.SubSeconds(30), []string{"30 секунд тому", "31 секунда тому"}},
		{utils.SubSeconds(31), []string{"31 секунда тому", "32 секунди тому"}},
		{utils.SubSeconds(32), []string{"32 секунди тому", "33 секунди тому"}},
		{utils.SubSeconds(33), []string{"33 секунди тому", "34 секунди тому"}},
		{utils.SubSeconds(34), []string{"34 секунди тому", "35 секунд тому"}},
		{utils.SubSeconds(35), []string{"35 секунд тому", "36 секунд тому"}},
		{utils.SubSeconds(36), []string{"36 секунд тому", "37 секунд тому"}},
		{utils.SubSeconds(37), []string{"37 секунд тому", "38 секунд тому"}},
		{utils.SubSeconds(39), []string{"39 секунд тому", "40 секунд тому"}},
		{utils.SubSeconds(40), []string{"40 секунд тому", "41 секунда тому"}},
		{utils.SubSeconds(41), []string{"41 секунда тому", "42 секунди тому"}},
		{utils.SubSeconds(59), []string{"59 секунд тому", "1 хвилина тому"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangUk})

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
