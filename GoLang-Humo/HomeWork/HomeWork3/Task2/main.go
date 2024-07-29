package main

import "fmt"

//В функции не передаются аргументы, но функция возвращает значение (значения)
//	Создайте функцию GetWelcomeMessage, которая возвращает строку "Welcome!".
//	Создайте функцию GetPiValue, которая возвращает значение числа π с точностью до
//двух знаков после запятой (3.14).
//	Создайте функцию IsServerActive, которая возвращает булево значение true.

func main() {
	fmt.Println(GetWelcomeMessage())
	fmt.Println(GetPiValue())
	fmt.Println(IsServerActive())
}
func GetWelcomeMessage() string {
	return "Welcome!"
}
func GetPiValue() float64 {
	return 3.14
}
func IsServerActive() bool {
	return true
}
