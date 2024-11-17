package config

type Config struct {
	Language string

	// Location is the timezone location neeed for parsing
	// string date like "2019-01-01 00:00:00" to time.Time.
	// If Location is not set it will default to the server's.
	// Example: "America/New_York", "Asia/China"
	Location string

	Translations []Translation
}

func New(lang, loc string, trans []Translation) *Config {
	return &Config{
		Language:     lang,
		Location:     loc,
		Translations: trans,
	}
}

func (c Config) LocationIsSet() bool {
	return c.Location != ""
}
