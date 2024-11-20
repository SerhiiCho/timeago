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
		name    string
		lang    string
		expect  string
		langSet LangSet
	}{
		{
			name:   "",
			lang:   "en",
			expect: "10 h a",
			langSet: LangSet{
				Lang: "en",
				Ago:  "a",
				Hour: LangForms{
					"one":   "h",
					"other": "h",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Configure(Config{
				Language:     tc.lang,
				Translations: []LangSet{tc.langSet},
			})

			date := getTimestampOfPastDate(time.Hour * 10)

			if res, _ := Parse(date); res != tc.expect {
				t.Errorf("Result must be %v, but got %v instead", tc.expect, res)
			}
		})
	}
}
