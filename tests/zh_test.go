package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
)

func TestParseZh(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{subMinutes(1), "1分钟前"},
		{subMinutes(2), "2分钟前"},
		{subMinutes(5), "5分钟前"},
		{subMinutes(9), "9分钟前"},
		{subMinutes(10), "10分钟前"},
		{subMinutes(11), "11分钟前"},
		{subMinutes(20), "20分钟前"},
		{subMinutes(21), "21分钟前"},
		{subMinutes(22), "22分钟前"},
		{subMinutes(30), "30分钟前"},
		{subMinutes(31), "31分钟前"},
		{subMinutes(59), "59分钟前"},
		{subHours(1), "1小时前"},
		{subHours(2), "2小时前"},
		{subHours(9), "9小时前"},
		{subHours(10), "10小时前"},
		{subHours(11), "11小时前"},
		{subHours(20), "20小时前"},
		{subHours(21), "21小时前"},
		{subHours(23), "23小时前"},
		{subDays(1), "1天前"},
		{subDays(2), "2天前"},
		{subDays(4), "4天前"},
		{subDays(5), "5天前"},
		{subDays(6), "6天前"},
		{subWeeks(1), "1周前"},
		{subWeeks(2), "2周前"},
		{subWeeks(3), "3周前"},
		{subMonths(1), "1个月前"},
		{subMonths(2), "2个月前"},
		{subMonths(9), "9个月前"},
		{subMonths(11), "11个月前"},
		{subYears(1), "1年前"},
		{subYears(2), "2年前"},
		{subYears(21), "21年前"},
		{subYears(31), "31年前"},
		{subYears(100), "100年前"},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangZh})

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

func TestParseZhWithSeconds(t *testing.T) {
	cases := []struct {
		date time.Time
		res  []string
	}{
		{subSeconds(0), []string{"0秒前", "1秒前"}},
		{subSeconds(1), []string{"1秒前", "2秒前"}},
		{subSeconds(2), []string{"2秒前", "3秒前"}},
		{subSeconds(9), []string{"9秒前", "10秒前"}},
		{subSeconds(10), []string{"10秒前", "11秒前"}},
		{subSeconds(11), []string{"11秒前", "12秒前"}},
		{subSeconds(20), []string{"20秒前", "21秒前"}},
		{subSeconds(21), []string{"21秒前", "22秒前"}},
		{subSeconds(22), []string{"22秒前", "23秒前"}},
		{subSeconds(30), []string{"30秒前", "31秒前"}},
		{subSeconds(59), []string{"59秒前", "1分钟前"}},
	}

	for _, tc := range cases {
		t.Run("result for "+tc.date.String(), func(test *testing.T) {
			ago.Reconfigure(ago.Config{Language: ago.LangZh})

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
