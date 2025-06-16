package timeago

import (
	"encoding/json"
	"fmt"
	"time"
)

func unixToTime(userDate int) time.Time {
	return time.Unix(int64(userDate), 0)
}

func parseLangSet(filePath string) (*LangSet, error) {
	content, err := langsFS.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read language file: %w", err)
	}

	langSet := &LangSet{}

	if err := json.Unmarshal(content, &langSet); err != nil {
		return nil, fmt.Errorf("could not unmarshal language file: %w", err)
	}

	return langSet, nil
}

func isUnsignedInteger(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}
