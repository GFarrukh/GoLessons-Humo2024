package main

import (
	"fmt"
	"net/http"
)

//	Task2
//Использование HandleFunc для обработки пути:
//
//Настройте сервер так, чтобы он использовал http.HandleFunc
//для обработки пути /greet. На этом пути сервер должен возвращать ответ "Greetings!".

// Обработчик для пути "/greet"
func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greetings!")
}

func main() {
	// Устанавливаем обработчики для различных путей
	http.HandleFunc("/greet", greetHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
