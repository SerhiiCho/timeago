package timeago

import (
	"encoding/json"
	"log"

	"github.com/SerhiiCho/timeago/models"
	"github.com/SerhiiCho/timeago/utils"
)

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
