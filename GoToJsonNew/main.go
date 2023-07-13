package main

import (
	"LearningGO/GoToJsonNew/array"

	"encoding/json"
	"fmt"
	"log"
	"os"
)

// -- вывод jSON в файл --
func saveToJSON(filename string, data []byte) {
	fileJSON, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Не создался файл JSON: %s", err)
	}
	defer fileJSON.Close()
	fmt.Fprintf(fileJSON, string(data))
}

func main() {
	fmt.Println(len(array.Bugs))
	data, err := json.MarshalIndent(array.Bugs, "", " ") // -- преобразование слайса в JSON в красивом виде --
	if err != nil {
		log.Fatalf("Сбой маршаллинга JSON: %s", err)
	}
	//fmt.Printf("%s\n", data) // -- вместо этого - запись в файл --

	saveToJSON("new-clang-bugs.json", data) // -- записали в файл --

}
