package main

import "fmt"

// Чтобы тут не запутаться в этой строке, она читается так:
// func MakeGreeter(greeting string) это функция MakeGreeter которая принимает параметр greeting тип string
// Возвращает анонимную функцию func(name string), которая принимает параметр name тип string, а эта функция возращает тоже тип string
func MakeGreeter(greeting string) func(name string) string {

	return func(name string) string {
		return greeting + name
	}

}

func main() {
	hello := MakeGreeter("Привет ")

	fmt.Println("Result", hello("Иван"))
	fmt.Println("Result", hello("Вася"))
}
