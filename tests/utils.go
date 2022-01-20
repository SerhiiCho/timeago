package tests

import "time"

func smallSubTime(d time.Duration) string {
	return time.Now().Add(d).Format("2006-01-02 15:04:05")
}

func bigSubTime(years int, months int, days int) string {
	return time.Now().AddDate(-years, -months, -days).Format("2006-01-02 15:04:05")
}
