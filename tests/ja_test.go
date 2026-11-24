package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v4"
)

func TestParseJa(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{subMins(1), "1分前"},
		{subMins(2), "2分前"},
		{subMins(5), "5分前"},
		{subMins(9), "9分前"},
		{subMins(10), "10分前"},
		{subMins(11), "11分前"},
		{subMins(20), "20分前"},
		{subMins(21), "21分前"},
		{subMins(22), "22分前"},
		{subMins(30), "30分前"},
		{subMins(31), "31分前"},
		{subMins(59), "59分前"},
		{subHours(1), "1時間前"},
		{subHours(2), "2時間前"},
		{subHours(9), "9時間前"},
		{subHours(10), "10時間前"},
		{subHours(11), "11時間前"},
		{subHours(20), "20時間前"},
		{subHours(21), "21時間前"},
		{subHours(23), "23時間前"},
		{subDays(1), "1日前"},
		{subDays(2), "2日前"},
		{subDays(4), "4日前"},
		{subDays(5), "5日前"},
		{subDays(6), "6日前"},
		{subWeeks(1), "1週間前"},
		{subWeeks(2), "2週間前"},
		{subWeeks(3), "3週間前"},
		{subMonths(1), "1ヶ月前"},
		{subMonths(2), "2ヶ月前"},
		{subMonths(9), "9ヶ月前"},
		{subMonths(11), "11ヶ月前"},
		{subYears(1), "1年前"},
		{subYears(2), "2年前"},
		{subYears(21), "21年前"},
		{subYears(31), "31年前"},
		{subYears(100), "100年前"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangJa})

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

func TestParseJaWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{subSecs(0), []string{"0秒前", "1秒前"}},
		{subSecs(1), []string{"1秒前", "2秒前"}},
		{subSecs(2), []string{"2秒前", "3秒前"}},
		{subSecs(9), []string{"9秒前", "10秒前"}},
		{subSecs(10), []string{"10秒前", "11秒前"}},
		{subSecs(11), []string{"11秒前", "12秒前"}},
		{subSecs(20), []string{"20秒前", "21秒前"}},
		{subSecs(21), []string{"21秒前", "22秒前"}},
		{subSecs(22), []string{"22秒前", "23秒前"}},
		{subSecs(30), []string{"30秒前", "31秒前"}},
		{subSecs(59), []string{"59秒前", "1分前"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangJa})

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
