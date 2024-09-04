package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Отправка JSON-ответа:
//
//Настройте обработчик для пути /json, который возвращает JSON-ответ с
//любым объектом (например, {"message": "Hello, JSON!"}).

// Обработчик для пути "/json"
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// Создаём объект для JSON-ответа
	response := map[string]string{
		"message": "Hello, JSON!",
	}

	// Устанавливаем заголовок Content-Type для JSON
	w.Header().Set("Content-Type", "application/json")

	// Кодируем объект в JSON и записываем его в ответ
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Устанавливаем обработчик для пути "/json"
	http.HandleFunc("/json", jsonHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
