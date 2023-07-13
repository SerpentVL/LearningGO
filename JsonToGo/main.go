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
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalf("Нифига не декодируется %s", err)
	}
	return result
}

func main() {

	array.LoadJsonToBugs("clang-bugs.json")

	fmt.Println("------------")
	array.PrintBugs()
	fmt.Println("------------")

}
