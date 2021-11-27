package timeago

import "testing"

func TestTrans(t *testing.T) {
	cases := []struct {
		lang   string
		trans  func() string
		result string
	}{
		{"ru", func() string { return trans().Online }, "В сети"},
		{"ru", func() string { return trans().Second }, "секунда"},
		{"ru", func() string { return trans().Hour }, "час"},
		{"ru", func() string { return trans().Day }, "день"},
		{"en", func() string { return trans().Online }, "Online"},
		{"en", func() string { return trans().Second }, "second"},
		{"en", func() string { return trans().Hour }, "hour"},
		{"en", func() string { return trans().Day }, "day"},
	}

	for _, tc := range cases {
		t.Run("returns "+tc.lang+" language", func(test *testing.T) {
			Set("language", tc.lang)

			if result := tc.trans(); result != tc.result {
				test.Errorf("Result must be %s but got %s", tc.result, result)
			}
		})
	}
}

func TestSet_for_language(t *testing.T) {
	cases := []struct {
		name  string
		value string
		err   string
	}{
		{"sets language to ru", "ru", "Set must set the `language` variable to `ru` but it didn't"},
		{"sets language to en", "en", "Set must set the `language` variable to `en` but it didn't"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(test *testing.T) {
			Set("language", tc.value)

			if language != tc.value {
				test.Error(tc.err)
			}
		})
	}
}

func TestSet_for_location(t *testing.T) {
	cases := []struct {
		name  string
		value string
		err   string
	}{
		{"sets location to India Delhi", "India/Delhi", "Set must set the `location` variable to `India/Delhi` but it didn't"},
		{"sets language to Europe/Kiev", "Europe/Kiev", "Set must set the `location` variable to `Europe/Kiev` but it didn't"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(test *testing.T) {
			Set("location", tc.value)

			if location != tc.value {
				test.Error(tc.err)
			}
		})
	}
}
