package main

import (
	"LearningGO/JsonToGo/array"
	"LearningGO/JsonToGo/clang_log"
	"fmt"
)

func main() {

	err := array.LoadJsonToBugs("clang-bugs.json")
	if err != nil {
		clang_log.Info("Ошибка при чтении json-файла: %s ", err)
	} else {
		fmt.Println("------------")
		array.PrintBugs()
		fmt.Println("------------")
	}
}
