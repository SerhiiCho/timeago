package timeago

import (
	"fmt"
)

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
	lastDigit := getLastNumber(num)
	rule := getRules(num, lastDigit)[language]

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
