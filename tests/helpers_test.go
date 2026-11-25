package tests

import (
	"testing"
	"time"
)

const (
	second time.Duration = time.Second
	minute time.Duration = time.Minute
	hour   time.Duration = time.Hour
	day    time.Duration = hour * 24
	week   time.Duration = day * 7
)

func unixFromPastDate(subDuration time.Duration) int {
	return int(subTime(subDuration).UnixNano() / 1000000000)
}

func unixFromFutureDate(duration time.Duration) int {
	return int(addTime(duration).Add(time.Second - time.Nanosecond).Unix())
}

func subYears(years int) time.Time {
	return wantAddMonths(time.Now().Truncate(time.Second), -years*12)
}

func subTime(duration time.Duration) time.Time {
	return time.Now().Add(-duration).Truncate(time.Second)
}

func addTime(duration time.Duration) time.Time {
	return time.Now().Add(duration)
}

func subSecs(duration time.Duration) time.Time {
	return subTime(second * duration)
}

func subMins(duration time.Duration) time.Time {
	return subTime(minute * duration)
}

func addMins(duration time.Duration) time.Time {
	return addTime(minute*duration + second)
}

func subHours(duration time.Duration) time.Time {
	return subTime(hour * duration)
}

func addHours(duration time.Duration) time.Time {
	return addTime(hour*duration + second)
}

func subHoursAndMinutes(hours, mins time.Duration) time.Time {
	return subTime(hour*hours + minute*mins)
}

func subDays(duration time.Duration) time.Time {
	return subTime(day * duration)
}

func subWeeks(duration time.Duration) time.Time {
	return subTime(week * duration)
}

func subMonths(months int) time.Time {
	return wantAddMonths(time.Now().Truncate(time.Second), -months)
}

func subMonthsAndDays(months, days int) time.Time {
	date := wantAddMonths(time.Now().Truncate(time.Second), -months)
	return date.AddDate(0, 0, -days)
}

// wantAddMonths is a deliberately naive oracle for addMonthsClamped: it leans on
// time.Date's normalization instead of manual month/year arithmetic, so the test
// doesn't just mirror the production implementation.
func wantAddMonths(t time.Time, months int) time.Time {
	first := time.Date(t.Year(), t.Month()+time.Month(months), 1, 0, 0, 0, 0, t.Location())
	lastDay := first.AddDate(0, 1, -1).Day()

	return time.Date(
		first.Year(), first.Month(), min(t.Day(), lastDay),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location(),
	)
}

func TestWantAddMonths(t *testing.T) {
	cases := []struct {
		name   string
		date   time.Time
		months int
		want   time.Time
	}{
		{
			name:   "leap year February",
			date:   time.Date(2024, time.March, 31, 12, 0, 0, 0, time.UTC),
			months: -1,
			want:   time.Date(2024, time.February, 29, 12, 0, 0, 0, time.UTC),
		},
		{
			name:   "non-leap year February",
			date:   time.Date(2025, time.March, 31, 12, 0, 0, 0, time.UTC),
			months: -1,
			want:   time.Date(2025, time.February, 28, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := wantAddMonths(tc.date, tc.months); !got.Equal(tc.want) {
				t.Errorf("want %s, got %s", tc.want, got)
			}
		})
	}
}

// assertResult checks the outcome of a parse call.
// Safe for goroutines: uses t.Errorf, not t.Fatalf.
func assertResult(t *testing.T, want string, got string, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("want nil error, got %v instead", err)
		return
	}

	if got != want {
		t.Errorf("want %q, got %q instead", want, got)
	}
}
