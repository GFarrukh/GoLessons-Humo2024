package main

import (
	"fmt"
	"math"
)

//Даны два неотрицательных числа a и b.
//Найти их среднее геометрическое,
//то есть квадратный корень из их произведения: √a·b.

func main() {
	var a float64 = 123
	var b float64 = 2
	N := math.Sqrt(a * b)
	fmt.Println(N)
}
