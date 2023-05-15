package main

import (
	"fmt"
	//"os"
)

func main() {

	var number int = 30
	fmt.Println("--------------")
	fmt.Println("Привет, Анюта!")
	fmt.Println("--------------")
	fmt.Println("Задумай число от 1 до", number)
	fmt.Println("Я его угадаю. ")
	fmt.Println("--------------")
	fmt.Println("Я буду задавать вопросы, а ты будешь мне отвечать на клавиатуре. ")
	fmt.Println("По-английски. ")
	fmt.Println("Если твой ответ - да,  то набирай на клавиатуре yes ")
	fmt.Println("Если твой ответ - нет, то набирай на клавиатуре no ")
	fmt.Println("И потом нажми Enter")
	fmt.Println("Ты поняла ?")

	var answer string
	fmt.Scan(&answer)

	if answer == "yes" {
		fmt.Println("--------------")
		fmt.Println("Погнали!")
	}

}
