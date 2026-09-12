package main

import (
	"errors"
	"fmt"
)

// Divide() - глобальная видимость
// divide() - локальная видимость
func Divide(a int, b int) (int, error) {

	// Проверка b == 0
	if b == 0 {
		return 0, errors.New("На ноль делить нельзя")
	}

	return a / b, nil
}

func main() {
	res, err := Divide(4, 0)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Result", res)
}
