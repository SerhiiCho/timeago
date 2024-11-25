package timeago

import (
	"testing"
	"time"
)

func TestIsLocationProvided(t *testing.T) {
	cases := []string{"Europe/Moscow", "UTC", "America/New_York", ""}

	for _, tc := range cases {
		t.Run(tc, func(t *testing.T) {
			c := defaultConfig()

			if !c.isLocationProvided() {
				t.Fatalf("Expected location to be provided, but got %v", tc)
			}
		})
	}
}

func TestCustomTranslations(t *testing.T) {
	cases := []struct {
		expect  string
		time    time.Duration
		langSet LangSet
	}{
		{
			expect: "10 h a",
			time:   time.Hour * 10,
			langSet: LangSet{
				Lang: "en",
				Ago:  "a",
				Hour: LangForms{"other": "h"},
			},
		},
		{
			expect: "1м н",
			time:   time.Minute,
			langSet: LangSet{
				Format: "{num}{timeUnit} {ago}",
				Lang:   "ru",
				Ago:    "н",
				Minute: LangForms{"one": "м"},
			},
		},
		{
			expect: "2 d ago",
			time:   time.Hour * 24 * 2,
			langSet: LangSet{
				Lang: "en",
				Day:  LangForms{"other": "d"},
			},
		},
	}

	for _, tc := range cases {
		name := tc.langSet.Lang + " " + tc.expect

		t.Run(name, func(t *testing.T) {
			Reconfigure(Config{
				Language:     tc.langSet.Lang,
				Translations: []LangSet{tc.langSet},
			})

			date := timestampFromPastDate(tc.time)

			if res, _ := Parse(date); res != tc.expect {
				t.Errorf("Result must be %q, but got %q instead", tc.expect, res)
			}
		})
	}
}
