package timeago

import (
	"strings"
)

// Take coverts given datetime into `x time ago` format.
// For displaying `Online` word if date interval within
// 60 seconds, add `|online` flag to the datetime string.
func Take(datetime string) string {
	option, hasOption := hasOption(datetime)

	if hasOption && option == "online" {
		return ""
	}

	return ""
}

func hasOption(datetime string) (string, bool) {
	spittedDateString := strings.Split(datetime, "|")

	if len(spittedDateString) > 1 {
		return spittedDateString[1], true
	}

	return "", false
}
