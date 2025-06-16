package utils

import (
	"fmt"
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

func UnixFromPastDate(subDuration time.Duration) int {
	return int(SubTime(subDuration).UnixNano() / 1000000000)
}

func UnixFromFutureDate(duration time.Duration) int {
	return int(AddTime(duration).UnixNano() / 1000000000)
}

func Errorf(msg string, a ...interface{}) error {
	return fmt.Errorf("[Timeago]: "+msg, a...)
}

func SubTime(duration time.Duration) time.Time {
	return time.Now().Add(-duration)
}

func AddTime(duration time.Duration) time.Time {
	return time.Now().Add(duration)
}

func SubSeconds(duration time.Duration) time.Time {
	return SubTime(second * duration)
}

func SubMinutes(duration time.Duration) time.Time {
	return SubTime(minute * duration)
}

func AddMinutes(duration time.Duration) time.Time {
	return AddTime(minute * duration)
}

func SubHours(duration time.Duration) time.Time {
	return SubTime(hour * duration)
}

func AddHours(duration time.Duration) time.Time {
	return AddTime(hour * duration)
}

func SubDays(duration time.Duration) time.Time {
	return SubTime(day * duration)
}

func SubWeeks(duration time.Duration) time.Time {
	return SubTime(week * duration)
}

func SubMonths(duration time.Duration) time.Time {
	return SubTime(month * duration)
}

func SubYears(duration time.Duration) time.Time {
	return SubTime(year * duration)
}
