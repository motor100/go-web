package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// Функция принимает ОБЫЧНУЮ структуру (копию)
// celebrateBirthday(u User) - это копия.
// celebrateBirthdayReal(u *User) - это оригинал или ссылка
// Полный аналог  function($name) use ($message) - копия (надеюсь не китайская), function($name) use (&$message) - оригинал или ссылка
func celebrateBirthday(u User) {
	u.Age++ // Увеличиваем возраст внутри функции
	fmt.Println("Внутри функции:", u.Age)
}

func celebrateBirthdayReal(u *User) {
	u.Age++ // Go сам поймет, что нужно изменить данные по адресу
}

func main() {
	pete := User{Name: "Петя", Age: 20}

	//celebrateBirthday(pete) // Передали Петю
	celebrateBirthdayReal(&pete)

	fmt.Println("В main():", pete.Age) // Что выведет? Всё еще 20!
}
