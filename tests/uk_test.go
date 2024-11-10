package tests

import (
	"testing"
	"time"

	. "github.com/SerhiiCho/timeago/v3"
)

const langUk = "uk"

func TestParseUk(t *testing.T) {
	cases := []struct {
		date   time.Time
		result string
	}{
		{subSeconds(60), "1 хвилина тому"},
		{subMinutes(1), "1 хвилина тому"},
		{subMinutes(2), "2 хвилини тому"},
		{subMinutes(5), "5 хвилин тому"},
		{subMinutes(9), "9 хвилин тому"},
		{subMinutes(10), "10 хвилин тому"},
		{subMinutes(11), "11 хвилин тому"},
		{subMinutes(20), "20 хвилин тому"},
		{subMinutes(21), "21 хвилина тому"},
		{subMinutes(22), "22 хвилини тому"},
		{subMinutes(30), "30 хвилин тому"},
		{subMinutes(31), "31 хвилина тому"},
		{subMinutes(59), "59 хвилин тому"},
		{subMinutes(60), "1 година тому"},
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
		{subYears(2), "2 року тому"},
		{subYears(5), "5 років тому"},
		{subYears(6), "6 років тому"},
		{subYears(7), "7 років тому"},
		{subYears(21), "21 рік тому"},
		{subYears(31), "31 рік тому"},
		{subYears(100), "100 років тому"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			SetConfig(Config{
				Language: langUk,
			})

			if res := Parse(tc.date); res != tc.result {
				test.Errorf("Result must be %s, but got %s instead", tc.result, res)
			}
		})
	}
}

func TestParseUkWithSeconds(t *testing.T) {
	cases := []struct {
		date   time.Time
		result []string
	}{
		{subSeconds(0), []string{"0 секунд тому", "1 секунда тому"}},
		{subSeconds(1), []string{"1 секунда тому", "2 секунди тому"}},
		{subSeconds(2), []string{"2 секунди тому", "3 секунди тому"}},
		{subSeconds(3), []string{"3 секунди тому", "4 секунди тому"}},
		{subSeconds(4), []string{"4 секунди тому", "5 секунд тому"}},
		{subSeconds(9), []string{"9 секунд тому", "10 секунд тому"}},
		{subSeconds(10), []string{"10 секунд тому", "11 секунд тому"}},
		{subSeconds(11), []string{"11 секунд тому", "12 секунд тому"}},
		{subSeconds(29), []string{"29 секунд тому", "30 секунд тому"}},
		{subSeconds(30), []string{"30 секунд тому", "31 секунда тому"}},
		{subSeconds(31), []string{"31 секунда тому", "32 секунди тому"}},
		{subSeconds(59), []string{"59 секунд тому", "1 хвилина тому"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			SetConfig(Config{
				Language: langUk,
			})

			if res := Parse(tc.date); res != tc.result[0] && res != tc.result[1] {
				test.Errorf("Result must be %s or %s, but got %s instead", tc.result[0], tc.result[1], res)
			}
		})
	}
}
