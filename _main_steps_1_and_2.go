package main

import (
	"fmt"
	"net/http"
)

// Функция-контроллер (в Go это называется Handler)
func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		// Говорим коробке ответа (w), что статус теперь 404 Not Found
		w.WriteHeader(http.StatusNotFound)
		// Пишем текст ошибки для пользователя
		fmt.Fprint(w, "Ошибка 404: Страница не найдена!")
		// Останавливаем выполнение функции, чтобы код ниже не выполнялся
		return
	}

	// Отправляем обычный текстовый ответ клиенту
	fmt.Fprint(w, "Hello, World! My first web server on Go is working!")
}

// Handler для страницы about
func aboutHandler(w http.ResponseWriter, r *http.Request) {

	// Полный способ объявления переменной var название переменной тип
	//var companyName string = "NaturaPharma"

	// Краткий способ объявления переменной с автоопределением типа
	foundationYear := 2022

	currentYear := 2026

	age := currentYear - foundationYear

	// Объявляем срез (динамический список, слайс) строк. Аналог массива в php
	// [] как массив в php, далее идет тип элементов. В данном случае это строки.
	services := []string{"Консультации врачей", "Экспертиза провизорами", "Подбор форм выпуска", "NaturaPharma не моя компания. Я сторонний разработчик."}

	// Получаю get параметр из запроса
	userName := r.URL.Query().Get("name")

	// Если есть get параметр name, то добавляю имя в строку
	if userName != "" {
		fmt.Fprintf(w, "Привет, %s! Это страница о нашей компании. Нашей компании уже %d лет. Наша первая услуга %s", userName, age, services[0])

		// return обязательно, чтобы код ниже не выполнялся
		return
	}

	fmt.Fprintf(w, "Это страница о нашей компании. Нам уже %d лет. Наша первая услуга %s", age, services[0])
}

// Создал отдельный хэндлер для изучение списков-слайсов. Чтобы не перегружать aboutHandler()
func loopForHandler(w http.ResponseWriter, r *http.Request) {

	// Объявляем срез (динамический список, слайс) строк. Аналог массива в php
	services := []string{"Консультации врачей", "Экспертиза провизорами", "Подбор форм выпуска", "NaturaPharma не моя компания. Я сторонний разработчик."}

	// range возвращает две вещи: индекс элемента (i) и сам элемент (value)
	// Наконец то старый добрый for, а не вот эти вот все while....
	// Вариант с переменными index и value
	/*
		for i, value := range services {
			// На каждой итерации цикла мы видим индекс и текст
			fmt.Printf("Индекс: %d, Услуга: %s\n", i, value) // Тут не Fprintf с ответом w. Это вывод просто в консоль.
		}
	*/

	// Вариант с перебором всего списка
	/*
		// Мы говорим компилятору: "Индекс нам не нужен, просто дай значение"
		// Вариант только с одной переменной value. До этапа компиляции vs code напишет что переменная index объявлена, но не используется.
		// Нижнее подчеркивание - это исключение переменной
		for _, value := range services {
			fmt.Println(value) // тут новый Println вместо Printf
		}
	*/

	// Вариант с ограниченным количеством итераций (старый синтаксис). компиляции проходит без ошибок, но vs code пишет сообщение про старый синтаксис
	// Как в PHP/JS: for ($i = 0; $i < 5; $i++) { ... }
	/*
		for i := 0; i < 5; i++ {
			fmt.Println(i, services[0]) // Зачем тут services[0]? Чтобы vs code и компилятор не выдавали ошибку
		}
	*/

	// Вариант с ограниченным количеством итераций (новый синтаксис)
	/*
		for i := range 10 {
			// fmt.Println(i, services[i]) компиляция проходит без ошибок, но при выполнении ошибка index out of range. Потому что в списке services 4 элемента. Странно почему компилятор эту ошибку нашел.
			fmt.Println(i, services[0])
		}
	*/

	// Вариант аналог while() с заранее объявленной переменной
	i := 0

	for i < 10 {
		fmt.Println(i, services[0])
		i++
	}

}

// Создал отдельный хэндлер для выводв списков-слайсов в браузер
func loopForBrowserHandler(w http.ResponseWriter, r *http.Request) {

	// Объявляем срез (динамический список, слайс) строк. Аналог массива в php
	services := []string{"Консультации врачей", "Экспертиза провизорами", "Подбор форм выпуска", "NaturaPharma не моя компания. Я сторонний разработчик."}

	// Создаем переменную для сбора текста. Это просто текст без html. /n это перенос строки.
	var listText string = "Наши услуги:\n" // тут переменная объявлена полным способом

	// Перебираем слайс и склеиваем строки
	for _, value := range services { // на этапе обучения понятнее value. Можно запутаться service и services
		listText = listText + "- " + value + "\n"
	}

	// Отправляем готовый список в коробку ответа
	fmt.Fprint(w, listText) // Вывожу браузер, а не в консоль
}

func main() {
	// Привязываем наш контроллер к главному URL-адресу "/"
	http.HandleFunc("/", homeHandler)

	// Привязываем наш контроллер к "/about"
	http.HandleFunc("/about", aboutHandler)

	// Привязываем наш контроллер к "/loop-for"
	http.HandleFunc("/loop-for", loopForHandler)

	// Привязываем наш контроллер к "/loop-for-browser"
	http.HandleFunc("/loop-for-browser", loopForBrowserHandler)

	fmt.Println("Сервер успешно запущен на http://localhost:8080")

	// Запускаем веб-сервер на порту 8080
	// Если что-то пойдет не так, программа выдаст ошибку
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
