package main

import "fmt"

//Создайте функцию PrintGreeting, которая печатает "Доброе утро!", если dayType равен "утро";
//"Добрый день!", если dayType равен "день"; и "Добрый вечер!", если dayType равен "вечер".
//Переменную  dayType вводить с консоли внутри функции.

func main() {
	PrintGreeting()
}
func PrintGreeting() {
	var dayType string
	fmt.Scan(&dayType)
	switch dayType {
	case "утро":

		fmt.Println("Доброе утро!")
	case "день":
		fmt.Println("Добрый день!")
	case "вечер":

		fmt.Println("Добрый вечер!")
	}
}
