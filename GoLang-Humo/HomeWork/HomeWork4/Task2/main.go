package main

import "fmt"

//Создайте функцию PrintWeather,
//которая печатает "Солнечно", если weatherType равен "солнечно";
//"Облачно", если weatherType равен "облачно";
//и "Дождливо", если weatherType равен "дождливо".
//Переменную  weatherType вводить с консоли внутри функции.

func main() {
	PrintWeather()
}

func PrintWeather() {
	var weatherType string
	fmt.Scan(&weatherType)
	switch weatherType {
	case "солнечно":
		fmt.Println("Солнечно")
	case "облачно":
		fmt.Println("Облачно")
	case "дождливо":
		fmt.Println("Дождливо")
	}
}
