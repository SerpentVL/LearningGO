package array

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Bug struct {
	Priority       string
	Type           string
	Category       string
	CWE            int
	TranslatedType string
}

func FindCWE(bug_type string, bug_category string) *Bug {
	// Simple cases
	for _, bug := range bugs {
		if strings.EqualFold(bug_type, bug.Type) && strings.EqualFold(bug_category, bug.Category) {
			return &bug
		}
	}
	// TODO: Special cases
	return nil
}

// -- заполнение bugs изjson-файла --
// TODO: запись в лог - отследить --
func LoadJsonToBugs(filename string) {

	fileJSON, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Не открывается файл JSON: %s", err)
	}
	defer fileJSON.Close()

	data, err := ioutil.ReadAll(fileJSON)
	err = json.Unmarshal(data, &bugs)
	if err != nil {
		log.Fatalf("Нифига не декодируется %s", err)
	}
}

func PrintBugs() {
	fmt.Println(len(bugs))

	for _, arg := range bugs {
		fmt.Printf("{%s %s %s %d %s},\n", arg.Priority, arg.Type, arg.Category, arg.CWE, arg.TranslatedType)
	}
}
