package langset

import "github.com/SerhiiCho/timeago/v3/ctx"

type CLDR struct {
	One   string
	Few   string
	Many  string
	Other string
}

type LangSet struct {
	Ago     string
	Online  string
	JustNow string
	Second  CLDR
	Minute  CLDR
	Hour    CLDR
	Day     CLDR
	Week    CLDR
	Month   CLDR
	Year    CLDR
}

func New(c *ctx.Ctx) LangSet {
	lang := c.Config().Language
	return parseJsonIntoTrans("langset/" + lang + ".json")
}
