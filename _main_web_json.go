package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// WeatherMood описывает наше текущее состояние и рекомендации
type WeatherMood struct {
	// [Имя поля]    [Тип данных]     [Тег (инструкция для библиотек)]
	/* json:"weather" — это инструкция (метаданные) для встроенного пакета json, которая говорит ему: «Когда ты будешь превращать эту структуру в JSON, переименуй это поле, сделав его с маленькой буквы» */
	Weather        string   `json:"weather"`
	Recommendation string   `json:"recommendation"`
	Playlist       []string `json:"playlist"`
}

// Правило с большой (заглавной) буквы (Weather), оно публичное (экспортируемое). Его видят другие пакеты, включая пакет json

func main() {
	// 1. Регистрируем хендлер (обработчик) для главной страницы
	http.HandleFunc("/", homeHandler)

	// 2. Регистрируем хендлер, который будет отдавать JSON со структурами
	http.HandleFunc("/api/mood", moodHandler)

	fmt.Println("🌧️  Сервер запущен на http://localhost:8080")
	fmt.Println("Попробуй открыть в браузере главную или http://localhost:8080/api/mood")

	// 3. Запускаем сервер на порту 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}

// homeHandler отдает обычный текст (или HTML)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "🌧️ За окном дождь... Добро пожаловать на уютный Go-сервер.")
}

// moodHandler переводит нашу структуру в JSON и отправляет клиенту
func moodHandler(w http.ResponseWriter, r *http.Request) {
	// Создаем экземпляр структуры с данными под настроение
	currentMood := WeatherMood{
		Weather:        "Дождливо и серо",
		Recommendation: "Заварить крепкий чай, завернуться в плед и кодить на Go",
		Playlist: []string{
			"Lofi Girl - Chillhop Radio",
			"The Neighborhood - Sweater Weather",
			"Chopin - Nocturne op.9 No.2",
		},
	}

	// Устанавливаем заголовок, что мы отдаем JSON
	w.Header().Set("Content-Type", "application/json")

	// Кодируем структуру прямо в ответ (w)
	// NewEncoder делает это «на лету», что очень эффективно по памяти
	err := json.NewEncoder(w).Encode(currentMood)
	if err != nil {
		http.Error(w, "Что-то пошло не так", http.StatusInternalServerError)
	}
}
