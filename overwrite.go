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

	for _, translationObj := range config.Translations {
		if translationObj.Language != config.Language {
			continue
		}

		result, suffix = modifyResult(result, translationObj.Translations)
	}

	if suffix == "" {
		return result
	}

	return result + " " + suffix
}

func modifyResult(result string, translations map[string]string) (string, string) {
	suffix := trans().Ago

	for key, translation := range translations {
		parts := strings.Split(result, " ")

		if key == trans().Ago {
			suffix = translation
		}

		if parts[1] == key {
			result = strings.Replace(result, key, translation, 1)
		}
	}

	return result, suffix
}
