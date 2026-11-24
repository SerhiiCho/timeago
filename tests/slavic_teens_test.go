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
		{
			name:  "ru 111 years",
			lang:  ago.LangRu,
			years: 111,
			want:  "111 лет назад",
		},
		{
			name:  "ru 112 years",
			lang:  ago.LangRu,
			years: 112,
			want:  "112 лет назад",
		},
		{
			name:  "ru 114 years",
			lang:  ago.LangRu,
			years: 114,
			want:  "114 лет назад",
		},
		{
			name:  "uk 111 years",
			lang:  ago.LangUk,
			years: 111,
			want:  "111 років тому",
		},
		{
			name:  "uk 112 years",
			lang:  ago.LangUk,
			years: 112,
			want:  "112 років тому",
		},
		{
			name:  "uk 114 years",
			lang:  ago.LangUk,
			years: 114,
			want:  "114 років тому",
		},
		{
			name:  "be 111 years",
			lang:  ago.LangBe,
			years: 111,
			want:  "111 гадоў таму",
		},
		{
			name:  "be 112 years",
			lang:  ago.LangBe,
			years: 112,
			want:  "112 гадоў таму",
		},
		{
			name:  "be 114 years",
			lang:  ago.LangBe,
			years: 114,
			want:  "114 гадоў таму",
		},
		{
			name:  "ru 121 years keeps one",
			lang:  ago.LangRu,
			years: 121,
			want:  "121 год назад",
		},
		{
			name:  "ru 122 years keeps few",
			lang:  ago.LangRu,
			years: 122,
			want:  "122 года назад",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ago.Reconfigure(ago.Config{Language: tc.lang})
			got, err := ago.Parse(subYears(time.Duration(tc.years)))
			assertResult(t, tc.want, got, err)
		})
	}
}
