package timeago

import (
	"math"
	"strconv"
	"strings"
	"time"
)

var (
	// cachedJsonRes saves parsed JSON translations to prevent
	// parsing the same JSON file multiple times.
	cachedJsonRes = map[string]*LangSet{}

	// options is a list of options that modify the final output.
	// Some options are noSuffix, upcoming, online, and justNow.
	options = []Option{}

	// conf is configuration provided by the user.
	conf = defaultConfig()

	// langSet is a pointer to the current language set that
	// is currently being used.
	langSet *LangSet
)

// Parse coverts privided datetime into `x time ago` format.
// The first argument can have 3 types:
// 1. int (Unix timestamp)
// 2. time.Time (Type from Go time package)
// 3. string (Datetime string in format 'YYYY-MM-DD HH:MM:SS')
func Parse(datetime interface{}, opts ...Option) (string, error) {
	options = []Option{}
	langSet = nil

	var input time.Time
	var err error

	switch date := datetime.(type) {
	case int:
		input = parseTimestampIntoTime(date)
	case string:
		input, err = parseStrIntoTime(date)
	default:
		input = datetime.(time.Time)
	}

	if err != nil {
		return "", err
	}

	enableOptions(opts)

	return calculateTimeAgo(input)
}

// Configure applies the given configuration to the timeago without
// overriding the previous configuration. It will only override the
// provided configuration. If you want to override the previous
// configurations, use Reconfigure function instead.
func Configure(c Config) {
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

// Reconfigure reconfigures the timeago with the provided configuration.
// It will override the previous configuration with the new one.
func Reconfigure(c Config) {
	conf = defaultConfig()
	cachedJsonRes = map[string]*LangSet{}

	Configure(c)
}

func defaultConfig() *Config {
	return NewConfig("en", "", []LangSet{})
}

func parseStrIntoTime(datetime string) (time.Time, error) {
	if !conf.isLocationProvided() {
		parsedTime, _ := time.Parse("2006-01-02 15:04:05", datetime)
		return parsedTime, nil
	}

	loc, err := location()

	if err != nil {
		return time.Time{}, err
	}

	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)

	if err != nil {
		return time.Time{}, createError("%v", err)
	}

	return parsedTime, nil
}

func location() (*time.Location, error) {
	if !conf.isLocationProvided() {
		return time.Now().Location(), nil
	}

	loc, err := time.LoadLocation(conf.Location)

	if err != nil {
		return nil, createError("%v", err)
	}

	return loc, nil
}

func calculateTimeAgo(t time.Time) (string, error) {
	now := time.Now()

	if conf.isLocationProvided() {
		loc, err := location()

		if err != nil {
			return "", err
		}

		t = t.In(loc)
		now = now.In(loc)
	}

	seconds := int(now.Sub(t).Seconds())

	if seconds < 0 {
		enableOption(Upcoming)
		seconds = -seconds
	}

	set, err := newLangSet()

	if err != nil {
		return "", err
	}

	langSet = set

	return generateTimeUnit(seconds)
}

func generateTimeUnit(seconds int) (string, error) {
	minutes, hours, days, weeks, months, years := getTimeCalculations(float64(seconds))

	switch {
	case optionIsEnabled("online") && seconds < 60:
		return langSet.Online, nil
	case optionIsEnabled("justNow") && seconds < 60:
		return langSet.JustNow, nil
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
// the time number.
func getWords(langForm LangForms, num int) (string, error) {
	timeUnit, err := finalTimeUnit(langForm, num)

	if err != nil {
		return "", err
	}

	res := langSet.Format
	res = strings.Replace(res, "{timeUnit}", timeUnit, -1)
	res = strings.Replace(res, "{num}", strconv.Itoa(num), -1)

	if optionIsEnabled("noSuffix") || optionIsEnabled("upcoming") {
		res = strings.Replace(res, "{ago}", "", -1)
		return strings.Trim(res, " "), nil
	}

	return strings.Replace(res, "{ago}", langSet.Ago, -1), nil
}

func finalTimeUnit(langForm LangForms, num int) (string, error) {
	form, err := identifyTimeUnitForm(num)

	if err != nil {
		return "", err
	}

	if unit, ok := langForm[form]; ok {
		return unit, nil
	}

	return langForm["other"], nil
}

func identifyTimeUnitForm(num int) (string, error) {
	rule, err := identifyGrammarRules(num, conf.Language)

	if err != nil {
		return "", err
	}

	switch {
	case rule.Zero:
		return "zero", nil
	case rule.One:
		return "one", nil
	case rule.Few:
		return "few", nil
	case rule.Two:
		return "two", nil
	case rule.Many:
		return "many", nil
	}

	return "other", nil
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
