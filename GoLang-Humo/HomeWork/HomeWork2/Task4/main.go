package main

import "fmt"

// Дан диаметр окружности d.
// Найти ее длину L = π·d. В качестве значения π использовать 3.14.
var pi float64 = 3.14

func main() {
	var d float64 = 3.8
	L := pi * d
	fmt.Println(L)
}
