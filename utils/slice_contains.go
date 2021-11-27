package utils

func SliceContains(slice []string, key string) bool {
	for _, a := range slice {
		if a == key {
			return true
		}
	}

	return false
}
