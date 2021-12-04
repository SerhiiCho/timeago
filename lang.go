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
			"cpecial": trans().SecondsSpecial,
		},
		"minutes": {
			"single":  trans().Minute,
			"plural":  trans().Minutes,
			"cpecial": trans().MinutesSpecial,
		},
		"hours": {
			"single":  trans().Hour,
			"plural":  trans().Hours,
			"cpecial": trans().HoursSpecial,
		},
		"days": {
			"single":  trans().Day,
			"plural":  trans().Days,
			"cpecial": trans().DaysSpecial,
		},
		"weeks": {
			"single":  trans().Week,
			"plural":  trans().Weeks,
			"cpecial": trans().WeeksSpecial,
		},
		"months": {
			"single":  trans().Month,
			"plural":  trans().Months,
			"cpecial": trans().MonthsSpecial,
		},
		"years": {
			"single":  trans().Year,
			"plural":  trans().Years,
			"cpecial": trans().YearsSpecial,
		},
	}
}

func getLanguageForm(num int) string {
	var form string
	lastDigit := getLastNumber(num)
	rule := getRules(num, lastDigit)[language]

	switch {
	case rule.Single:
		return "single"
	case rule.Plural:
		return "plural"
	}

	if rule.Special != nil {
		for _, isPassing := range rule.Special {
			if isPassing {
				return "special"
			}
		}
	}

	fmt.Errorf("Provided rules don't apply to a number %d", num)

	return form
}
