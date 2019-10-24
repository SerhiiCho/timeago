package timeago

var (
	lang     string
	location = "Europe/Kiev"
)

func SetLang(langParam string) {
	lang = langParam
}

func SetLocation(local string) {
	location = local
}

func trans(key string) string {
	if lang == "ru" {
		return getRussian()[key]
	} else {
		return getEnglish()[key]
	}
}
