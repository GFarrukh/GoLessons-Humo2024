package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	ID     int
	Title  string
	Author string
}

func main() {
	book := Book{
		ID:     1,
		Title:  "GoLang-Humo",
		Author: "Farrukh",
	}
	// Открытие файла для записи
	file, err := os.OpenFile("book.json", os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	// Создание нового JSON-энкодера и запись данных в файл
	encoder := json.NewEncoder(file)
	err = encoder.Encode(book)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
	}
	fmt.Println("Successfully wrote JSON to file")

	// Открытие файла для чтения
	file1, err1 := os.Open("book.json")
	if err1 != nil {
		fmt.Println("Error opening file:", err1)
		return
	}
	defer file1.Close()

	// Создание нового JSON-декодера и чтение данных из файла
	decoder := json.NewDecoder(file1)
	for {
		err = decoder.Decode(&book)
		if err != nil {
			fmt.Println("Конец файла")
			return
		}
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
