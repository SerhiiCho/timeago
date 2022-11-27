package timeago

import "strings"

func overwriteOutput(result string) string {
	if !translationsAreSet() {
		return result + " " + trans().Ago
	}

	suffix := trans().Ago

	for _, translation := range config.Translations {
		if translation.Language != config.Language {
			continue
		}

		for key, trans := range translation.Translations {
			parts := strings.Split(result, " ")

			if key == "ago" {
				suffix = trans
			}

			if parts[1] == key {
				result = strings.Replace(result, key, trans, 1)
			}
		}
	}

	if suffix == "" {
		return result
	}

	return result + " " + suffix
}
