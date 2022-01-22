package tests

import (
	"time"

	"github.com/SerhiiCho/timeago"
)

var (
	second time.Duration = time.Second
	minute time.Duration = time.Minute
	hour   time.Duration = time.Hour
	day    time.Duration = hour * 24
	week   time.Duration = day * 7
)

func subTime(duration time.Duration) string {
	return time.Now().Add(-duration).Format("2006-01-02 15:04:05")
}

func subMinutes(duration time.Duration) string {
	return subTime(minute * duration)
}

func subHours(duration time.Duration) string {
	return subTime(hour * duration)
}

func subDays(duration time.Duration) string {
	return subTime(day * duration)
}

func subWeeks(duration time.Duration) string {
	return subTime(week * duration)
}

func subMonths(duration time.Duration) string {
	daysInMonth := getLastDayOfMonth(time.Now()).Day()
	result := duration * time.Duration(daysInMonth) * day
	return subTime(result)
}

func subYears(duration time.Duration) string {
	lastDayOfTheYear := time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	daysInYear := time.Duration(lastDayOfTheYear.YearDay())

	return subTime(duration * daysInYear * day)
}

func getLastDayOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func smallSubTime(d time.Duration) string {
	return time.Now().Add(d).Format("2006-01-02 15:04:05")
}

func bigSubTime(years int, months int, days int) string {
	return time.Now().AddDate(-years, -months, -days).Format("2006-01-02 15:04:05")
}

func setup() {
	timeago.SetConfig(timeago.Config{
		Location: "Europe/Kiev",
	})
}
