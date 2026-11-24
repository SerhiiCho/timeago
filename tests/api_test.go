package tests

import (
	"strconv"
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
)

func TestParse(t *testing.T) {
	cases := []struct {
		date any
		res  string
	}{
		// Integer timestamp input parsing
		{unixFromPastDate(time.Minute * 5), "5 minutes ago"},
		{unixFromPastDate(time.Hour), "1 hour ago"},
		{unixFromPastDate(time.Hour * 3), "3 hours ago"},
		{unixFromPastDate(time.Hour * 5), "5 hours ago"},
		{unixFromPastDate(time.Hour * 24), "1 day ago"},
		{unixFromPastDate(time.Hour * 24 * 2), "2 days ago"},
		{unixFromPastDate(time.Hour * 24 * 3), "3 days ago"},
		{unixFromPastDate(time.Hour * 24 * 4), "4 days ago"},
		{unixFromPastDate(time.Hour * 24 * 5), "5 days ago"},
		{unixFromPastDate(time.Hour * 24 * 6), "6 days ago"},
		{unixFromPastDate(time.Hour * 24 * 7), "1 week ago"},
		// Integer 64 timestamp input parsing
		{int64(unixFromPastDate(time.Minute * 5)), "5 minutes ago"},
		{int64(unixFromPastDate(time.Hour)), "1 hour ago"},
		{int64(unixFromPastDate(time.Hour * 3)), "3 hours ago"},
		{int64(unixFromPastDate(time.Hour * 5)), "5 hours ago"},
		{int64(unixFromPastDate(time.Hour * 24)), "1 day ago"},
		{int64(unixFromPastDate(time.Hour * 24 * 2)), "2 days ago"},
		{int64(unixFromPastDate(time.Hour * 24 * 3)), "3 days ago"},
		{int64(unixFromPastDate(time.Hour * 24 * 4)), "4 days ago"},
		{int64(unixFromPastDate(time.Hour * 24 * 5)), "5 days ago"},
		{int64(unixFromPastDate(time.Hour * 24 * 6)), "6 days ago"},
		{int64(unixFromPastDate(time.Hour * 24 * 7)), "1 week ago"},
		// Unsigned Integer timestamp input parsing
		{uint(unixFromPastDate(time.Minute * 5)), "5 minutes ago"},
		{uint(unixFromPastDate(time.Hour)), "1 hour ago"},
		{uint(unixFromPastDate(time.Hour * 3)), "3 hours ago"},
		{uint(unixFromPastDate(time.Hour * 5)), "5 hours ago"},
		{uint(unixFromPastDate(time.Hour * 24)), "1 day ago"},
		{uint(unixFromPastDate(time.Hour * 24 * 2)), "2 days ago"},
		{uint(unixFromPastDate(time.Hour * 24 * 3)), "3 days ago"},
		{uint(unixFromPastDate(time.Hour * 24 * 4)), "4 days ago"},
		{uint(unixFromPastDate(time.Hour * 24 * 5)), "5 days ago"},
		{uint(unixFromPastDate(time.Hour * 24 * 6)), "6 days ago"},
		{uint(unixFromPastDate(time.Hour * 24 * 7)), "1 week ago"},
		// Unsigned Integer 32 timestamp input parsing
		{uint32(unixFromPastDate(time.Minute * 5)), "5 minutes ago"},
		{uint32(unixFromPastDate(time.Hour)), "1 hour ago"},
		{uint32(unixFromPastDate(time.Hour * 3)), "3 hours ago"},
		{uint32(unixFromPastDate(time.Hour * 5)), "5 hours ago"},
		{uint32(unixFromPastDate(time.Hour * 24)), "1 day ago"},
		{uint32(unixFromPastDate(time.Hour * 24 * 2)), "2 days ago"},
		{uint32(unixFromPastDate(time.Hour * 24 * 3)), "3 days ago"},
		{uint32(unixFromPastDate(time.Hour * 24 * 4)), "4 days ago"},
		{uint32(unixFromPastDate(time.Hour * 24 * 5)), "5 days ago"},
		{uint32(unixFromPastDate(time.Hour * 24 * 6)), "6 days ago"},
		{uint32(unixFromPastDate(time.Hour * 24 * 7)), "1 week ago"},
		// Unsigned Integer 64 timestamp input parsing
		{uint64(unixFromPastDate(time.Minute * 5)), "5 minutes ago"},
		{uint64(unixFromPastDate(time.Hour)), "1 hour ago"},
		{uint64(unixFromPastDate(time.Hour * 3)), "3 hours ago"},
		{uint64(unixFromPastDate(time.Hour * 5)), "5 hours ago"},
		{uint64(unixFromPastDate(time.Hour * 24)), "1 day ago"},
		{uint64(unixFromPastDate(time.Hour * 24 * 2)), "2 days ago"},
		{uint64(unixFromPastDate(time.Hour * 24 * 3)), "3 days ago"},
		{uint64(unixFromPastDate(time.Hour * 24 * 4)), "4 days ago"},
		{uint64(unixFromPastDate(time.Hour * 24 * 5)), "5 days ago"},
		{uint64(unixFromPastDate(time.Hour * 24 * 6)), "6 days ago"},
		{uint64(unixFromPastDate(time.Hour * 24 * 7)), "1 week ago"},
		// Integer 32 timestamp input parsing
		{int32(unixFromPastDate(time.Minute * 5)), "5 minutes ago"},
		{int32(unixFromPastDate(time.Hour)), "1 hour ago"},
		{int32(unixFromPastDate(time.Hour * 3)), "3 hours ago"},
		{int32(unixFromPastDate(time.Hour * 5)), "5 hours ago"},
		{int32(unixFromPastDate(time.Hour * 24)), "1 day ago"},
		{int32(unixFromPastDate(time.Hour * 24 * 2)), "2 days ago"},
		{int32(unixFromPastDate(time.Hour * 24 * 3)), "3 days ago"},
		{int32(unixFromPastDate(time.Hour * 24 * 4)), "4 days ago"},
		{int32(unixFromPastDate(time.Hour * 24 * 5)), "5 days ago"},
		{int32(unixFromPastDate(time.Hour * 24 * 6)), "6 days ago"},
		{int32(unixFromPastDate(time.Hour * 24 * 7)), "1 week ago"},
		// Negative integer timestamp input parsing
		{unixFromFutureDate(time.Minute * 5), "5 minutes"},
		{unixFromFutureDate(time.Hour), "1 hour"},
		{unixFromFutureDate(time.Hour * 3), "3 hours"},
		{unixFromFutureDate(time.Hour * 5), "5 hours"},
		{unixFromFutureDate(time.Hour * 24), "1 day"},
		{unixFromFutureDate(time.Hour * 24 * 2), "2 days"},
		{unixFromFutureDate(time.Hour * 24 * 3), "3 days"},
		{unixFromFutureDate(time.Hour * 24 * 4), "4 days"},
		{unixFromFutureDate(time.Hour * 24 * 5), "5 days"},
		{unixFromFutureDate(time.Hour * 24 * 6), "6 days"},
		{unixFromFutureDate(time.Hour * 24 * 7), "1 week"},
		// time.Time input parsing
		{subMins(2), "2 minutes ago"},
		{subMins(3), "3 minutes ago"},
		{subMins(4), "4 minutes ago"},
		{subMins(5), "5 minutes ago"},
		{subMins(6), "6 minutes ago"},
		{subHours(7), "7 hours ago"},
		{subHours(8), "8 hours ago"},
		{subHours(9), "9 hours ago"},
		{subHours(10), "10 hours ago"},
		{subHours(11), "11 hours ago"},
		// time.Time future date parsing
		{addMins(2), "2 minutes"},
		{addMins(5), "5 minutes"},
		{addMins(10), "10 minutes"},
		{addHours(1), "1 hour"},
		{addHours(24), "1 day"},
		{addHours(48), "2 days"},
		// Timestamp string input parsing
		{strconv.Itoa(unixFromPastDate(time.Minute * 5)), "5 minutes ago"},
		{strconv.Itoa(unixFromPastDate(time.Hour)), "1 hour ago"},
		{strconv.Itoa(unixFromPastDate(time.Hour * 3)), "3 hours ago"},
		{strconv.Itoa(unixFromPastDate(time.Hour * 5)), "5 hours ago"},
		{strconv.Itoa(unixFromPastDate(time.Hour * 24)), "1 day ago"},
		{strconv.Itoa(unixFromPastDate(time.Hour * 24 * 2)), "2 days ago"},
		{strconv.Itoa(unixFromPastDate(time.Hour * 24 * 3)), "3 days ago"},
		{strconv.Itoa(unixFromPastDate(time.Hour * 24 * 4)), "4 days ago"},
		{strconv.Itoa(unixFromPastDate(time.Hour * 24 * 5)), "5 days ago"},
		{strconv.Itoa(unixFromPastDate(time.Hour * 24 * 6)), "6 days ago"},
		{strconv.Itoa(unixFromPastDate(time.Hour * 24 * 7)), "1 week ago"},
		// Edge cases that prove flooring is working
		{subMonthsAndDays(1, 27), "1 month ago"},
		{subMonthsAndDays(6, 27), "6 months ago"},
		{subMonthsAndDays(10, 27), "10 months ago"},
		{subMonthsAndDays(11, 1), "11 months ago"},
		{subMonthsAndDays(11, 27), "11 months ago"},
		{subHoursAndMinutes(23, 59), "23 hours ago"},
		{subHoursAndMinutes(23, 60), "1 day ago"},
	}

	ago.Reconfigure(ago.Config{Language: ago.LangEn})

	for _, tc := range cases {
		t.Run(tc.res, func(t *testing.T) {
			res, err := ago.Parse(tc.date)
			if err != nil {
				t.Errorf("want nil error, got %v instead", err)
			}

			if res != tc.res {
				t.Errorf("want %q, got %q instead", tc.res, res)
			}
		})
	}
}
