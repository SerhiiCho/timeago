package timeago

import "testing"

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
