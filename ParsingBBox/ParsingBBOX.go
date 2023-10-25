package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	s := "1.1,2.2,3.3,4.4"
	// -- обязательная замена запятой на пробел --
	coord := strings.Fields(strings.ReplaceAll(s, ",", " "))
	fmt.Println(len(coord))
	fmt.Println(coord[0])
	fmt.Println(coord[3])
	var d float64
	var err error
	d, err = strconv.ParseFloat(coord[0], 64)
	if err != nil {
		// -- выводим сообщение об ошибке --
	}
	fmt.Println(d)
}
