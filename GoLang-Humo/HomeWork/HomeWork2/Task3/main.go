package main

import "fmt"

// Даны стороны прямоугольника a и b.
// Найти его площадь S = a·b и периметр P = 2·(a + b).
func main() {
	var a float64 = 4
	var b float64 = 5
	S := a * b
	P := 2 * (a + b)
	fmt.Println(S)
	fmt.Println(P)
}
