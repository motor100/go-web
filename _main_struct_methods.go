package main

import "fmt"

// Описываем структуру Product
type Product struct {
	Title       string  // Название товара
	Price       float64 // Цена с цифрами после запятой для скидок
	OldPrice    float64 // Добавим поле, чтобы помнить цену БЕЗ скидки
	Description string  // Описание товара
	Image       string  // Изображение товара
	IsActive    bool    // Флаг: активен ли товар
	Quantity    int     // Количество товара на складе
}

// Это обычная функция
/*
func applyDiscount(p Product) {
	p.Price = p.Price - p.Price * 0.1 // Увеличиваем возраст внутри функции
	fmt.Println("Цена со скидкой внутри функции:", p.Price)
}
*/

// Это МЕТОД структуры Product.
// (p *Product) — это "получатель" (receiver). Аналог $this в PHP.
// Мы используем указатель *, чтобы изменить цену прямо внутри оригинального товара.
// Без указателя * (p Product), чтобы просто прочитать цену.
func (p *Product) applyDiscount(percent float64) {
	p.OldPrice = p.Price                  // Сохраняем старую цену
	p.Price = p.Price * (1 - percent/100) // Вычисляем новую цену со скидкой
}

// Метод Sell структуры Product. Уменьшает количество товара на складе. С указателем *, потому что количество товара на складе уменьшается.
func (p *Product) Sell(amount int) {
	// Проверка amount меньше или равен p.Quantity
	if amount <= p.Quantity {
		p.Quantity = p.Quantity - amount // Или сокращенно p.Quantity -= amount

		return
	}

	fmt.Printf("Ошибка: столько товара нет в наличии\n")
}

func main() {
	/*
		product1 := Product{Title: "Товар", Price: 200, Description: "Описание", Image: "image", IsActive: true, Quantity: 5}

		applyDiscount(product1)

		fmt.Println("Цена со скидкой", product1.Price)
	*/

	product1 := Product{
		Title:    "Смартфон GoPhone",
		Price:    50000.0,
		Quantity: 10,
	}

	// Вызываем метод через точку, прямо как в ООП!
	// Обрати внимание: Go сам автоматически передаст адрес (&product1) под капотом.
	// Вызываем метод applyDiscount() структуры Product с параметром 10 discount
	product1.applyDiscount(10) // Скидка 10%

	// Вызываем метод Sell и продаем товар
	product1.Sell(5) // Продали 5 товаров

	fmt.Printf("Товар: %s\n", product1.Title)
	fmt.Printf("Старая цена: %.2f руб.\n", product1.OldPrice)
	fmt.Printf("Новая цена со скидкой: %.2f руб.\n", product1.Price)
	fmt.Printf("В наличии: %d\n", product1.Quantity)

}
