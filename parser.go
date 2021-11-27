package timeago

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/SerhiiCho/timeago/models"
	"github.com/SerhiiCho/timeago/utils"
)

func GetLanguage(lang string, key string) string {
	// allLangs := getAllLanguages()

	return ""
	// if lang == "ru" {
	// 	return ru[key]
	// }

	// return en[key]
}

// Parses json file and unmarshals result into a struct
func parseNeededFile(fileName string) models.Lang {
	content := utils.GetFileContent(fileName)

	var result models.Lang

	err := json.Unmarshal(content, &result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

// Scans 'langs' directory and returns names
// of all the files inside without extension
func getAllLanguages() []string {
	files, err := ioutil.ReadDir("./langs")
	var result []string

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		lang := strings.Split(file.Name(), ".")[0]
		result = append(result, lang)
	}

	return result
}
