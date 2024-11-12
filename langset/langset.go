package langset

import "github.com/SerhiiCho/timeago/v3/ctx"

type CLDR struct {
	One   string `json:"one"`
	Few   string `json:"few"`
	Many  string `json:"many"`
	Other string `json:"other"`
}

type LangSet struct {
	Ago     string `json:"ago"`
	Online  string `json:"online"`
	JustNow string `json:"justnow"`
	Second  CLDR   `json:"second"`
	Minute  CLDR   `json:"minute"`
	Hour    CLDR   `json:"hour"`
	Day     CLDR   `json:"day"`
	Week    CLDR   `json:"week"`
	Month   CLDR   `json:"month"`
	Year    CLDR   `json:"year"`
}

func New(c *ctx.Ctx) LangSet {
	lang := c.Config().Language
	return parseJsonIntoTrans("langs/" + lang + ".json")
}
