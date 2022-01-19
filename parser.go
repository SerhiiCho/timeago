package timeago

import (
	"encoding/json"
	"log"

	"github.com/SerhiiCho/timeago/utils"
)

// Parses json file and unmarshals result into a struct
func parseNeededFile(fileName string) Lang {
	content := utils.GetFileContent(fileName)

	var result Lang

	err := json.Unmarshal(content, &result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
