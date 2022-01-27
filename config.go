package timeago

type Config struct {
	Language string
	Location string
}

var config = Config{
	Language: "en",
	Location: "",
}

func SetConfig(conf Config) {
	if conf.Language == "" {
		conf.Language = config.Language
	}

	config = conf
}
