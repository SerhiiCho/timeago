package rules

import (
	"github.com/SerhiiCho/timeago/models"
)

// number: The number of seconds, minutes, hours, days, weeks, months and years.
// lastDigit: Last digit of the number, if number with 1 digit return it.
func ru(number int, lastDigit int) models.Rule {
	return models.Rule{
		Single: lastDigit == 1 || number == 0,
		Plural: lastDigit >= 2 && lastDigit < 5,
		Special: []bool{
			number >= 5 && number <= 20,
			lastDigit == 0,
			lastDigit >= 5 && lastDigit <= 9,
		},
	}
}
