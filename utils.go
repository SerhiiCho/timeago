package timeago

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func unixToTime(userDate int) time.Time {
	return time.Unix(int64(userDate), 0)
}

func parseLangSet(filePath string, fs embed.FS) (*LangSet, error) {
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

func readFile(filePath string) []byte {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	return content
}

func isFilePresent(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, err
}
