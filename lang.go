package timeago

var (
	language = "ru"
	location = "Europe/Kiev"
)

func Set(paramName string, value string) {
	switch paramName {
	case "language":
		language = value
		break
	case "location":
		location = value
	}
}

func trans(key string) string {
	if language == "ru" {
		return getRussian()[key]
	} else {
		return getEnglish()[key]
	}
}
