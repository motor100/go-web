package main

import (
	"errors"
	"fmt"
)

// Описываем структуру User
// Это аналог класса. А создание переменной с типом данных User - это аналог объекта
type User struct {
	// public/private Name (с заглавной) - доступно из других пакетов. name (со строчной) - доступно только локально.
	// В этом случае все поля public. Это поля а не ключи массива и не свойства объекта
	Name     string  // Имя пользователя
	Age      int     // Возраст
	IsActive bool    // Флаг: активен ли аккаунт
	NewStr   string  // Строка для тестирования конкатенации строк
	Res      float64 // Число для тестирование деления в методе структуры
}

// Описываем структуру Product
type Product struct {
	Title       string  // Название товара
	Price       float64 // Цена с цифрами после запятой для скидок
	Description string  // Описание товара
	IsActive    bool    // Активность товара
}

/*
// Это просто функция которая сама по себе. Это не метод структуры
func setString(u *User) string {
	u.NewStr = u.Name + strconv.Itoa(u.Age)
	fmt.Println("Внутри функции:", u.NewStr)
	return u.NewStr <- функция возвращает. Как и любая другая обычная функция в PHP
}
*/

/*
u.Age — это int.
num — это int.
Операция int / int всегда даёт результат int (целочисленное деление, остаток просто отбрасывается).
То есть 28 / 100 в мире int — это ровно 0.
Затем вы пытаетесь записать этот int в поле u.Res, которое имеет тип float64.
Компилятор видит несоответствие типов и выдает ошибку.
*/

// Метод стуктуры User
func (u *User) setString(num int) (float64, error) {
	if num == 0 {
		return 0, errors.New("На ноль делить нельзя")
	}
	u.Res = float64(u.Age) / float64(num)

	return u.Res, nil
}

// Функция которая возвращает структуру (функция-фабрика)
func NewProduct(title string, price float64, description string) *Product {
	return &Product{
		Title:       title,
		Price:       price,
		Description: description,
		IsActive:    true, // Значение по умолчанию для всех новых пользователей
	}
}

func main() {
	admin := User{
		Name:     "Евгений",
		Age:      28,
		IsActive: true,
	}

	user := User{
		Name:     "Иван",
		Age:      20,
		IsActive: false,
	}

	//fmt.Println("Результат", user)

	//fmt.Println("Результат", admin.Name)
	//fmt.Println("Результат", user.IsActive)

	//user.IsActive = true

	fmt.Println("Результат", user)

	//setString(&admin) <- просто вызываю функцию, без присваивания переменной

	_, err := admin.setString(100500)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Результат", admin.Res)

	newProduct := NewProduct("Товар1", 65, "Описание")

	fmt.Println("Товар", newProduct)

}
