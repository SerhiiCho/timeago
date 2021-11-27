package timeago

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/SerhiiCho/timeago/models"
	"github.com/SerhiiCho/timeago/utils"
)

func getLanguage(lang string) models.Lang {
	filePath := fmt.Sprintf("/langs/%s.json", lang)

	thereIsFile, err := utils.FileExists(filePath)

	if !thereIsFile {
		log.Fatalf("File with the path: %s, doesn't exist", filePath)
	}

	if err != nil {
		log.Fatalf("Error while trying to read file %s. Error: %v", filePath, err)
	}

	return parseNeededFile(filePath)
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
