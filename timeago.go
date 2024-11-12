package timeago

import (
	"log"
	"math"
	"strconv"
	"time"

	"github.com/SerhiiCho/timeago/v3/config"
	"github.com/SerhiiCho/timeago/v3/ctx"
	locale "github.com/SerhiiCho/timeago/v3/langset"
	"github.com/SerhiiCho/timeago/v3/option"
)

var cachedJsonResults = map[string]locale.LocaleSet{}
var options = []option.Option{}
var localeSet *locale.LocaleSet

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
func Parse(datetime interface{}, opts ...option.Option) string {
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
		log.Fatalf("Error in timeago package: %v", err)
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
		enableOption(option.Upcoming)
		seconds = math.Abs(seconds)
	}

	context := ctx.New(conf, options)
	localeSet = locale.New(context)

	switch {
	case seconds < 60 && optionIsEnabled("online"):
		return localeSet.Online
	case seconds < 0:
		return getWords(localeSet.Second, 0)
	}

	timeUnit := generateTimeUnit(int(seconds))

	return timeUnit
}

func generateTimeUnit(seconds int) string {
	minutes, hours, days, weeks, months, years := getTimeCalculations(float64(seconds))

	switch {
	case optionIsEnabled("online") && seconds < 60:
		return localeSet.Online
	case optionIsEnabled("justNow") && seconds < 60:
		return localeSet.JustNow
	case seconds < 60:
		return getWords(localeSet.Second, seconds)
	case minutes < 60:
		return getWords(localeSet.Minute, minutes)
	case hours < 24:
		return getWords(localeSet.Hour, hours)
	case days < 7:
		return getWords(localeSet.Day, days)
	case weeks < 4:
		return getWords(localeSet.Week, weeks)
	case months < 12:
		if months == 0 {
			months = 1
		}

		return getWords(localeSet.Month, months)
	}

	return getWords(localeSet.Year, years)
}

// getWords decides rather the word must be singular or plural,
// and depending on the result it adds the correct word after
// the time number
func getWords(final locale.LocaleForms, num int) string {
	form := identifyLocaleForm(num)
	time := getTimeTranslations()

	translation := time[timeUnit][form]
	result := strconv.Itoa(num) + " " + translation

	if optionIsEnabled("noSuffix") || optionIsEnabled("upcoming") {
		return result
	}

	return result
}

func identifyLocaleForm(num int) string {
	rule := identifyGrammarRules(num)[conf.Language]

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
		log.Fatalf("Location error in timeago package: %v\n", err)
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
