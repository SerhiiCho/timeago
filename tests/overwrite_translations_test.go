package tests

import (
	"testing"
	"time"

	. "github.com/SerhiiCho/timeago"
)

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
