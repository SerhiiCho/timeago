package tests

import (
	"testing"
	"time"

	ago "github.com/SerhiiCho/timeago/v3"
)

func TestSlavicTeensAboveHundredUseManyForm(t *testing.T) {
	cases := []struct {
		name  string
		lang  string
		years int
		want  string
	}{
		{"ru 111 years", ago.LangRu, 111, "111 лет назад"},
		{"ru 112 years", ago.LangRu, 112, "112 лет назад"},
		{"ru 114 years", ago.LangRu, 114, "114 лет назад"},
		{"uk 111 years", ago.LangUk, 111, "111 років тому"},
		{"uk 112 years", ago.LangUk, 112, "112 років тому"},
		{"uk 114 years", ago.LangUk, 114, "114 років тому"},
		{"be 111 years", ago.LangBe, 111, "111 гадоў таму"},
		{"be 112 years", ago.LangBe, 112, "112 гадоў таму"},
		{"be 114 years", ago.LangBe, 114, "114 гадоў таму"},
		{"ru 121 years keeps one", ago.LangRu, 121, "121 год назад"},
		{"ru 122 years keeps few", ago.LangRu, 122, "122 года назад"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ago.Reconfigure(ago.Config{Language: tc.lang})
			got, err := ago.Parse(subYears(time.Duration(tc.years)))
			assertResult(t, tc.want, got, err)
		})
	}
}
