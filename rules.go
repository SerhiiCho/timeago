package timeago

import (
	"strings"
)

type rule struct {
	Zero  bool
	One   bool
	Two   bool
	Few   bool
	Many  bool
	Other bool
}

var grammarRules = func(num int) map[string]*rule {
	lastDigit := num % 10

	return map[string]*rule{
		"en,nl,de": {
			Zero: num == 0,
			One:  num == 1,
			Two:  num == 2,
			Few:  num > 1,
			Many: num > 1,
		},
		"ru,uk": {
			Zero: num == 0,
			One:  num == 1 || (num > 20 && lastDigit == 1),
			Two:  num == 2,
			Few:  lastDigit == 2 || lastDigit == 3 || lastDigit == 4,
			Many: (num >= 5 && num <= 20) || lastDigit == 0 || (lastDigit >= 5 && lastDigit <= 9),
		},
	}
}

func identifyGrammarRules(num int, lang string) (*rule, error) {
	rules := grammarRules(num)

	if v, ok := rules[lang]; ok {
		return v, nil
	}

	for langs, v := range rules {
		if strings.Contains(langs, lang) {
			return v, nil
		}
	}

	return nil, createError("Language '" + lang + "' not found")
}
