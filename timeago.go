package timeago

import (
	"fmt"
	"log"
	"math"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/SerhiiCho/timeago/models"
	"github.com/SerhiiCho/timeago/utils"
)

var cache = map[string]models.Lang{}

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
	case seconds < 60 && option == "online":
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
	form := getLanguageForm(num)
	time := getTimeTranslations()

	translation := time[timeKind][form]

	return strconv.Itoa(num) + " " + translation + " " + trans().Ago
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
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		panic("No caller information")
	}

	rootPath := path.Dir(filename)

	filePath := fmt.Sprintf(rootPath+"/langs/%s.json", language)

	if cachedResult, ok := cache[filePath]; ok {
		return cachedResult
	}

	thereIsFile, err := utils.FileExists(filePath)

	if !thereIsFile {
		log.Fatalf("File with the path: %s, doesn't exist", filePath)
	}

	if err != nil {
		log.Fatalf("Error while trying to read file %s. Error: %v", filePath, err)
	}

	parseResult := parseNeededFile(filePath)

	cache[filePath] = parseResult

	return parseResult
}
