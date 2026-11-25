package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v4"
)

func TestParseUk(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{subSecs(60), "1 хвилина тому"},
		{subMins(1), "1 хвилина тому"},
		{subMins(2), "2 хвилини тому"},
		{subMins(5), "5 хвилин тому"},
		{subMins(9), "9 хвилин тому"},
		{subMins(10), "10 хвилин тому"},
		{subMins(11), "11 хвилин тому"},
		{subMins(20), "20 хвилин тому"},
		{subMins(21), "21 хвилина тому"},
		{subMins(22), "22 хвилини тому"},
		{subMins(30), "30 хвилин тому"},
		{subMins(31), "31 хвилина тому"},
		{subMins(59), "59 хвилин тому"},
		{subMins(60), "1 година тому"},
		{subHours(1), "1 година тому"},
		{subHours(2), "2 години тому"},
		{subHours(9), "9 годин тому"},
		{subHours(10), "10 годин тому"},
		{subHours(11), "11 годин тому"},
		{subHours(20), "20 годин тому"},
		{subHours(21), "21 година тому"},
		{subHours(23), "23 години тому"},
		{subHours(24), "1 день тому"},
		{subDays(1), "1 день тому"},
		{subDays(2), "2 дні тому"},
		{subDays(3), "3 дні тому"},
		{subDays(4), "4 дні тому"},
		{subDays(5), "5 днів тому"},
		{subDays(6), "6 днів тому"},
		{subDays(7), "1 тиждень тому"},
		{subWeeks(1), "1 тиждень тому"},
		{subWeeks(2), "2 тижні тому"},
		{subWeeks(3), "3 тижні тому"},
		{subMonths(1), "1 місяць тому"},
		{subMonths(2), "2 місяці тому"},
		{subMonths(3), "3 місяці тому"},
		{subMonths(4), "4 місяці тому"},
		{subMonths(5), "5 місяців тому"},
		{subMonths(6), "6 місяців тому"},
		{subMonths(7), "7 місяців тому"},
		{subMonths(8), "8 місяців тому"},
		{subMonths(9), "9 місяців тому"},
		{subMonths(10), "10 місяців тому"},
		{subMonths(11), "11 місяців тому"},
		{subMonths(12), "1 рік тому"},
		{subYears(1), "1 рік тому"},
		{subYears(2), "2 роки тому"},
		{subYears(5), "5 років тому"},
		{subYears(6), "6 років тому"},
		{subYears(7), "7 років тому"},
		{subYears(21), "21 рік тому"},
		{subYears(31), "31 рік тому"},
		{subYears(100), "100 років тому"},
		{subYears(111), "111 років тому"},
		{subYears(112), "112 років тому"},
		{subYears(113), "113 років тому"},
		{subYears(114), "114 років тому"},
		{subYears(115), "115 років тому"},
		{subYears(117), "117 років тому"},
		{subYears(118), "118 років тому"},
		{subYears(119), "119 років тому"},
		{subYears(120), "120 років тому"},
		{subYears(121), "121 рік тому"},
		{subYears(122), "122 роки тому"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangUk})

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

func TestParseUkWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{subSecs(0), []string{"0 секунд тому", "1 секунда тому"}},
		{subSecs(1), []string{"1 секунда тому", "2 секунди тому"}},
		{subSecs(2), []string{"2 секунди тому", "3 секунди тому"}},
		{subSecs(3), []string{"3 секунди тому", "4 секунди тому"}},
		{subSecs(4), []string{"4 секунди тому", "5 секунд тому"}},
		{subSecs(5), []string{"5 секунд тому", "6 секунд тому"}},
		{subSecs(6), []string{"6 секунд тому", "7 секунд тому"}},
		{subSecs(7), []string{"7 секунд тому", "8 секунд тому"}},
		{subSecs(8), []string{"8 секунд тому", "9 секунд тому"}},
		{subSecs(9), []string{"9 секунд тому", "10 секунд тому"}},
		{subSecs(10), []string{"10 секунд тому", "11 секунд тому"}},
		{subSecs(11), []string{"11 секунд тому", "12 секунд тому"}},
		{subSecs(14), []string{"14 секунд тому", "15 секунд тому"}},
		{subSecs(16), []string{"16 секунд тому", "17 секунд тому"}},
		{subSecs(21), []string{"21 секунда тому", "22 секунди тому"}},
		{subSecs(23), []string{"23 секунди тому", "24 секунди тому"}},
		{subSecs(25), []string{"25 секунд тому", "26 секунд тому"}},
		{subSecs(29), []string{"29 секунд тому", "30 секунд тому"}},
		{subSecs(30), []string{"30 секунд тому", "31 секунда тому"}},
		{subSecs(31), []string{"31 секунда тому", "32 секунди тому"}},
		{subSecs(32), []string{"32 секунди тому", "33 секунди тому"}},
		{subSecs(33), []string{"33 секунди тому", "34 секунди тому"}},
		{subSecs(34), []string{"34 секунди тому", "35 секунд тому"}},
		{subSecs(35), []string{"35 секунд тому", "36 секунд тому"}},
		{subSecs(36), []string{"36 секунд тому", "37 секунд тому"}},
		{subSecs(37), []string{"37 секунд тому", "38 секунд тому"}},
		{subSecs(39), []string{"39 секунд тому", "40 секунд тому"}},
		{subSecs(40), []string{"40 секунд тому", "41 секунда тому"}},
		{subSecs(41), []string{"41 секунда тому", "42 секунди тому"}},
		{subSecs(59), []string{"59 секунд тому", "1 хвилина тому"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangUk})

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
