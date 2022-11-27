package timeago

import (
	"testing"
	"time"
)

func TestSetConfigTranslations(t *testing.T) {
	cases := []struct {
		lang         string
		name         string
		time         time.Time
		expect       string
		translations []Translation
	}{
		{
			lang:   "en",
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
			lang:   "en",
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
			lang:   "en",
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
		{
			lang:   "en",
			name:   "you can remove 'ago' suffix by overwriting it to an empty string",
			time:   time.Now().Add(-time.Minute * 2),
			expect: "2 minutes",
			translations: []Translation{
				{
					Language: "en",
					Translations: map[string]string{
						"ago": "",
					},
				},
			},
		},
		{
			lang:   "en",
			name:   "you can overwrite 'ago' suffix with different value",
			time:   time.Now().Add(-time.Minute * 2),
			expect: "2 minutes passed",
			translations: []Translation{
				{
					Language: "en",
					Translations: map[string]string{
						"ago": "passed",
					},
				},
			},
		},
		{
			lang:   "ru",
			name:   "you can overwrite 'минуты' with new value",
			time:   time.Now().Add(-time.Minute * 2),
			expect: "2 мин назад",
			translations: []Translation{
				{
					Language: "ru",
					Translations: map[string]string{
						"минуты": "мин",
					},
				},
			},
		},
		{
			lang:   "ru",
			name:   "you can remove 'назад' suffix by overwriting it to an empty string",
			time:   time.Now().Add(-time.Minute * 2),
			expect: "2 минуты",
			translations: []Translation{
				{
					Language: "ru",
					Translations: map[string]string{
						"назад": "",
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(test *testing.T) {
			config := Config{
				Language:     tc.lang,
				Translations: tc.translations,
			}

			SetConfig(config)

			result := Parse(tc.time)

			if result != tc.expect {
				test.Errorf("Expected '%s', but got '%s'", tc.expect, result)
			}
		})
	}

	t.Run("english does not change if you overwrite russian", func(test *testing.T) {
		config := Config{
			Language: "en",
			Translations: []Translation{
				{
					Language: "ru",
					Translations: map[string]string{
						"назад":  "xx",
						"минут":  "xx",
						"минуты": "xx",
					},
				},
			},
		}

		SetConfig(config)

		result := Parse(time.Now().Add(-time.Minute * 2))

		if result != "2 minutes ago" {
			test.Errorf("Expected '2 minutes ago', but got '%s'", result)
		}
	})
}
