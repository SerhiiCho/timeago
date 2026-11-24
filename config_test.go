package timeago

import (
	"testing"
)

func TestIsLocationProvided(t *testing.T) {
	cases := []struct {
		name string
		loc  string
		want bool
	}{
		{
			name: "Location is set to Europe/Paris",
			loc:  "Europe/Paris",
			want: true,
		},
		{
			name: "Location is set to Asia/Shanghai",
			loc:  "Asia/Shanghai",
			want: true,
		},
		{
			name: "Location is not set",
			loc:  "UTC",
			want: false,
		},
		{
			name: "Location is empty",
			loc:  "",
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			conf := NewConfig("en", tc.loc, []LangSet{}, 60, 60)
			actual := conf.isLocationProvided()

			if actual != tc.want {
				t.Fatalf("want %v, got %v instead", tc.want, actual)
			}
		})
	}
}
