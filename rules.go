package timeago

type rule struct {
	Single  bool
	Plural  bool
	Special bool
}

func getRules(number, lastDigit int) map[string]rule {
	return map[string]rule{
		"en": {
			Single: number == 1,
			Plural: number > 1 || number == 0,
		},
		"ru": {
			Special: (number >= 5 && number <= 20) || lastDigit == 0 || (lastDigit >= 5 && lastDigit <= 9),
			Single:  lastDigit == 1,
			Plural:  lastDigit >= 2 && lastDigit < 5,
		},
		"uk": {
			Special: (number >= 5 && number <= 20) || lastDigit == 0 || (lastDigit >= 5 && lastDigit <= 9),
			Single:  lastDigit == 1,
			Plural:  lastDigit >= 2 && lastDigit < 5,
		},
		"nl": {
			Single: number == 1,
			Plural: number > 1 || number == 0,
		},
		"de": {
			Single: number == 1,
			Plural: number > 1 || number == 0,
		},
	}
}
