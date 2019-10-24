package timeago

var (
	lang string
)

func SetLang(langParam string) {
	lang = langParam
}

func trans(key string) string {
	if lang == "ru" {
		return getRussian()[key]
	} else {
		return getEnglish()[key]
	}
}
