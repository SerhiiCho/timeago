package timeago

import (
	"strings"

	"github.com/SerhiiCho/timeago/v3/internal/utils"
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
			Few:  end == 2 || end == 3 || end == 4,
			Many: (num >= 5 && num <= 20) || end == 0 || (end >= 5 && end <= 9),
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

	return nil, utils.Errorf("Language '" + lang + "' not found")
}
