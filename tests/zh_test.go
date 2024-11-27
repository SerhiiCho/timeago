package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParseZh(t *testing.T) {
	cases := []struct {
		date time.Time
		res  string
	}{
		{utils.SubMinutes(1), "1分钟前"},
		{utils.SubMinutes(2), "2分钟前"},
		{utils.SubMinutes(5), "5分钟前"},
		{utils.SubMinutes(9), "9分钟前"},
		{utils.SubMinutes(10), "10分钟前"},
		{utils.SubMinutes(11), "11分钟前"},
		{utils.SubMinutes(20), "20分钟前"},
		{utils.SubMinutes(21), "21分钟前"},
		{utils.SubMinutes(22), "22分钟前"},
		{utils.SubMinutes(30), "30分钟前"},
		{utils.SubMinutes(31), "31分钟前"},
		{utils.SubMinutes(59), "59分钟前"},
		{utils.SubHours(1), "1小时前"},
		{utils.SubHours(2), "2小时前"},
		{utils.SubHours(9), "9小时前"},
		{utils.SubHours(10), "10小时前"},
		{utils.SubHours(11), "11小时前"},
		{utils.SubHours(20), "20小时前"},
		{utils.SubHours(21), "21小时前"},
		{utils.SubHours(23), "23小时前"},
		{utils.SubDays(1), "1天前"},
		{utils.SubDays(2), "2天前"},
		{utils.SubDays(4), "4天前"},
		{utils.SubDays(5), "5天前"},
		{utils.SubDays(6), "6天前"},
		{utils.SubWeeks(1), "1周前"},
		{utils.SubWeeks(2), "2周前"},
		{utils.SubWeeks(3), "3周前"},
		{utils.SubMonths(1), "1个月前"},
		{utils.SubMonths(2), "2个月前"},
		{utils.SubMonths(9), "9个月前"},
		{utils.SubMonths(11), "11个月前"},
		{utils.SubYears(1), "1年前"},
		{utils.SubYears(2), "2年前"},
		{utils.SubYears(21), "21年前"},
		{utils.SubYears(31), "31年前"},
		{utils.SubYears(100), "100年前"},
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
		{utils.SubSeconds(59), []string{"59秒前", "1分钟前"}},
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
