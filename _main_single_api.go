package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Структура для нашего "дождливого" состояния
type WeatherMood struct {
	Weather        string   `json:"weather"`
	Recommendation string   `json:"recommendation"`
	Playlist       []string `json:"playlist"`
}

// Структура специально для приема нового трека через POST-запрос
type TrackRequest struct {
	TrackName string `json:"track_name"`
}

// Глобальные переменные (для простоты, пока без БД)
// sync.Mutex нужен, чтобы безопасно изменять данные из разных запросов одновременно
var (
	mutex       sync.Mutex
	currentMood = WeatherMood{
		Weather:        "Дождливо и серо",
		Recommendation: "Заварить крепкий чай и кодить на Go",
		Playlist: []string{
			"Lofi Girl - Chillhop Radio",
			"The Neighborhood - Sweater Weather",
		},
	}
)

func main() {
	// Оставляем один роут, но внутри хендлера будем разделять GET и POST
	http.HandleFunc("/api/mood", moodHandler)

	fmt.Println("🌧️  Сервер запущен на http://localhost:8080/api/mood")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func moodHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем общий заголовок ответа для всех методов
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// --- ОБРАБОТКА GET ЗАПРОСА (Отдаем данные) ---
		mutex.Lock() // Блокируем для безопасного чтения
		json.NewEncoder(w).Encode(currentMood)
		mutex.Unlock()

	case http.MethodPost:
		// --- ОБРАБОТКА POST ЗАПРОСА (Принимаем данные) ---
		var req TrackRequest

		// json.NewDecoder читает тело запроса (r.Body) и записывает данные в структуру req.
		// Обрати внимание на оператор & — мы передаем указатель, чтобы метод мог изменить req!
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Невалидный JSON"})
			return
		}

		// Проверяем, что нам не прислали пустую строку
		if req.TrackName == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Поле track_name не должно быть пустым"})
			return
		}

		// Модифицируем наши данные в памяти
		mutex.Lock()
		currentMood.Playlist = append(currentMood.Playlist, req.TrackName)
		mutex.Unlock()

		// Возвращаем статус 201 Created и обновленный плейлист
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Трек успешно добавлен в дождливый плейлист!",
			"added":   req.TrackName,
		})

	default:
		// Если метод не GET и не POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Метод не поддерживается"})
	}
}
