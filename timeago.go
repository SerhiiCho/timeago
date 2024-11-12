package timeago

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/SerhiiCho/timeago/v3/config"
	"github.com/SerhiiCho/timeago/v3/ctx"
	"github.com/SerhiiCho/timeago/v3/langset"
	"github.com/SerhiiCho/timeago/v3/option"
)

var cachedJsonResults = map[string]langset.LangSet{}
var options = []option.Option{}

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
	langSet := langset.New(context)

	fmt.Printf("-------> %#v\n", langSet)
	// switch {
	// case seconds < 60 && optionIsEnabled(option.Online):
	// 	return trans().Online
	// case seconds < 0:
	// 	return getWords("seconds", 0)
	// }
	return ""
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
