package timeago

type rule struct {
	Zero  bool
	One   bool
	Two   bool
	Few   bool
	Many  bool
	Other bool
}

func getRules(num int) map[string]rule {
	lastDigit := num % 10
	// lastTwoDigits := num % 100

	return map[string]rule{
		"en": {
			Zero:  num == 0,
			One:   num == 1,
			Few:   num > 1,
			Two:   num == 2,
			Many:  num > 1,
			Other: num > 1,
		},
		"ru": {
			// Zero: num == 0,
			One: lastDigit == 1,
			// Two:  num == 2,
			Few:  lastDigit == 2 || lastDigit == 3 || lastDigit == 4,
			Many: (num >= 5 && num <= 20) || lastDigit == 0 || (lastDigit >= 5 && lastDigit <= 9),
		},
		"uk": {
			Zero:  num == 0,
			One:   lastDigit == 1,
			Two:   num == 2,
			Few:   num == 2,
			Many:  lastDigit >= 2 && lastDigit < 5,
			Other: (num >= 5 && num <= 20) || lastDigit == 0 || (lastDigit >= 5 && lastDigit <= 9),
		},
		"nl": {
			Zero:  num == 0,
			One:   num == 1,
			Few:   num > 1,
			Two:   num == 2,
			Many:  num > 1,
			Other: num > 1,
		},
		"de": {
			Zero:  num == 0,
			One:   num == 1,
			Few:   num > 1,
			Two:   num == 2,
			Many:  num > 1,
			Other: num > 1,
		},
	}
}
