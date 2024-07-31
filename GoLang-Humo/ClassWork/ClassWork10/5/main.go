package main

import "fmt"

// Logger
// Пустая структура
type Logger struct{}

// Log
// Метод с пустым recever
func (Logger) Log(message string) {
	fmt.Println("Log", message)
}

func main() {
	Logger{}.Log("This is a log message.")
}
