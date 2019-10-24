package timeago

import "testing"

func TestGetOption(test *testing.T) {
	test.Parallel()

	test.Run("case has online option", func(t *testing.T) {
		option, hasOption := getOption("2017-02-01 00:00:00|online")

		if !hasOption || option != "online" {
			t.Errorf("Result of getOption func must return `online` string and `true`, but `%s` string and `%v` returned returned", option, hasOption)
		}
	})

	test.Run("case has random option", func(t *testing.T) {
		option, hasOption := getOption("2017-02-01 00:00:00|random")

		if !hasOption || option != "random" {
			t.Errorf("Result of getOption func must return `random` string and `true`, but `%s` string and `%v` returned returned", option, hasOption)
		}
	})

	test.Run("case don't have option", func(t *testing.T) {
		option, hasOption := getOption("2017-02-01 00:00:00")

		if hasOption || option != "" {
			t.Errorf("Result of getOption func must return empty string and `false`, but `%s` string and `%v` returned returned", option, hasOption)
		}
	})
}

func TestTrans(t *testing.T) {
	cases := []struct {
		lang   string
		key    string
		result string
	}{
		{"ru", "online", "В сети"},
		{"ru", "second", "секунда"},
		{"ru", "hour", "час"},
		{"ru", "day", "день"},
		{"en", "online", "Online"},
		{"en", "second", "second"},
		{"en", "hour", "hour"},
		{"en", "day", "day"},
	}

	for _, tc := range cases {
		t.Run("returns "+tc.lang+" language", func(test *testing.T) {
			SetLang(tc.lang)

			if result := trans(tc.key); result != tc.result {
				test.Errorf("Result mast be %s but got %s", tc.result, result)
			}
		})
	}
}
