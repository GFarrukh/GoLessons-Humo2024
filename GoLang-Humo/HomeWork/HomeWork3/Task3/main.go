package main

import "fmt"

//В функции передаются аргументы, но функция ничего не возвращает
//	Создайте функцию PrintNumber, которая принимает целое число и выводит его на
//экран.
//	Создайте функцию GreetUser, которая принимает строку с именем пользователя и
//выводит приветствие.
//	Создайте функцию ToggleLight, которая принимает булевое значение и выводит
//"Light on" или "Light off" в зависимости от значения аргумента.

func main() {
	var number int64 = 5
	//var str string = "Farrukh"
	//var l bool = true
	PrintNumber(number)
	//fmt.Println(GreeUser(str))
	//fmt.Println(ToggleLight(l))
}
func PrintNumber(number int64) {
	fmt.Println(number)
}
//func GreeUser(str string) {
//	fmt.Printf(str)
//}
//func ToggleLight(l bool) {
//	if l == true {
//		fmt.Println("Light on")
//	} else {
//		fmt.Println("Light off")
//	}
}
