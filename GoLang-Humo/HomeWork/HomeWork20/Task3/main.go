package main

import (
	"fmt"
	"net/http"
)

//	Task3
//Обработка запроса с параметром:
//
//Создайте обработчик для пути /hello, который принимает параметр
//name из строки запроса (например, /hello?name=John) и отвечает "Hello, John!".

// Обработчик для пути "/hello"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр "name" из строки запроса
	name := r.URL.Query().Get("name")

	// Если параметр "name" пустой, используем значение по умолчанию
	if name == "" {
		name = "World"
	}

	// Формируем ответ
	response := fmt.Sprintf("Hello, %s!", name)
	fmt.Fprintf(w, response)
}

func main() {
	// Устанавливаем обработчик для пути "/hello"
	http.HandleFunc("/hello", helloHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
