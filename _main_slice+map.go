package main

import "fmt"

func Unique(strings []string) []string {

	var newStrings []string

	seen := make(map[string]bool) // вот тут непонятно что за make(). Почему не как обычно map[string]bool

	/**
	Это в случае с nil картой
	Не просто создай переменную, а выдели в памяти реальную пустую хэш-таблицу, подготовь её к работе и привяжи к переменной seen

	Или сразу создать карту и заполнить ее. Как в примере _main_map.go

	make(map[string]bool, 100) Создание карты с размером 100
	*/

	for _, str := range strings {
		if !seen[str] {
			newStrings = append(newStrings, str)
			seen[str] = true
		}
	}
	return newStrings
}

func main() {

	strings := []string{"Раз", "Два", "Три", "Четыре", "Раз"}

	result := Unique(strings)

	fmt.Println("Результат", result)
}
