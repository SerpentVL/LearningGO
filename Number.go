package main

import (
	"fmt"
)

func Start() string {

	fmt.Println("Я буду задавать вопросы, а ты будешь мне отвечать на клавиатуре. ")
	fmt.Println("По-английски. ")
	fmt.Println("Если твой ответ - да,  то набирай на клавиатуре yes ")
	fmt.Println("Если твой ответ - нет, то набирай на клавиатуре no ")
	fmt.Println("И потом нажми Enter")
	fmt.Println("Ты поняла ?")

	var answer string
	fmt.Scan(&answer)
	return answer
}

func main() {

	var number int = 30
	fmt.Println("--------------")
	fmt.Println("Привет, Анюта!")
	fmt.Println("--------------")
	fmt.Println("Задумай число от 1 до", number)
	fmt.Println("Я его угадаю. ")
	fmt.Println("--------------")
	var answer string
	answer = Start()

	if answer == "yes" {
		fmt.Println("--------------")
		fmt.Println("Погнали!")
	} else if answer == "no" {
		fmt.Println("--------------")
		fmt.Println("До свидания!")
	} else {
		fmt.Println("--------------")
		fmt.Println("Ты ошиблась. Давай попробуем еще раз.")
	}
	fmt.Println("--------------")
	fmt.Println("----- КОНЕЦ -----")
}
