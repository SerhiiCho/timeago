package timeago

import (
	"fmt"
	"log"
	"path"
	"runtime"
)

type lang struct {
	Ago            string
	Online         string
	Second         string
	Seconds        string
	SecondsSpecial string
	Minute         string
	Minutes        string
	MinutesSpecial string
	Hour           string
	Hours          string
	HoursSpecial   string
	Day            string
	Days           string
	DaysSpecial    string
	Week           string
	Weeks          string
	WeeksSpecial   string
	Month          string
	Months         string
	MonthsSpecial  string
	Year           string
	Years          string
	YearsSpecial   string
}

// getTimeTranslations returns array of translations for different
// cases. For example `1 second` must not have `s` at the end
// but `2 seconds` requires `s`. So this method keeps all
// possible options for the translated word.
func getTimeTranslations() map[string]map[string]string {
	return map[string]map[string]string{
		"seconds": {
			"single":  trans().Second,
			"plural":  trans().Seconds,
			"special": trans().SecondsSpecial,
		},
		"minutes": {
			"single":  trans().Minute,
			"plural":  trans().Minutes,
			"special": trans().MinutesSpecial,
		},
		"hours": {
			"single":  trans().Hour,
			"plural":  trans().Hours,
			"special": trans().HoursSpecial,
		},
		"days": {
			"single":  trans().Day,
			"plural":  trans().Days,
			"special": trans().DaysSpecial,
		},
		"weeks": {
			"single":  trans().Week,
			"plural":  trans().Weeks,
			"special": trans().WeeksSpecial,
		},
		"months": {
			"single":  trans().Month,
			"plural":  trans().Months,
			"special": trans().MonthsSpecial,
		},
		"years": {
			"single":  trans().Year,
			"plural":  trans().Years,
			"special": trans().YearsSpecial,
		},
	}
}

func getLanguageForm(num int) string {
	var form string
	lastDigit := getLastNumberDigit(num)
	rule := getRules(num, lastDigit)[config.Language]

	switch {
	case rule.Special:
		return "special"
	case rule.Single:
		return "single"
	case rule.Plural:
		return "plural"
	}

	fmt.Printf("Provided rules don't apply to a number %d", num)

	return form
}

func trans() lang {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		panic("No caller information")
	}

	rootPath := path.Dir(filename)

	filePath := fmt.Sprintf(rootPath+"/langs/%s.json", config.Language)

	if cachedResult, ok := cachedJsonResults[filePath]; ok {
		return cachedResult
	}

	thereIsFile, err := fileExists(filePath)

	if !thereIsFile {
		log.Fatalf("File with the path: %s, doesn't exist", filePath)
	}

	if err != nil {
		log.Fatalf("Error while trying to read file %s. Error: %v", filePath, err)
	}

	parseResult := parseJsonIntoLang(filePath)

	cachedJsonResults[filePath] = parseResult

	return parseResult
}
