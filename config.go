package timeago

// Config is the main configurations for the package
type Config struct {
	// Default value is "ru"
	Language string
	// Default value is "Europe/Kiev"
	Location string
}

var config = Config{
	Language: "ru",
	Location: "Europe/Kiev",
}

// SetConfig sets configurations
func SetConfig(conf Config) {
	if conf.Language == "" {
		conf.Language = config.Language
	}

	if conf.Location == "" {
		conf.Location = config.Location
	}

	config = conf
}
