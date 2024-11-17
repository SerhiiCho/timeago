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
			c := NewConfig("ru", tc.loc, []Translation{})
			actual := c.LocationIsSet()

			if actual != tc.expect {
				t.Fatalf("Expected %v, but got %v", tc.expect, actual)
			}
		})
	}
}

func TestCustomTranslations(t *testing.T) {
	t.Skip()

	t.Run("can overwrite translations", func(t *testing.T) {
		Configure(&Config{
			Language: "en",
			Translations: []Translation{
				{
					Language: "en",
					Translations: &LangSet{
						Hour: LangForms{
							"other": "h",
						},
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
