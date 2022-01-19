package timeago

import . "github.com/SerhiiCho/timeago/models"

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
