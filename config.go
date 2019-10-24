package timeago

var (
	language = "ru"
	location = "America/New_York"
)

// Set sets configurations parameters to given value
// paramName is the name of the parameter, can be
// `language` or `location`.
// Seconds parameter is value of the configuration.
// For parameter `language` can be `ru` or `en`
func Set(paramName string, value string) {
	switch paramName {
	case "language":
		language = value
		break
	case "location":
		location = value
	}
}
