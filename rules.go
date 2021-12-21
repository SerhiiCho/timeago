package timeago

import "github.com/SerhiiCho/timeago/models"

func getRules(number, lastDigit int) map[string]models.Rule {
	return map[string]models.Rule{
		"en": {
			Single: number == 1,
			Plural: number > 1 || number == 0,
		},
		"ru": {
			Single: lastDigit == 1 || number == 0,
			Plural: lastDigit >= 2 && lastDigit < 5,
			Special: []bool{
				number >= 5 && number <= 20,
				lastDigit == 0,
				lastDigit >= 5 && lastDigit <= 9,
			},
		},
		"uk": {
			Single: lastDigit == 1 || number == 0,
			Plural: lastDigit >= 2 && lastDigit < 5,
			Special: []bool{
				number >= 5 && number <= 20,
				lastDigit == 0,
				lastDigit >= 5 && lastDigit <= 9,
			},
		},
	}
}
