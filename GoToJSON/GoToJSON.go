package main

import (
	"LearningGO/GoToJSON/array"

	"encoding/json"
	"fmt"
	"log"
	"os"
)

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	data, err := json.MarshalIndent(array.Bugs, "", " ")
	if err != nil {
		log.Fatalf("Сбой маршаллинга JSON: %s", err)
	}
	fmt.Printf("%s\n", data)
	//saveToJSON(os.Stdout, array.Bugs)

}
