package timeago

const (
	LangEn string = "en" // English
	LangRu string = "ru" // Russian
	LangUk string = "uk" // Ukrainian
	LangNl string = "nl" // Dutch
	LangDe string = "de" // German
	LangZh string = "zh" // Chinese
	LangBe string = "be" // Belarusian
	LangEs string = "es" // Spanish
	LangJa string = "ja" // Japanese
	LangFr string = "fr" // French
)

type Config struct {
	// Language is an ISO 639 language code like en, ru, de, nl, etc.
	// Default: LangEn ("en").
	Language string

	// Location is the timezone location needed for parsing
	// string date like "2019-01-01 00:00:00" to time.Time.
	// If Location is not set it will default to the server's.
	// Example: "America/New_York", "Europe/Moscow".
	// Default: "UTC".
	Location string

	// Translations is a slice of language sets that can be used to
	// overwrite the default translations. Read more about it in the docs
	// https://time-ago.github.io/configurations.html#overwrite-translations
	// Default: []LangSet{}
	Translations []LangSet

	// OnlineThreshold is the threshold in seconds to determine when
	// Timeago should show "Online" instead of "X seconds ago". If the
	// time difference is less than the threshold, it will show "Online".
	// Minimum value: "1"
	// Default: "60"
	OnlineThreshold int

	// JustNowThreshold is the threshold in seconds to determine when
	// Timeago should show "Just now" instead of "X seconds ago". If the
	// time difference is less than the threshold, it will show "Just now".
	// Minimum value: "1"
	// Default: "60"
	JustNowThreshold int
}

// NewConfig creates a new Config instance with the given language, location
// and language sets to overwrite the default translations.
func NewConfig(lang, loc string, langSets []LangSet, online, justNow int) *Config {
	return &Config{
		Language:         lang,
		Location:         loc,
		Translations:     langSets,
		OnlineThreshold:  online,
		JustNowThreshold: justNow,
	}
}

func (c Config) isLocationProvided() bool {
	return c.Location != "UTC" && c.Location != ""
}
