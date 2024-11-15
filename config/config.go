package config

type Config struct {
	Language     string
	Location     string
	Translations []Translation
}

func New(lang, loc string, trans []Translation) Config {
	return Config{
		Language:     lang,
		Location:     loc,
		Translations: trans,
	}
}

func (c Config) LocationIsSet() bool {
	return c.Location != ""
}
