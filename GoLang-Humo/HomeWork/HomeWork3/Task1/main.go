package main

import "fmt"

// В функции не передаются аргументы и ничего не возвращает функция
//
//	Создайте функцию PrintGreeting, которая печатает "Hello, World!" на экран.
//	Создайте функцию DisplaySeparator, которая печатает строку из 20 символов - для
// разделения текста.
//	Создайте функцию NotifyStart, которая выводит на экран сообщение "Program
// started".

func main() {
	PrintGreeting()
	DisplaySteparator()
	NotifyStart()
}
func PrintGreeting() {
	fmt.Println("Hello World!")
}
func DisplaySteparator() {
	fmt.Println("--------------------")
}
func NotifyStart() {
	fmt.Println("Program started")
}
