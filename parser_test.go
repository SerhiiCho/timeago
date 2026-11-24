package timeago

import "testing"

func TestParserOptionIsEnabled(t *testing.T) {
	cases := []struct {
		name    string
		options []opt
		want    map[opt]bool
	}{
		{
			name:    "returns true if option is enabled",
			options: []opt{OptNoSuffix},
			want:    map[opt]bool{OptNoSuffix: true},
		},
		{
			name:    "returns true if option is enabled with other option",
			options: []opt{OptNoSuffix, OptUpcoming},
			want:    map[opt]bool{OptUpcoming: true},
		},
		{
			name:    "returns true for the first and the last option",
			options: []opt{OptUpcoming, OptOnline, OptNoSuffix},
			want:    map[opt]bool{OptUpcoming: true, OptNoSuffix: true},
		},
		{
			name:    "returns true for duplicated option",
			options: []opt{OptNoSuffix, OptNoSuffix},
			want:    map[opt]bool{OptNoSuffix: true},
		},
		{
			name:    "returns true for every option when all are enabled",
			options: []opt{OptUpcoming, OptOnline, OptJustNow, OptNoSuffix},
			want: map[opt]bool{
				OptUpcoming: true,
				OptOnline:   true,
				OptJustNow:  true,
				OptNoSuffix: true,
			},
		},
		{
			name:    "returns false if option is disabled",
			options: nil,
			want:    map[opt]bool{OptNoSuffix: false},
		},
		{
			name:    "returns false for empty options slice",
			options: []opt{},
			want:    map[opt]bool{OptNoSuffix: false},
		},
		{
			name:    "returns false if option is not in the list",
			options: []opt{OptNoSuffix, OptUpcoming},
			want:    map[opt]bool{OptOnline: false},
		},
		{
			name:    "returns false for unknown option",
			options: []opt{OptNoSuffix},
			want:    map[opt]bool{opt("unknown"): false},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := &parser{options: tc.options}

			for option, expect := range tc.want {
				if res := p.optionIsEnabled(option); res != expect {
					t.Errorf("want optionIsEnabled(%q) = %v, got %v instead", option, expect, res)
				}
			}
		})
	}
}
