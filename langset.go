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
		return nil, errorf("No called information")
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

		mergeSetValue(trans.Lang, &langSet.Lang)
		mergeSetValue(trans.Format, &langSet.Format)
		mergeSetValue(trans.Ago, &langSet.Ago)
		mergeSetValue(trans.Online, &langSet.Online)
		mergeSetValue(trans.JustNow, &langSet.JustNow)

		mergeLangForms(langSet.Second, trans.Second)
		mergeLangForms(langSet.Minute, trans.Minute)
		mergeLangForms(langSet.Hour, trans.Hour)
		mergeLangForms(langSet.Day, trans.Day)
		mergeLangForms(langSet.Week, trans.Week)
		mergeLangForms(langSet.Month, trans.Month)
		mergeLangForms(langSet.Year, trans.Year)
	}

	return langSet
}

func mergeSetValue(newVal string, currVal *string) {
	if newVal != "" {
		*currVal = newVal
	}
}

func mergeLangForms(curr, langForms LangForms) {
	if len(langForms) == 0 {
		return
	}

	for k, v := range langForms {
		curr[k] = v
	}
}
