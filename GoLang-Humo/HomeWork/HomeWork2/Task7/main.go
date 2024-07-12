package main

import "fmt"

// Найти длину окружности L и площадь круга S
// заданного радиуса R: L = 2·π·R, S = π·R
// В качестве значения π использовать 3.14.
var pi float64 = 3.14

func main() {
	var R float64 = 8.4
	L := 2 * pi * R
	S := pi * R
	fmt.Println(L)
	fmt.Println(S)
}
