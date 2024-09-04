package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//	Task6
//Чтение данных из тела запроса:
//
//Создайте обработчик для пути /echo, который читает
//текстовые данные из тела запроса и возвращает их обратно клиенту.

// Обработчик для пути "/echo"
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Чтение данных из тела запроса
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// Записываем прочитанные данные обратно в ответ
	fmt.Fprintf(w, "Received: %s", string(body))
}

func main() {
	// Устанавливаем обработчик для пути "/echo"
	http.HandleFunc("/echo", echoHandler)

	// Запускаем сервер на порту 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
