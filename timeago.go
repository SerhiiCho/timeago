package timeago

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/SerhiiCho/timeago/v3/config"
)

var cachedJsonResults = map[string]LangSet{}
var options = []Option{}
var langSet *LangSet

var conf = &config.Config{
	Language:     "en",
	Location:     "",
	Translations: []config.Translation{},
}

// Parse coverts given datetime into `x time ago` format.
// First argument can have 3 types:
// 1. int (Unix timestamp)
// 2. time.Time (Type from Go time package)
// 3. string (Datetime string in format 'YYYY-MM-DD HH:MM:SS')
func Parse(datetime interface{}, opts ...Option) string {
	options = []Option{}
	var input time.Time

	switch date := datetime.(type) {
	case int:
		input = parseTimestampIntoTime(date)
	case string:
		input = parseStrIntoTime(date)
	default:
		input = datetime.(time.Time)
	}

	enableOptions(opts)

	return calculateTimeAgo(input)
}

func Configure(c *config.Config) {
	if c.Language != "" {
		conf.Language = c.Language
	}

	if c.Location != "" {
		conf.Location = c.Location
	}

	if len(c.Translations) > 0 {
		conf.Translations = c.Translations
	}
}

func parseStrIntoTime(datetime string) time.Time {
	if !conf.LocationIsSet() {
		parsedTime, _ := time.Parse("2006-01-02 15:04:05", datetime)
		return parsedTime
	}

	location, err := time.LoadLocation(conf.Location)

	if err != nil {
		log.Fatalf("[Timeago]: ERROR: %v", err)
	}

	parsedTime, _ := time.ParseInLocation("2006-01-02 15:04:05", datetime, location)

	return parsedTime
}

func calculateTimeAgo(datetime time.Time) string {
	now := time.Now()

	if conf.LocationIsSet() {
		now = applyLocationToTime(now)
	}

	seconds := now.Sub(datetime).Seconds()

	if seconds < 0 {
		enableOption(Upcoming)
		seconds = math.Abs(seconds)
	}

	langSet = NewLangSet()

	switch {
	case seconds < 60 && optionIsEnabled("online"):
		return langSet.Online
	case seconds < 0:
		return getWords(langSet.Second, 0)
	}

	timeUnit := generateTimeUnit(int(seconds))

	return timeUnit
}

func generateTimeUnit(seconds int) string {
	minutes, hours, days, weeks, months, years := getTimeCalculations(float64(seconds))

	switch {
	case optionIsEnabled("online") && seconds < 60:
		return langSet.Online
	case optionIsEnabled("justNow") && seconds < 60:
		return langSet.JustNow
	case seconds < 60:
		return getWords(langSet.Second, seconds)
	case minutes < 60:
		return getWords(langSet.Minute, minutes)
	case hours < 24:
		return getWords(langSet.Hour, hours)
	case days < 7:
		return getWords(langSet.Day, days)
	case weeks < 4:
		return getWords(langSet.Week, weeks)
	case months < 12:
		if months == 0 {
			months = 1
		}

		return getWords(langSet.Month, months)
	}

	return getWords(langSet.Year, years)
}

// getWords decides rather the word must be singular or plural,
// and depending on the result it adds the correct word after
// the time number
func getWords(final LangForms, num int) string {
	form := identifyLocaleForm(num)

	result := langSet.Format
	result = strings.Replace(result, "{timeUnit}", final[form], -1)
	result = strings.Replace(result, "{num}", strconv.Itoa(num), -1)

	if optionIsEnabled("noSuffix") || optionIsEnabled("upcoming") {
		result = strings.Replace(result, "{ago}", "", -1)
		return strings.Trim(result, " ")
	}

	return strings.Replace(result, "{ago}", langSet.Ago, -1)
}

func identifyLocaleForm(num int) string {
	rule, err := identifyGrammarRules(num, conf.Language)

	if err != nil {
		fmt.Println(err)
		return "other"
	}

	switch {
	case rule.Zero:
		return "zero"
	case rule.One:
		return "one"
	case rule.Few:
		return "few"
	case rule.Two:
		return "two"
	case rule.Many:
		return "many"
	}

	return "other"
}

func applyLocationToTime(datetime time.Time) time.Time {
	location, err := time.LoadLocation(conf.Location)

	if err != nil {
		log.Fatalf("[Timeago]: Location error: %v\n", err)
	}

	return datetime.In(location)
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
