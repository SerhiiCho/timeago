package timeago

import (
	"strconv"
	"strings"
	"time"
)

// Take coverts given datetime into `x time ago` format.
// For displaying `Online` word if date interval within
// 60 seconds, add `|online` flag to the datetime string.
// Format must be [year-month-day hours:minutes:seconds}
func Take(datetime string) string {
	option, hasOption := getOption(&datetime)

	loc, _ := time.LoadLocation(location)
	parsedTime, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)
	seconds := int(time.Now().In(loc).Sub(parsedTime).Seconds())

	switch {
	case seconds < 0 && option == "online":
		return trans("online")
	case seconds < 0:
		return getWords("seconds", 0)
	}

	return calculateTheResult(seconds, hasOption, option)
}

func calculateTheResult(seconds int, hasOption bool, option string) string {
	minutes, hours, days, weeks, months, years := getTimeCalculations(seconds)

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

func getTimeCalculations(seconds int) (int, int, int, int, int, int) {
	minutes := seconds / 60
	hours := seconds / 3600
	days := seconds / 86400
	weeks := seconds / 604800
	months := seconds / 2629440
	years := seconds / 31553280

	return minutes, hours, days, weeks, months, years
}

func getLastNumber(num int) int {
	numStr := strconv.Itoa(num)
	result, _ := strconv.Atoi(numStr[len(numStr)-1:])

	return result
}

// getWords desides rather the word must be singular or plural,
// and depending on the result it adds the correct word after
// the time number
func getWords(timeKind string, num int) string {
	lastNum := getLastNumber(num)
	index := 2

	switch {
	case lastNum == 1 && num == 11:
		index = 2
		break
	case lastNum == 1 && language == "ru" || num == 1 && language == "en":
		index = 0
		break
	case lastNum > 1 && lastNum < 5:
		index = 1
	}

	timeTrans := getTimeTranslations()

	return strconv.Itoa(num) + " " + timeTrans[timeKind][index] + " " + trans("ago")
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

// getOption check if datetime has option with time,
// if yes, it will return this option and remove it
// from datetime
func getOption(datetime *string) (string, bool) {
	date := *datetime
	spittedDateString := strings.Split(date, "|")

	if len(spittedDateString) > 1 {
		*datetime = spittedDateString[0]
		return spittedDateString[1], true
	}

	return "", false
}

func trans(key string) string {
	if language == "ru" {
		return getRussian()[key]
	}

	return getEnglish()[key]
}
