package timeago

type Config struct {
	Language string
	Location string
}

var config = Config{
	Language: "en",
	Location: "America/New_York",
}

func SetConfig(conf Config) {
	if conf.Language == "" {
		conf.Language = config.Language
	}

	if conf.Location == "" {
		conf.Location = config.Location
	}

	config = conf
}
