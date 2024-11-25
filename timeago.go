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
	options = []opt{}

	// conf is configuration provided by the user.
	conf = defaultConfig()

	// langSet is a pointer to the current language set that
	// is currently being used.
	langSet *LangSet
)

type timeNumbers struct {
	Seconds int
	Minutes int
	Hours   int
	Days    int
	Weeks   int
	Months  int
	Years   int
}

// Parse coverts privided datetime into `x time ago` format.
// The first argument can have 3 types:
// 1. int (Unix timestamp)
// 2. time.Time (Type from Go time package)
// 3. string (Datetime string in format 'YYYY-MM-DD HH:MM:SS')
func Parse(datetime interface{}, opts ...opt) (string, error) {
	options = []opt{}
	langSet = nil

	var input time.Time
	var err error

	switch date := datetime.(type) {
	case int:
		input = unixToTime(date)
	case string:
		input, err = strToTime(date)
	default:
		input = datetime.(time.Time)
	}

	if err != nil {
		return "", err
	}

	enableOptions(opts)

	return computeTimeSince(input)
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

func strToTime(datetime string) (time.Time, error) {
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
		return time.Time{}, errorf("%v", err)
	}

	return parsedTime, nil
}

// location loads location from the time package
func location() (*time.Location, error) {
	if !conf.isLocationProvided() {
		return time.Now().Location(), nil
	}

	loc, err := time.LoadLocation(conf.Location)
	if err != nil {
		return nil, errorf("%v", err)
	}

	return loc, nil
}

func computeTimeSince(t time.Time) (string, error) {
	now := time.Now()
	var err error

	// Adjust times based on location if provided
	if t, now, err = adjustTimesForLocation(t, now); err != nil {
		return "", err
	}

	timeInSec := computeTimeDifference(t, now)

	if langSet, err = newLangSet(); err != nil {
		return "", err
	}

	if optionIsEnabled(OptOnline) && timeInSec < 60 {
		return langSet.Online, nil
	}

	if optionIsEnabled(OptJustNow) && timeInSec < 60 {
		return langSet.JustNow, nil
	}

	var timeUnit string

	langForms, timeNum := findLangForms(timeInSec)

	if timeUnit, err = computeTimeUnit(langForms, timeNum); err != nil {
		return "", err
	}

	suffix := computeSuffix()

	return mergeFinalOutput(timeNum, timeUnit, suffix)
}

// adjustTimesForLocation adjusts the given times based on the provided location.
func adjustTimesForLocation(t, now time.Time) (time.Time, time.Time, error) {
	if !conf.isLocationProvided() {
		return t, now, nil
	}

	loc, err := location()
	if err != nil {
		return t, now, err
	}

	return t.In(loc), now.In(loc), nil
}

// computeTimeDifference returns the absolute time difference in seconds.
func computeTimeDifference(t, now time.Time) int {
	timeInSec := int(now.Sub(t).Seconds())
	if timeInSec < 0 {
		enableOption(OptUpcoming)
		return -timeInSec
	}

	return timeInSec
}

func mergeFinalOutput(timeNum int, timeUnit, suffix string) (string, error) {
	replacer := strings.NewReplacer(
		"{timeUnit}", timeUnit,
		"{num}", strconv.Itoa(timeNum),
		"{ago}", suffix,
	)

	out := replacer.Replace(langSet.Format)

	return strings.TrimSpace(out), nil
}

func findLangForms(timeInSec int) (LangForms, int) {
	nums := calculateTimeNumbers(float64(timeInSec))

	switch {
	case timeInSec < 60:
		return langSet.Second, nums.Seconds
	case nums.Minutes < 60:
		return langSet.Minute, nums.Minutes
	case nums.Hours < 24:
		return langSet.Hour, nums.Hours
	case nums.Days < 7:
		return langSet.Day, nums.Days
	case nums.Weeks < 4:
		return langSet.Week, nums.Weeks
	case nums.Months < 12:
		if nums.Months == 0 {
			nums.Months = 1
		}

		return langSet.Month, nums.Months
	}

	return langSet.Year, nums.Years
}

func computeSuffix() string {
	if optionIsEnabled(OptNoSuffix) || optionIsEnabled(OptUpcoming) {
		return ""
	}

	return langSet.Ago
}

func calculateTimeNumbers(seconds float64) timeNumbers {
	minutes := math.Round(seconds / 60)
	hours := math.Round(seconds / 3600)
	days := math.Round(seconds / 86400)
	weeks := math.Round(seconds / 604800)
	months := math.Round(seconds / 2629440)
	years := math.Round(seconds / 31553280)

	return timeNumbers{
		Seconds: int(seconds),
		Minutes: int(minutes),
		Hours:   int(hours),
		Days:    int(days),
		Weeks:   int(weeks),
		Months:  int(months),
		Years:   int(years),
	}
}

func computeTimeUnit(langForm LangForms, num int) (string, error) {
	form, err := timeUnitForm(num)
	if err != nil {
		return "", err
	}

	if unit, ok := langForm[form]; ok {
		return unit, nil
	}

	return langForm["other"], nil
}

func timeUnitForm(num int) (string, error) {
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
