package main

import "fmt"

// СЛАЙСЫ (динамический массив)

// Аналог массивов в PHP, но с обязательным обозначением типов элементов

// Добавление элементов: numbers = append(numbers, 40)

// Перебор слайса for index, value := range numbers { ... }

func FilterEven(numbers []int) []int {

	// Полный способ объявления переменной var название переменной и ее тип. В данном случае слайс из целых чисел (динамический массив)
	var result []int

	for _, val := range numbers {
		if val%2 == 0 {
			result = append(result, val)
		}
	}

	return result
}

func main() {

	numbers := []int{5, 8, 10, 20, 30, 31}

	evenNumbers := FilterEven(numbers)

	fmt.Println("Результат", evenNumbers)
}
