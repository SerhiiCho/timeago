package timeago

var (
	translations map[string]string
	lang         string
)

func SetLang(langParam string) {
	lang = langParam
}

// if translation not found returns null
func trans(index string) string {
	return translations[index]
}

// includeTranslations Includes array of translations
// from lang directory nto the translations variable.
func includeTranslations() {
	if lang == "ru" {
		translations = getRussian()
	} else {
		translations = getEnglish()
	}
}
