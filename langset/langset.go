package langset

import "github.com/SerhiiCho/timeago/v3/ctx"

type LangForms map[string]string

type LangSet struct {
	Format  string    `json:"format"`
	Ago     string    `json:"ago"`
	Online  string    `json:"online"`
	JustNow string    `json:"justnow"`
	Second  LangForms `json:"second"`
	Minute  LangForms `json:"minute"`
	Hour    LangForms `json:"hour"`
	Day     LangForms `json:"day"`
	Week    LangForms `json:"week"`
	Month   LangForms `json:"month"`
	Year    LangForms `json:"year"`
}

func New(c *ctx.Ctx) *LangSet {
	return parseJsonIntoTrans("../langs/" + c.Config().Language + ".json")
}
