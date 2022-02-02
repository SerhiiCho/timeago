package tests

import (
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

func subTime(duration time.Duration) time.Time {
	return time.Now().Add(-duration)
}

func subSeconds(duration time.Duration) time.Time {
	return subTime(second * duration)
}

func subMinutes(duration time.Duration) time.Time {
	return subTime(minute * duration)
}

func subHours(duration time.Duration) time.Time {
	return subTime(hour * duration)
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

func subYears(duration time.Duration) time.Time {
	return subTime(year * duration)
}
