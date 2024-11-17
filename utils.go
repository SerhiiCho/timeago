package timeago

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func parseTimestampIntoTime(timestamp int) time.Time {
	return time.Unix(int64(timestamp), 0)
}

func getTimestampOfPastDate(subDuration time.Duration) int {
	return int(time.Now().Add(-subDuration).UnixNano() / 1000000000)
}

func createError(msg string, a ...interface{}) error {
	return fmt.Errorf("[Timeago]: "+msg, a...)
}

func parseLangSet(fileName string) *LangSet {
	content := fileContent(fileName)

	var res LangSet

	err := json.Unmarshal(content, &res)

	if err != nil {
		log.Fatalln(err)
	}

	return &res
}

func fileContent(filePath string) []byte {
	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatalln(err)
	}

	return content
}

func fileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)

	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, err
}
