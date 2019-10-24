package timeago

import "testing"

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

func TestSetLang(t *testing.T) {
	t.Run("sets language to ru", func(test *testing.T) {
		SetLang("ru")

		if lang != "ru" {
			test.Error("SetLang must set the `lang` variable to `ru` but it didn't")
		}
	})

	t.Run("sets language to en", func(test *testing.T) {
		SetLang("en")

		if lang != "en" {
			test.Error("SetLang must set the `lang` variable to `en` but it didn't")
		}
	})
}

func TestSetLocation(t *testing.T) {
	t.Run("sets location to India Delhi", func(test *testing.T) {
		SetLocation("India/Delhi")

		if location != "India/Delhi" {
			test.Error("SetLocation must set the `location` variable to `India/Delhi` but it didn't")
		}
	})

	t.Run("sets language to Asia/China", func(test *testing.T) {
		SetLocation("Europe/Kiev")

		if location != "Europe/Kiev" {
			test.Error("SetLocation must set the `location` variable to `Europe/Kiev` but it didn't")
		}
	})
}
