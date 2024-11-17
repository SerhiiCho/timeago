package timeago

import (
	"strings"
)

type Rule struct {
	Zero  bool
	One   bool
	Two   bool
	Few   bool
	Many  bool
	Other bool
}

var grammarRules = func(num int) map[string]*Rule {
	end := num % 10

	return map[string]*Rule{
		"en,nl,de": {
			Zero: num == 0,
			One:  num == 1,
			Two:  num == 2,
			Few:  num > 1,
			Many: num > 1,
		},
		"ru,uk": {
			Zero: num == 0,
			One:  num == 1 || (num > 20 && end == 1),
			Two:  num == 2,
			Few:  end == 2 || end == 3 || end == 4,
			Many: (num >= 5 && num <= 20) || end == 0 || (end >= 5 && end <= 9),
		},
	}
}

func identifyGrammarRules(num int, lang string) (*Rule, error) {
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
