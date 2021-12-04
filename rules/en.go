package rules

import (
	"github.com/SerhiiCho/timeago/models"
)

// number: The number of seconds, minutes, hours, days, weeks, months and years.
// lastDigit: Last digit of the number, if number with 1 digit return it.
func en(number int, lastDigit int) models.Rule {
	return models.Rule{
		Single: number == 0 || number == 1,
		Plural: number > 1,
	}
}
