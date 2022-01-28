package timeago

import (
	"log"
	"math"
	"strconv"
	"time"
)

var cachedJsonResults = map[string]lang{}
var globalOptions = []string{}

// Parse coverts given datetime into `x time ago` format.
// First argument can have 3 types:
// 1. int (Unix timestamp)
// 2. time.Time (Type from Go time package)
// 3. string (Datetime string in format 'YYYY-MM-DD HH:MM:SS')
func Parse(datetime interface{}, options ...string) string {
	var input time.Time

	switch date := datetime.(type) {
	case int:
		input = parseTimestampToTime(date)
	case string:
		if config.Location == "" {
			parsedTime, _ := time.Parse("2006-01-02 15:04:05", date)
			input = parsedTime
		} else {
			location, err := time.LoadLocation(config.Location)

			if err != nil {
				log.Fatalf("Error in timeago package: %v", err)
			}

			parsedTime, _ := time.ParseInLocation("2006-01-02 15:04:05", date, location)
			input = parsedTime
		}
	default:
		input = datetime.(time.Time)
	}

	globalOptions = options

	return process(input)
}

func process(datetime time.Time) string {
	now := time.Now()

	if config.Location != "" {
		location, err := time.LoadLocation(config.Location)

		if err != nil {
			log.Fatalf("Location error in timeago package: %v\n", err)
		} else {
			now = now.In(location)
		}
	}

	seconds := now.Sub(datetime).Seconds()

	if seconds < 0 {
		enableOption("upcoming")
		seconds = math.Abs(seconds)
	}

	switch {
	case seconds < 60 && optionIsEnabled("online"):
		return trans().Online
	case seconds < 0:
		return getWords("seconds", 0)
	}

	return calculateTheResult(int(seconds))
}

func calculateTheResult(seconds int) string {
	minutes, hours, days, weeks, months, years := getTimeCalculations(float64(seconds))

	switch {
	case optionIsEnabled("online") && seconds < 60:
		return trans().Online
	case optionIsEnabled("justNow") && seconds < 60:
		return trans().JustNow
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

// getWords decides rather the word must be singular or plural,
// and depending on the result it adds the correct word after
// the time number
func getWords(timeKind string, num int) string {
	form := getLanguageForm(num)
	time := getTimeTranslations()

	translation := time[timeKind][form]
	result := strconv.Itoa(num) + " " + translation

	if optionIsEnabled("noSuffix") || optionIsEnabled("upcoming") {
		return result
	}

	return result + " " + trans().Ago
}

func optionIsEnabled(searchOption string) bool {
	for _, option := range globalOptions {
		if option == searchOption {
			return true
		}
	}

	return false
}

func enableOption(option string) {
	globalOptions = append(globalOptions, option)
}
