package main

import (
	"LearningGO/JsonToGo/array"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// -- чтение JSON-файла --

func loadJSON(filename string) []array.Bug {

	fileJSON, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Не открывается файл JSON: %s", err)
	}
	defer fileJSON.Close()

	data, err := ioutil.ReadAll(fileJSON)

	var result []array.Bug

	jsonErr := json.Unmarshal(data, &result)

	if jsonErr != nil {
		log.Fatalf("Нифига не декодируется %s", err)
	}
	return result
}

func main() {

	data := loadJSON("clang-bugs.json")
	fmt.Println(len(data))

	for _, arg := range data {
		fmt.Printf("{%s %s %s %d %s},\n", arg.Priority, arg.Type, arg.Category, arg.CWE, arg.TranslatedType)

	}

}