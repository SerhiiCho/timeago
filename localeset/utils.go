package localeset

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func parseJsonIntoTrans(fileName string) *LocaleSet {
	content := getFileContent(fileName)

	var res LocaleSet

	err := json.Unmarshal(content, &res)

	if err != nil {
		log.Fatalln(err)
	}

	return &res
}

func getFileContent(filePath string) []byte {
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
