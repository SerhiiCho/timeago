package timeago

type Translation struct {
	Language     string
	Translations map[string]string
}

type Config struct {
	Language     string
	Location     string
	Translations []Translation
}

var config = Config{
	Language:     "en",
	Location:     "",
	Translations: []Translation{},
}

func SetConfig(conf Config) {
	if conf.Language == "" {
		conf.Language = config.Language
	}

	config = conf
}

func locationIsSet() bool {
	return config.Location != ""
}

func locationIsNotSet() bool {
	return !locationIsSet()
}

func translationsAreSet() bool {
	return len(config.Translations) > 0
}
