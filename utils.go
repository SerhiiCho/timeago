package timeago

import (
	"time"
)

func parseTimestampIntoTime(timestamp int) time.Time {
	return time.Unix(int64(timestamp), 0)
}
