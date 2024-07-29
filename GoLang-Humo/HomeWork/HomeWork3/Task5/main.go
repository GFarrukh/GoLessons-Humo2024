package main

import "fmt"

//Функция - как переменная
//	Создайте переменную adder, которая является функцией, принимающей два целых
//числа и возвращающей их сумму.
//	Создайте переменную concatenator, которая является функцией, принимающей две
//строки и возвращающей их конкатенацию.
//	Создайте переменную isPositive, которая является функцией, принимающей целое
//число и возвращающей true, если число положительное.

func main() {
	f := adder()
	fmt.Println(f(5, 8))
	l := concatenator()
	fmt.Println(l("Go", "Lang"))
	k := isPositive()
	fmt.Println(k(20))
}
func adder() func(num1, num2 int) int {
	return func(num1, num2 int) int {
		return num1 + num2
	}
}
func concatenator() func(str1, str2 string) string {
	return func(str1, str2 string) string {
		return str1 + str2
	}
}
func isPositive() func(n int) bool {
	return func(n int) bool {
		return n%2 == 0
	}
}
