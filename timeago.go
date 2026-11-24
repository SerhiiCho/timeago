package timeago

import (
	"sync"
)

var (
	// confMu protects conf from concurrent reads and writes.
	confMu sync.RWMutex

	// conf is configuration provided by the user.
	conf = defaultConfig()

	// cacheMu protects cachedJsonRes from concurrent reads and writes.
	cacheMu sync.RWMutex

	// cachedJsonRes saves parsed JSON translations to prevent
	// parsing the same JSON file multiple times.
	cachedJsonRes = map[string]*LangSet{}
)

// Parse coverts provided datetime into `x time ago` format.
// The first argument can different types of inputs:
// 1. Unix timestamp: int, int32, int64, unit, uint32, uint64, string
// 2. Type from Go time package: time.Time
// 3. Datetime string in format 'YYYY-MM-DD HH:MM:SS': string
func Parse(date any, opts ...opt) (string, error) {
	p := &parser{
		conf:    getConfig(),
		options: append([]opt{}, opts...),
	}

	t, err := p.dateToTime(date)
	if err != nil {
		return "", err
	}

	return p.computeTimeSince(t)
}

// Configure applies the given configuration to the timeago without
// overriding the previous configuration. It will only override the
// provided configuration. If you want to override the previous
// configurations, use Reconfigure function instead.
func Configure(c Config) {
	confMu.Lock()
	defer confMu.Unlock()

	if c.OnlineThreshold > 0 {
		conf.OnlineThreshold = c.OnlineThreshold
	}

	if c.JustNowThreshold > 0 {
		conf.JustNowThreshold = c.JustNowThreshold
	}

	if c.Language != "" {
		conf.Language = c.Language
	}

	if c.Location != "" {
		conf.Location = c.Location
	}

	if len(c.Translations) > 0 {
		conf.Translations = c.Translations
		clearCache()
	}
}

// Reconfigure reconfigures the timeago with the provided configuration.
// It will override the previous configuration with the new one.
func Reconfigure(c Config) {
	confMu.Lock()
	conf = defaultConfig()
	confMu.Unlock()

	clearCache()
	Configure(c)
}

func clearCache() {
	cacheMu.Lock()
	cachedJsonRes = map[string]*LangSet{}
	cacheMu.Unlock()
}

func getConfig() Config {
	confMu.RLock()
	defer confMu.RUnlock()

	return *conf
}

func defaultConfig() *Config {
	return NewConfig("en", "UTC", []LangSet{}, 60, 60)
}
