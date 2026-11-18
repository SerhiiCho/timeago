package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseJa(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{utils.SubMinutes(1), "1分前"},
		{utils.SubMinutes(2), "2分前"},
		{utils.SubMinutes(5), "5分前"},
		{utils.SubMinutes(9), "9分前"},
		{utils.SubMinutes(10), "10分前"},
		{utils.SubMinutes(11), "11分前"},
		{utils.SubMinutes(20), "20分前"},
		{utils.SubMinutes(21), "21分前"},
		{utils.SubMinutes(22), "22分前"},
		{utils.SubMinutes(30), "30分前"},
		{utils.SubMinutes(31), "31分前"},
		{utils.SubMinutes(59), "59分前"},
		{utils.SubHours(1), "1時間前"},
		{utils.SubHours(2), "2時間前"},
		{utils.SubHours(9), "9時間前"},
		{utils.SubHours(10), "10時間前"},
		{utils.SubHours(11), "11時間前"},
		{utils.SubHours(20), "20時間前"},
		{utils.SubHours(21), "21時間前"},
		{utils.SubHours(23), "23時間前"},
		{utils.SubDays(1), "1日前"},
		{utils.SubDays(2), "2日前"},
		{utils.SubDays(4), "4日前"},
		{utils.SubDays(5), "5日前"},
		{utils.SubDays(6), "6日前"},
		{utils.SubWeeks(1), "1週間前"},
		{utils.SubWeeks(2), "2週間前"},
		{utils.SubWeeks(3), "3週間前"},
		{utils.SubMonths(1), "1ヶ月前"},
		{utils.SubMonths(2), "2ヶ月前"},
		{utils.SubMonths(9), "9ヶ月前"},
		{utils.SubMonths(11), "11ヶ月前"},
		{utils.SubYears(1), "1年前"},
		{utils.SubYears(2), "2年前"},
		{utils.SubYears(21), "21年前"},
		{utils.SubYears(31), "31年前"},
		{utils.SubYears(100), "100年前"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangJa})

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

func TestParseJaWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{utils.SubSeconds(0), []string{"0秒前", "1秒前"}},
		{utils.SubSeconds(1), []string{"1秒前", "2秒前"}},
		{utils.SubSeconds(2), []string{"2秒前", "3秒前"}},
		{utils.SubSeconds(9), []string{"9秒前", "10秒前"}},
		{utils.SubSeconds(10), []string{"10秒前", "11秒前"}},
		{utils.SubSeconds(11), []string{"11秒前", "12秒前"}},
		{utils.SubSeconds(20), []string{"20秒前", "21秒前"}},
		{utils.SubSeconds(21), []string{"21秒前", "22秒前"}},
		{utils.SubSeconds(22), []string{"22秒前", "23秒前"}},
		{utils.SubSeconds(30), []string{"30秒前", "31秒前"}},
		{utils.SubSeconds(59), []string{"59秒前", "1分前"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangJa})

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
