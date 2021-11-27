package utils

import (
	"io/ioutil"
	"log"
)

func GetFileContent(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return content
}
