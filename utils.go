package timeago

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"
)

func unixToTime(userDate int) time.Time {
	return time.Unix(int64(userDate), 0)
}

func parseLangSet(fileName string) *LangSet {
	content := readFile(fileName)

	var res LangSet

	err := json.Unmarshal(content, &res)
	if err != nil {
		log.Fatalln(err)
	}

	return &res
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
