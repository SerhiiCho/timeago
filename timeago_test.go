package timeago

import (
	"testing"
	"time"

	"github.com/SerhiiCho/timeago/v3/internal/utils"
)

func TestParse(t *testing.T) {
	cases := []struct {
		date interface{}
		res  string
	}{
		{utils.UnixFromPastDate(time.Minute), "1 minute ago"},
		{utils.UnixFromPastDate(time.Minute * 5), "5 minutes ago"},
		{utils.UnixFromPastDate(time.Hour), "1 hour ago"},
		{utils.UnixFromPastDate(time.Hour * 3), "3 hours ago"},
		{utils.UnixFromPastDate(time.Hour * 5), "5 hours ago"},
		{utils.UnixFromPastDate(time.Hour * 24), "1 day ago"},
		{utils.UnixFromPastDate(time.Hour * 24 * 2), "2 days ago"},
		{utils.UnixFromPastDate(time.Hour * 24 * 3), "3 days ago"},
		{utils.UnixFromPastDate(time.Hour * 24 * 4), "4 days ago"},
		{utils.UnixFromPastDate(time.Hour * 24 * 5), "5 days ago"},
		{utils.UnixFromPastDate(time.Hour * 24 * 6), "6 days ago"},
		{utils.UnixFromPastDate(time.Hour * 24 * 7), "1 week ago"},
		{utils.SubMinutes(1), "1 minute ago"},
		{utils.SubMinutes(2), "2 minutes ago"},
		{utils.SubMinutes(3), "3 minutes ago"},
		{utils.SubMinutes(4), "4 minutes ago"},
		{utils.SubMinutes(5), "5 minutes ago"},
		{utils.SubMinutes(6), "6 minutes ago"},
		{utils.SubHours(7), "7 hours ago"},
		{utils.SubHours(8), "8 hours ago"},
		{utils.SubHours(9), "9 hours ago"},
		{utils.SubHours(10), "10 hours ago"},
		{utils.SubHours(11), "11 hours ago"},
		{utils.AddMinutes(2), "2 minutes"},
		{utils.AddMinutes(5), "5 minutes"},
		{utils.AddMinutes(10), "10 minutes"},
		{utils.AddHours(1), "1 hour"},
		{utils.AddHours(24), "1 day"},
		{utils.AddHours(48), "2 days"},
	}

	Reconfigure(Config{Language: LangEn})

	for _, tc := range cases {
		t.Run(tc.res, func(t *testing.T) {
			res, err := Parse(tc.date)

			if err != nil {
				t.Errorf("Error must be nil, but got %q instead", err)
			}

			if res != tc.res {
				t.Errorf("Result must be %q, but got %q instead", tc.res, res)
			}
		})
	}
}
