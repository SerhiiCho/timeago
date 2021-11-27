package timeago

import (
	"io/ioutil"
	"log"
	"strings"
)

func GetLanguage(lang string, key string) string {
	// allLangs := getAllLanguages()

	return ""
	// if lang == "ru" {
	// 	return ru[key]
	// }

	// return en[key]
}

func parseNeededFile() {
	// result := make(map[string]map[string]string)

	// for _, fileName := range fileNames {
	// 	var langItem models.Lang

	// read file

	// json.NewDecoder().Decode(&langItem)
	// newItem := map[string]string{
	// 	"Title": "",
	// 	"Body":  "",
	// }

	// result[lang] = append(result[lang], newItem)

	// fmt.Println(name)
	// }
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
