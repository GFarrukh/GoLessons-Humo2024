package main

import "fmt"

//В функции передаются аргументы, но функция возвращает значение(значения)
//	Создайте функцию Add, которая принимает два целых числа и возвращает их сумму.
//	Создайте функцию Concat, которая принимает две строки и возвращает их
//конкатенацию.
//	Создайте функцию IsEven, которая принимает целое число и возвращает true, если
//число четное, и false в противном случае.

func main() {
	fmt.Println(Add(5, 9))
	fmt.Println(Concat("Far", "rukh"))
	fmt.Println(IsEven(19))
}
func Add(num1, num2 int) int {
	return num1 + num2
}
func Concat(str1, str2 string) string {
	return str1 + str2
}
func IsEven(num int) bool {
	return num%2 == 0
}
