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

func newLangSet() (*LangSet, error) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return nil, createError("No called information")
	}

	rootPath := path.Dir(filename)
	filePath := fmt.Sprintf(rootPath+"/langs/%s.json", conf.Language)

	if cache, ok := cachedJsonRes[filePath]; ok {
		return cache, nil
	}

	langSet := parseLangSet(filePath)
	langSet = applyCustomTranslations(langSet)

	cachedJsonRes[filePath] = langSet

	return langSet, nil
}

func applyCustomTranslations(langSet *LangSet) *LangSet {
	if len(conf.Translations) == 0 {
		return langSet
	}

	for _, trans := range conf.Translations {
		if trans.Lang != langSet.Lang {
			continue
		}

		if trans.Format != "" {
			langSet.Format = trans.Format
		}

		if trans.Ago != "" {
			langSet.Ago = trans.Ago
		}

		if trans.Online != "" {
			langSet.Online = trans.Online
		}

		if trans.JustNow != "" {
			langSet.JustNow = trans.JustNow
		}

		if trans.Second != nil {
			for k, v := range trans.Second {
				langSet.Second[k] = v
			}
		}

		if trans.Minute != nil {
			for k, v := range trans.Minute {
				langSet.Minute[k] = v
			}
		}

		if trans.Hour != nil {
			for k, v := range trans.Hour {
				langSet.Hour[k] = v
			}
		}

		if trans.Day != nil {
			for k, v := range trans.Day {
				langSet.Day[k] = v
			}
		}

		if trans.Week != nil {
			for k, v := range trans.Week {
				langSet.Week[k] = v
			}
		}

		if trans.Month != nil {
			for k, v := range trans.Month {
				langSet.Month[k] = v
			}
		}

		if trans.Year != nil {
			for k, v := range trans.Year {
				langSet.Year[k] = v
			}
		}
	}

	return langSet
}
