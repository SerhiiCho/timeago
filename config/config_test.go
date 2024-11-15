package config

import "testing"

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
			c := New("ru", tc.loc, []Translation{})
			actual := c.LocationIsSet()

			if actual != tc.expect {
				t.Fatalf("Expected %v, but got %v", tc.expect, actual)
			}
		})
	}
}
