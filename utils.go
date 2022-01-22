package timeago

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

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

func getFileContent(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return content
}

func getTimestampOfPastDate(subDuration time.Duration) int {
	return int(time.Now().Add(-subDuration).UnixNano() / 1000000000)
}

func parseTimestampToString(timestamp int) string {
	return time.Unix(int64(timestamp), 0).Format("2006-01-02 15:04:05")
}

func getLastNumberDigit(num int) int {
	numStr := strconv.Itoa(num)
	result, _ := strconv.Atoi(numStr[len(numStr)-1:])

	return result
}

func parseJsonIntoLang(fileName string) lang {
	content := getFileContent(fileName)

	var result lang

	err := json.Unmarshal(content, &result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
