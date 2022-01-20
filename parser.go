package timeago

import (
	"encoding/json"
	"log"
)

// Parses json file and includes result into a Lang type
func parseNeededFile(fileName string) Lang {
	content := getFileContent(fileName)

	var result Lang

	err := json.Unmarshal(content, &result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}
