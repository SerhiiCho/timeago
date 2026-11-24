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

func TestConfigureCustomTranslationsClearsCache(t *testing.T) {
	t.Cleanup(func() {
		ago.Reconfigure(ago.Config{Language: "en"})
	})

	ago.Reconfigure(ago.Config{Language: "en"})

	date := unixFromPastDate(10 * time.Hour)
	if res, err := ago.Parse(date); err != nil || res != "10 hours ago" {
		t.Fatalf("want %q, got %q (error: %v)", "10 hours ago", res, err)
	}

	ago.Configure(ago.Config{
		Translations: []ago.LangSet{{
			Lang: "en",
			Ago:  "a",
			Hour: ago.LangForms{"other": "h"},
		}},
	})

	if res, err := ago.Parse(date); err != nil || res != "10 h a" {
		t.Errorf("want %q, got %q (error: %v)", "10 h a", res, err)
	}
}
