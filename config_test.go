package timeago

import (
	"testing"
	"time"
)

func TestLocationIsSet(t *testing.T) {
	cases := []struct {
		name   string
		loc    string
		expect bool
	}{
		{
			name:   "Location is set",
			loc:    "Russia/Moscow",
			expect: true,
		},
		{
			name:   "Location is not set",
			loc:    "",
			expect: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewConfig("ru", tc.loc, []LangSet{})
			actual := c.isLocationProvided()

			if actual != tc.expect {
				t.Fatalf("Expected %v, but got %v", tc.expect, actual)
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

			date := getTimestampOfPastDate(tc.time)

			if res, _ := Parse(date); res != tc.expect {
				t.Errorf("Result must be %q, but got %q instead", tc.expect, res)
			}
		})
	}
}
