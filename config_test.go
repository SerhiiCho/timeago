package timeago

import (
	"testing"
	"time"
)

func TestTrans(t *testing.T) {
	cases := []struct {
		lang   string
		trans  func() string
		result string
	}{
		{"ru", func() string { return trans().Online }, "В сети"},
		{"ru", func() string { return trans().Second }, "секунда"},
		{"ru", func() string { return trans().Hour }, "час"},
		{"ru", func() string { return trans().Day }, "день"},
		{"en", func() string { return trans().Online }, "Online"},
		{"en", func() string { return trans().Second }, "second"},
		{"en", func() string { return trans().Hour }, "hour"},
		{"en", func() string { return trans().Day }, "day"},
	}

	for _, tc := range cases {
		t.Run("returns "+tc.lang+" language", func(test *testing.T) {
			config := Config{
				Language: tc.lang,
			}

			SetConfig(config)

			if result := tc.trans(); result != tc.result {
				test.Errorf("Result must be %s but got %s", tc.result, result)
			}
		})
	}
}

func TestSetConfigLanguage(t *testing.T) {
	cases := []struct {
		name  string
		value string
		err   string
	}{
		{"sets language to ru", "ru", "Set must set language to 'ru' but it didn't"},
		{"sets language to en", "en", "Set must set language to 'en' but it didn't"},
		{"sets language to nl", "nl", "Set must set language to 'nl' but it didn't"},
		{"sets language to uk", "uk", "Set must set language to 'uk' but it didn't"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(test *testing.T) {
			config := Config{
				Language: tc.value,
			}

			SetConfig(config)

			if config.Language != tc.value {
				test.Error(tc.err)
			}
		})
	}
}

func TestSetConfigLocation(t *testing.T) {
	cases := []struct {
		name  string
		value string
		err   string
	}{
		{"sets location to India Delhi", "India/Delhi", "Set must set the location to 'India/Delhi' but it didn't"},
		{"sets language to Europe/Kiev", "Europe/Kiev", "Set must set the location to 'Europe/Kiev' but it didn't"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(test *testing.T) {
			config := Config{
				Location: tc.value,
			}

			SetConfig(config)

			if config.Location != tc.value {
				test.Error(tc.err)
			}
		})
	}
}

func TestSetConfigTranslations(t *testing.T) {
	cases := []struct {
		name         string
		time         time.Time
		expect       string
		translations []Translation
	}{
		{
			name:   "changes 'minutes' to 'm'",
			time:   time.Now().Add(-time.Minute * 2),
			expect: "2 m ago",
			translations: []Translation{
				{
					Language: "en",
					Translations: map[string]string{
						"minutes": "m",
					},
				},
			},
		},
		{
			name:   "changes 'minute' to 'm'",
			time:   time.Now().Add(-time.Minute),
			expect: "1 m ago",
			translations: []Translation{
				{
					Language: "en",
					Translations: map[string]string{
						"minute": "m",
					},
				},
			},
		},
		{
			name:   "not changes 'minute' to 'm' when output should be '4 minutes ago'",
			time:   time.Now().Add(-time.Minute * 4),
			expect: "4 minutes ago",
			translations: []Translation{
				{
					Language: "en",
					Translations: map[string]string{
						"minute": "m",
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(test *testing.T) {
			config := Config{
				Translations: tc.translations,
			}

			SetConfig(config)

			result := Parse(tc.time)

			if result != tc.expect {
				test.Errorf("Expected '%s', but got '%s'", tc.expect, result)
			}
		})
	}
}
