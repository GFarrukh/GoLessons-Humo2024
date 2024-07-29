package main

import "fmt"

//Создайте функцию PrintTrafficLight, которая печатает
//"Стоп", если lightColor равен "красный";
//"Внимание", если lightColor равен "желтый"; и
//"Идите", если lightColor равен "зеленый".
//Переменную lightColor вводить с консоли внутри функции.

func main() {
	PrintTrafficLight()
}
func PrintTrafficLight() {
	var lightColor string
	fmt.Scan(&lightColor)
	switch lightColor {
	case "красный":
		fmt.Println("Стоп")
	case "желтый":
		fmt.Println("Внимание")
	case "зеленый":
		fmt.Println("Идите")
	}
}
