package timeago

import (
	"fmt"
	"path"
	"runtime"
)

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

func NewLangSet() *LangSet {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		panic("[Timeago]: No caller information")
	}

	rootPath := path.Dir(filename)
	filePath := fmt.Sprintf(rootPath+"/langs/%s.json", conf.Language)

	if cache, hasCache := cachedJsonRes[filePath]; hasCache {
		return cache
	}

	return parseLangSet(filePath)
}
