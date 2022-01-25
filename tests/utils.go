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
	// we cannot add month and year because months and years cannot have static values
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
	return subTime(duration * getDaysInMonth() * day)
}

func subYears(duration time.Duration) time.Time {
	return subTime(duration * getDaysInYear() * day)
}

func getDaysInYear() time.Duration {
	lastDayOfTheYear := time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	return time.Duration(lastDayOfTheYear.YearDay())
}

func getDaysInMonth() time.Duration {
	return time.Duration(getLastDayOfMonth(time.Now()).Day())
}

func getLastDayOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}
