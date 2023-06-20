package main

import (
	"encoding/json"
	"fmt"
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
	saveToJSON(os.Stdout, array.bugs)
}
