package timeago

import "testing"

func TestOptionIsEnabled(t *testing.T) {
	t.Run("returns true if option is enabled", func(test *testing.T) {
		options = []Option{"noSuffix"}

		if res := optionIsEnabled("noSuffix"); res == false {
			test.Error("Result must be true, but got false instead")
		}

		options = []Option{}
	})

	t.Run("returns true if option is enabled with other option", func(test *testing.T) {
		options = []Option{"noSuffix", "upcoming"}

		if res := optionIsEnabled("upcoming"); res == false {
			test.Error("Result must be true, but got false instead")
		}

		options = []Option{}
	})

	t.Run("returns false if option is disabled", func(test *testing.T) {
		if res := optionIsEnabled("noSuffix"); res == true {
			test.Error("Result must be true, but got false instead")
		}
	})
}
