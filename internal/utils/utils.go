package utils

import (
	"fmt"
	"time"
)

func UnixFromPastDate(subDuration time.Duration) int {
	return int(time.Now().Add(-subDuration).UnixNano() / 1000000000)
}

func Errorf(msg string, a ...interface{}) error {
	return fmt.Errorf("[Timeago]: "+msg, a...)
}
