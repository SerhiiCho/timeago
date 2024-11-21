package timeago

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func unixToTime(timestamp int) time.Time {
	return time.Unix(int64(timestamp), 0)
}

func timestampFromPastDate(subDuration time.Duration) int {
	return int(time.Now().Add(-subDuration).UnixNano() / 1000000000)
}

func errorf(msg string, a ...interface{}) error {
	return fmt.Errorf("[Timeago]: "+msg, a...)
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
