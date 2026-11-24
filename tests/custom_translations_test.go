package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
)

func TestCustomTranslations(t *testing.T) {
	cases := []struct {
		want    string
		time    time.Duration
		langSet ago.LangSet
	}{
		{
			want: "10 h a",
			time: time.Hour * 10,
			langSet: ago.LangSet{
				Lang: "en",
				Ago:  "a",
				Hour: ago.LangForms{"other": "h"},
			},
		},
		{
			want: "1м н",
			time: time.Minute,
			langSet: ago.LangSet{
				Format: "{num}{timeUnit} {ago}",
				Lang:   "ru",
				Ago:    "н",
				Minute: ago.LangForms{"one": "м"},
			},
		},
		{
			want: "2 d ago",
			time: time.Hour * 24 * 2,
			langSet: ago.LangSet{
				Lang: "en",
				Day:  ago.LangForms{"other": "d"},
			},
		},
	}

	for _, tc := range cases {
		name := tc.langSet.Lang + " " + tc.want

		t.Run(name, func(t *testing.T) {
			ago.Reconfigure(ago.Config{
				Language:     tc.langSet.Lang,
				Translations: []ago.LangSet{tc.langSet},
			})

			date := unixFromPastDate(tc.time)
			if res, _ := ago.Parse(date); res != tc.want {
				t.Errorf("want %q, got %q instead", tc.want, res)
			}
		})
	}
}
