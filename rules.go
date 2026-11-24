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

	// In ru, uk and be, numbers ending in 11-14 take the many form
	// regardless of their last digit (111 лет, 112 лет). Treat them
	// as a zero ending so the one/few last-digit rules do not match.
	if lastTwo := num % 100; lastTwo >= 11 && lastTwo <= 14 {
		end = 0
	}

	return map[string]*Rule{
		"en,nl,de,es,fr": {
			Zero: num == 0,
			One:  num == 1,
			Two:  num == 2,
			Few:  num > 1,
			Many: num > 1,
		},
		"ru,uk,be": {
			Zero: num == 0,
			One:  num == 1 || (num > 20 && end == 1),
			Two:  num == 2,
			Few:  (end == 2 || end == 3 || end == 4) && (num < 10 || num > 20),
			Many: (num >= 5 && num <= 20) || end == 0 || end >= 5,
		},
		"zh,ja": {
			Zero: true,
			One:  true,
			Two:  true,
			Few:  true,
			Many: true,
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

	return nil, errorf("Language '%s' not found", lang)
}
