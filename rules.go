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

// slavicEnding returns the ending that drives the ru, uk and be grammar
// rules. Numbers ending in 11-14 take the many form regardless of their
// last digit (111 лет, 112 лет), so they are treated as a zero ending
// and the one/few last-digit rules do not match.
func slavicEnding(num int) int {
	if lastTwo := num % 100; lastTwo >= 11 && lastTwo <= 14 {
		return 0
	}

	return num % 10
}

var grammarRules = func(num int) map[string]*Rule {
	slavicEnd := slavicEnding(num)

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
			One:  num == 1 || (num > 20 && slavicEnd == 1),
			Two:  num == 2,
			Few:  (slavicEnd == 2 || slavicEnd == 3 || slavicEnd == 4) && (num < 10 || num > 20),
			Many: (num >= 5 && num <= 20) || slavicEnd == 0 || slavicEnd >= 5,
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
