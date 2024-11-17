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

var (
	cachedJsonResults = map[string]*LangSet{}
	options           = []Option{}
	langSet           *LangSet
	conf              = config.New("en", "", []config.Translation{})
)

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

	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, location())

	if err != nil {
		log.Fatalf("[Timeago]: ERROR: %v", err)
	}

	return parsedTime
}

func location() *time.Location {
	if !conf.LocationIsSet() {
		return time.Now().Location()
	}

	loc, err := time.LoadLocation(conf.Location)

	if err != nil {
		log.Fatalf("[Timeago]: ERROR: %v", err)
	}

	return loc
}

func calculateTimeAgo(t time.Time) string {
	now := time.Now()

	if conf.LocationIsSet() {
		loc := location()
		t = t.In(loc)
		now = now.In(loc)
	}

	seconds := int(now.Sub(t).Seconds())

	if seconds < 0 {
		enableOption(Upcoming)
		seconds = -seconds
	}

	langSet = NewLangSet()

	return generateTimeUnit(seconds)
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

func getTimeCalculations(seconds float64) (int, int, int, int, int, int) {
	minutes := math.Round(seconds / 60)
	hours := math.Round(seconds / 3600)
	days := math.Round(seconds / 86400)
	weeks := math.Round(seconds / 604800)
	months := math.Round(seconds / 2629440)
	years := math.Round(seconds / 31553280)

	return int(minutes), int(hours), int(days), int(weeks), int(months), int(years)
}
