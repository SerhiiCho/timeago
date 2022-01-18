package utils

import "time"

func GetTimestampOfPastDate(subDuration time.Duration) int {
	return int(time.Now().Add(-subDuration).UnixNano() / 1000000000)
}

func SmallSubTime(d time.Duration) string {
	return time.Now().Add(d).Format("2006-01-02 15:04:05")
}

func BigSubTime(years int, months int, days int) string {
	return time.Now().AddDate(-years, -months, -days).Format("2006-01-02 15:04:05")
}
