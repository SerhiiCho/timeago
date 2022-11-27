package timeago

import "strings"

// overwriteOutput takes the final result, for example "4 days",
// and modifies the output if user added custom translations
// via config.Translations configuration.
func overwriteOutput(result string) string {
	if !translationsAreSet() {
		return result + " " + trans().Ago
	}

	suffix := trans().Ago

	for _, translation := range config.Translations {
		if translation.Language != config.Language {
			continue
		}

		for key, translation := range translation.Translations {
			parts := strings.Split(result, " ")

			if key == trans().Ago {
				suffix = translation
			}

			if parts[1] == key {
				result = strings.Replace(result, key, translation, 1)
			}
		}
	}

	if suffix == "" {
		return result
	}

	return result + " " + suffix
}
