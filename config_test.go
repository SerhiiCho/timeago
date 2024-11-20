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
	t.Run("can overwrite translations", func(t *testing.T) {
		Configure(Config{
			Language: "en",
			Translations: []LangSet{
				{
					Lang: "en",
					Ago:  "",
					Day: LangForms{
						"one":   "h",
						"other": "h",
					},
					Week: LangForms{
						"one":   "w",
						"other": "w",
					},
				},
			},
		})

		date := getTimestampOfPastDate(time.Hour * 10)
		expect := "10 h ago"

		if res, _ := Parse(date); res != expect {
			t.Errorf("Result must be %v, but got %v instead", expect, res)
		}
	})
}
