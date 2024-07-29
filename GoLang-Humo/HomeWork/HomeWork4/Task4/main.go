package main

import "fmt"

// Создайте функцию GetGrade, которая возвращает оценку
// "A", "B" или "C" в зависимости от значения переменной score.
// Переменную scope вводить с консоли внутри функции.
func main() {
	fmt.Println(GetGrade())
}
func GetGrade() string {
	var scope string
	fmt.Scan(&scope)
	return scope
}
