package timeago

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/SerhiiCho/timeago/models"
	"github.com/SerhiiCho/timeago/utils"
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
		return trans().Online
	case seconds < 0:
		return getWords("seconds", 0)
	}

	return calculateTheResult(seconds, hasOption, option)
}

func calculateTheResult(seconds int, hasOption bool, option string) string {
	minutes, hours, days, weeks, months, years := getTimeCalculations(float64(seconds))

	switch {
	case hasOption && option == "online" && seconds < 60:
		return trans().Online
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
		if months == 0 {
			months = 1
		}

		return getWords("months", months)
	}

	return getWords("years", years)
}

func getTimeCalculations(seconds float64) (int, int, int, int, int, int) {
	minutes := math.Round(seconds / 60)
	hours := math.Round(seconds / 3600)
	days := math.Round(seconds / 86400)
	weeks := math.Round(seconds / 604800)
	months := math.Round(seconds / 2629440)
	years := math.Round(seconds / 31553280)

	return int(minutes), int(hours), int(days), int(weeks), int(months), int(years)
}

// get the last number of a given integer
func getLastNumber(num int) int {
	numStr := strconv.Itoa(num)
	result, _ := strconv.Atoi(numStr[len(numStr)-1:])

	return result
}

// getWords decides rather the word must be singular or plural,
// and depending on the result it adds the correct word after
// the time number
func getWords(timeKind string, num int) string {
	lastNum := getLastNumber(num)
	index := 2

	switch {
	case lastNum == 1 && num == 11:
		index = 2
	case lastNum == 1 && language == "ru" || num == 1 && language == "en":
		index = 0
	case lastNum > 1 && lastNum < 5:
		index = 1
	}

	timeTrans := getTimeTranslations()

	return strconv.Itoa(num) + " " + timeTrans[timeKind][index] + " " + trans().Ago
}

// getTimeTranslations returns array of translations for different
// cases. For example `1 second` must not have `s` at the end
// but `2 seconds` requires `s`. So this method keeps all
// possible options for the translated word.
func getTimeTranslations() map[string][]string {
	return map[string][]string{
		"seconds": {trans().Second, trans().Seconds, trans().Seconds2},
		"minutes": {trans().Minute, trans().Minutes, trans().Minutes2},
		"hours":   {trans().Hour, trans().Hours, trans().Hours2},
		"days":    {trans().Day, trans().Days, trans().Days2},
		"weeks":   {trans().Week, trans().Weeks, trans().Weeks2},
		"months":  {trans().Month, trans().Months, trans().Months2},
		"years":   {trans().Year, trans().Years, trans().Years2},
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

func trans() models.Lang {
	filePath := fmt.Sprintf("./langs/%s.json", language)

	thereIsFile, err := utils.FileExists(filePath)

	if !thereIsFile {
		log.Fatalf("File with the path: %s, doesn't exist", filePath)
	}

	if err != nil {
		log.Fatalf("Error while trying to read file %s. Error: %v", filePath, err)
	}

	return parseNeededFile(filePath)
}
