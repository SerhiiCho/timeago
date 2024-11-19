package timeago

import (
	"fmt"
	"path"
	"runtime"
)

type LangForms map[string]string

type LangSet struct {
	Lang    string    `json:"lang"`
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

func NewLangSet() (*LangSet, error) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return nil, createError("No called information")
	}

	rootPath := path.Dir(filename)
	filePath := fmt.Sprintf(rootPath+"/langs/%s.json", conf.Language)

	if cache, hasCache := cachedJsonRes[filePath]; hasCache {
		return cache, nil
	}

	return parseLangSet(filePath), nil
}
