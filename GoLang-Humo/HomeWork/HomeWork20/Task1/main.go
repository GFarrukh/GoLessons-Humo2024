package main

import (
	"fmt"
	"net/http"
)

// Task1
// Создание простого HTTP-сервера:
//
// Создайте простой HTTP-сервер на Go, который слушает порт 8080 и
// возвращает текст "Hello, World!" на любой HTTP-запрос.

// Обработчик для пути "/"
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Устанавливаем обработчики для различных путей
	http.HandleFunc("/", defaultHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
