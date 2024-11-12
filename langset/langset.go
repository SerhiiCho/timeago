package langset

import "github.com/SerhiiCho/timeago/v3/ctx"

type LocaleForms struct {
	One   string `json:"one"`
	Few   string `json:"few"`
	Many  string `json:"many"`
	Other string `json:"other"`
}

type LocaleSet struct {
	Ago     string      `json:"ago"`
	Online  string      `json:"online"`
	JustNow string      `json:"justnow"`
	Second  LocaleForms `json:"second"`
	Minute  LocaleForms `json:"minute"`
	Hour    LocaleForms `json:"hour"`
	Day     LocaleForms `json:"day"`
	Week    LocaleForms `json:"week"`
	Month   LocaleForms `json:"month"`
	Year    LocaleForms `json:"year"`
}

func New(c *ctx.Ctx) *LocaleSet {
	lang := c.Config().Language
	return parseJsonIntoTrans("langs/" + lang + ".json")
}
