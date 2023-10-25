package main

import "fmt"

var data []int

func main() {
	data = make([]int, 4, 7) // элементы нулевые
	fmt.Print(data)          // создается 4 элемента = 0
	data = append(data, 1)   // работает -- добавляется 5 элемент
	fmt.Println(data)

	t := make([]byte, 0, 3)
	t = append(t, 100)
	fmt.Println(t)

	/*	// зацикливание, так как все время увеличивается размер среза
		for i := 0; i < len(t); i++ {
			t = append(t, 1)
		}
		fmt.Println(t)
	*/
	// это работает -- добавляются элементы в срез и не зацикливается --
	length := 4
	countP := 3
	for i := 0; i < length; i++ {
		fmt.Println(i, length, countP)
		t = append(t, byte(i))
		if countP > 0 {
			length += countP
			countP--
		}
	}
	fmt.Println(t)
}
