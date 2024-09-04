package main

import (
	"fmt"
	"net/http"
)

//	Task5
//Создание маршрута для POST-запроса:
//
//Реализуйте обработчик для пути /post, который обрабатывает только
//POST-запросы и возвращает ответ "Received POST request".

// Обработчик для пути "/post"
func postHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что запрос является POST
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Received POST request")
	} else {
		// Если метод не POST, возвращаем ошибку 405 Method Not Allowed
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Устанавливаем обработчик для пути "/post"
	http.HandleFunc("/post", postHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
