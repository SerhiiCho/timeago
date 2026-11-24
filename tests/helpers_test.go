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
	month  time.Duration = day * 30
	year   time.Duration = day * 365
)

func unixFromPastDate(subDuration time.Duration) int {
	return int(subTime(subDuration).UnixNano() / 1000000000)
}

func unixFromFutureDate(duration time.Duration) int {
	return int(addTime(duration).Add(time.Second - time.Nanosecond).Unix())
}

func subYears(duration time.Duration) time.Time {
	return subTime(year * duration)
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

func subMonths(duration time.Duration) time.Time {
	return subTime(month * duration)
}

func subMonthsAndDays(months, days time.Duration) time.Time {
	return subTime(month*months + day*days)
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
