package main

import "fmt"

// Карта map это аналог ассоциативного массива в php $arr = [1 => 'Аня', 2 => 'Петя']
// Слайс slice это аналог нумерованного массива в php $arr = ['Аня', 'Петя']
// Тут строго ключ тип int, значение тип string

func FindUser(id int) (string, bool) {
	users := map[int]string{
		1: "Аня",
		2: "Петя",
		3: "Вова",
	}

	/*
		* Вариант с перебором карты
		for key, value := range users {
			if key == id {
				fmt.Println("Name ", value)
			}
		}
	*/

	name, ok := users[id]

	return name, ok
}

func main() {
	name, ok := FindUser(1)

	if !ok {
		fmt.Println("Ошибка. Нет такого id")
		return
	}

	fmt.Println("Имя ", name)
}
