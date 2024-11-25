package timeago

import "testing"

func TestOptionIsEnabled(t *testing.T) {
	t.Run("returns true if option is enabled", func(test *testing.T) {
		options = []opt{OptNoSuffix}

		if res := optionIsEnabled(OptNoSuffix); res == false {
			test.Error("Result must be true, but got false instead")
		}

		options = []opt{}
	})

	t.Run("returns true if option is enabled with other option", func(test *testing.T) {
		options = []opt{OptNoSuffix, OptUpcoming}

		if res := optionIsEnabled(OptUpcoming); res == false {
			test.Error("Result must be true, but got false instead")
		}

		options = []opt{}
	})

	t.Run("returns false if option is disabled", func(test *testing.T) {
		if res := optionIsEnabled(OptNoSuffix); res == true {
			test.Error("Result must be true, but got false instead")
		}
	})
}
