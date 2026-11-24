package timeago

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// parser holds the state of a single Parse call. Keeping it local
// makes Parse safe to call from multiple goroutines.
type parser struct {
	conf    Config
	langSet *LangSet
	options []opt
}

type timeNumbers struct {
	Seconds int
	Minutes int
	Hours   int
	Days    int
	Weeks   int
	Months  int
	Years   int
}

func (p *parser) dateToTime(date any) (time.Time, error) {
	switch userDate := date.(type) {
	case int:
		return unixToTime(userDate), nil
	case int32:
		return unixToTime(int(userDate)), nil
	case int64:
		return unixToTime(int(userDate)), nil
	case uint:
		return unixToTime(int(userDate)), nil
	case uint32:
		return unixToTime(int(userDate)), nil
	case uint64:
		return unixToTime(int(userDate)), nil
	case string:
		if isUnsignedInteger(userDate) {
			return strTimestampToTime(userDate)
		}
		return p.strToTime(userDate)
	default:
		parsed, ok := date.(time.Time)
		if !ok {
			return time.Time{}, errorf("unsupported date type %T", date)
		}
		return parsed, nil
	}
}

func strTimestampToTime(userDate string) (time.Time, error) {
	sec, err := strconv.Atoi(userDate)
	if err != nil {
		return time.Time{}, err
	}

	return unixToTime(sec), nil
}

func (p *parser) strToTime(userDate string) (time.Time, error) {
	if !p.conf.isLocationProvided() {
		parsedTime, _ := time.Parse(time.DateTime, userDate)
		return parsedTime, nil
	}

	loc, err := p.conf.location()
	if err != nil {
		return time.Time{}, err
	}

	parsedTime, err := time.ParseInLocation(time.DateTime, userDate, loc)
	if err != nil {
		return time.Time{}, errorf("%v", err)
	}

	return parsedTime, nil
}

func (p *parser) computeTimeSince(t time.Time) (string, error) {
	now := time.Now()

	// Adjust times based on location if provided
	var err error
	if t, now, err = p.adjustTimesForLocation(t, now); err != nil {
		return "", err
	}

	timeInSec := p.computeTimeDifference(t, now)

	langSet, err := p.newLangSet()
	if err != nil {
		return "", err
	}

	p.langSet = langSet

	if p.optionIsEnabled(OptOnline) && timeInSec < p.conf.OnlineThreshold {
		return langSet.Online, nil
	}

	if p.optionIsEnabled(OptJustNow) && timeInSec < p.conf.JustNowThreshold {
		return langSet.JustNow, nil
	}

	langForms, timeNum := p.findLangForms(timeInSec)

	timeUnit, err := p.computeTimeUnit(langForms, timeNum)
	if err != nil {
		return "", err
	}

	return p.mergeFinalOutput(timeNum, timeUnit, p.computeSuffix()), nil
}

// adjustTimesForLocation adjusts the given times based on the provided location.
func (p *parser) adjustTimesForLocation(t, now time.Time) (time.Time, time.Time, error) {
	if !p.conf.isLocationProvided() {
		return t, now, nil
	}

	loc, err := p.conf.location()
	if err != nil {
		return t, now, err
	}

	return t.In(loc), now.In(loc), nil
}

// computeTimeDifference returns the absolute time difference in seconds.
func (p *parser) computeTimeDifference(t, now time.Time) int {
	timeInSec := int(now.Sub(t).Seconds())
	if timeInSec < 0 {
		p.enableOption(OptUpcoming)
		return -timeInSec
	}

	return timeInSec
}

func (p *parser) mergeFinalOutput(timeNum int, timeUnit, suffix string) string {
	replacer := strings.NewReplacer(
		"{timeUnit}", timeUnit,
		"{num}", strconv.Itoa(timeNum),
		"{ago}", suffix,
	)

	return strings.TrimSpace(replacer.Replace(p.langSet.Format))
}

func (p *parser) findLangForms(timeInSec int) (LangForms, int) {
	nums := calculateTimeNumbers(float64(timeInSec))

	switch {
	case timeInSec < 60:
		return p.langSet.Second, nums.Seconds
	case nums.Minutes < 60:
		return p.langSet.Minute, nums.Minutes
	case nums.Hours < 24:
		return p.langSet.Hour, nums.Hours
	case nums.Days < 7:
		return p.langSet.Day, nums.Days
	case nums.Weeks < 4:
		return p.langSet.Week, nums.Weeks
	case nums.Months < 12:
		if nums.Months == 0 {
			nums.Months = 1
		}
		return p.langSet.Month, nums.Months
	}

	return p.langSet.Year, nums.Years
}

func (p *parser) computeSuffix() string {
	if p.optionIsEnabled(OptNoSuffix) || p.optionIsEnabled(OptUpcoming) {
		return ""
	}
	return p.langSet.Ago
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

func (p *parser) computeTimeUnit(langForm LangForms, num int) (string, error) {
	form, err := p.timeUnitForm(num)
	if err != nil {
		return "", err
	}

	if unit, ok := langForm[form]; ok {
		return unit, nil
	}

	return langForm["other"], nil
}

func (p *parser) timeUnitForm(num int) (string, error) {
	rule, err := identifyGrammarRules(num, p.conf.Language)
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

func (p *parser) newLangSet() (*LangSet, error) {
	filePath := fmt.Sprintf("langs/%s.json", p.conf.Language)

	cacheMu.RLock()
	cached, ok := cachedJsonRes[filePath]
	cacheMu.RUnlock()

	if ok {
		return cached, nil
	}

	langSet, err := parseLangSet(filePath)
	if err != nil {
		return nil, err
	}

	langSet = applyCustomTranslations(langSet, p.conf.Translations)

	cacheMu.Lock()
	cachedJsonRes[filePath] = langSet
	cacheMu.Unlock()

	return langSet, nil
}

func (p *parser) enableOption(opt opt) {
	p.options = append(p.options, opt)
}

func (p *parser) optionIsEnabled(opt opt) bool {
	for i := range p.options {
		if p.options[i] == opt {
			return true
		}
	}

	return false
}
