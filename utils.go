package timeago

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

func unixToTime(userDate int) time.Time {
	return time.Unix(int64(userDate), 0)
}

func uintToTime(userDate uint64) (time.Time, error) {
	if userDate > math.MaxInt64 {
		return time.Time{}, errorf("timestamp %d overflows int64", userDate)
	}

	return unixToTime(int(userDate)), nil
}

func parseLangSet(filePath string) (*LangSet, error) {
	content, err := langsFS.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read language file: %w", err)
	}

	langSet := &LangSet{}

	if err := json.Unmarshal(content, &langSet); err != nil {
		return nil, fmt.Errorf("could not unmarshal language file: %w", err)
	}

	return langSet, nil
}

func isUnsignedInteger(date string) bool {
	if len(date) == 0 {
		return false
	}

	for i := 0; i < len(date); i++ {
		char := date[i]
		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}

func errorf(msg string, a ...any) error {
	return fmt.Errorf("[Timeago]: "+msg, a...)
}
