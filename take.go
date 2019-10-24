package timeago

import (
	"strconv"
	"strings"
	"time"
)

// Take coverts given datetime into `x time ago` format.
// For displaying `Online` word if date interval within
// 60 seconds, add `|online` flag to the datetime string.
func Take(datetime string) string {
	includeTranslations()

	seconds := int(time.Since(time.Now()).Seconds())

	minutes := seconds / 60
	hours := seconds / 3600
	days := seconds / 86400
	weeks := seconds / 604800
	months := seconds / 2629440
	years := seconds / 31553280

	option, hasOption := getOption(datetime)

	switch {
	case hasOption && option == "online" && seconds < 60:
		return trans("online")
	case seconds < 60:
		return getWords("seconds", seconds)
	case minutes < 60:
		return getWords("minutes", minutes)
	case hours < 24:
		return getWords("hours", hours)
	case days < 7:
		return getWords("days", days)
	case weeks < 4:
		return getWords("weeks", weeks)
	case months < 12:
		return getWords("months", months)
	}

	return getWords("years", years)
}

func getLastNumber(num int) int {
	numStr := strconv.Itoa(num)
	result, err := strconv.Atoi(numStr[len(numStr)-1:])

	if err != nil {
		return 0
	}

	return result
}

func getWords(typeOf string, num int) string {
	lastNum := getLastNumber(num)
	var index int

	switch {
	case lastNum == 1 && num == 11:
		index = 2
		break
	case lastNum == 1:
		index = 0
		break
	case lastNum > 1 && lastNum < 5:
		index = 1
		break
	default:
		index = 2
	}

	timeTrans := getTimeTranslations()

	return strconv.Itoa(num) + " " + timeTrans[typeOf][index] + " " + trans("ago")
}

// getTimeTranslations returns array of translations for different
// cases. For example `1 second` must not have `s` at the end
// but `2 seconds` requires `s`. So this method keeps all
// possible options for the translated word.
func getTimeTranslations() map[string][]string {
	return map[string][]string{
		"seconds": {trans("second"), trans("seconds"), trans("seconds2")},
		"minutes": {trans("minute"), trans("minutes"), trans("minutes2")},
		"hours":   {trans("hour"), trans("hours"), trans("hours2")},
		"days":    {trans("day"), trans("days"), trans("days2")},
		"weeks":   {trans("week"), trans("weeks"), trans("weeks2")},
		"months":  {trans("month"), trans("months"), trans("months2")},
		"years":   {trans("year"), trans("years"), trans("years2")},
	}
}

func getOption(datetime string) (string, bool) {
	spittedDateString := strings.Split(datetime, "|")

	if len(spittedDateString) > 1 {
		return spittedDateString[1], true
	}

	return "", false
}
