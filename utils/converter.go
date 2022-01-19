package utils

import "time"

func ParseertTimestampToString(timestamp int) string {
	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}
