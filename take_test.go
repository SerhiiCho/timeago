package timeago

import "testing"

func TestGetOption(test *testing.T) {
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

func TestGetTimeTranslations(t *testing.T) {
	t.Parallel()

	t.Run("returns 7 items", func(test *testing.T) {
		if res := len(getTimeTranslations()); res != 7 {
			test.Errorf("Must return 7 items but got %d", res)
		}
	})

	t.Run("every slice has 3 items", func(test *testing.T) {
		for _, slice := range getTimeTranslations() {
			if res := len(slice); res != 3 {
				test.Errorf("Slice Must return 3 items but got %d", res)
			}
		}
	})
}
