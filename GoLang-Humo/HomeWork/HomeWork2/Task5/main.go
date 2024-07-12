package main

import "fmt"

//Дана длина ребра куба a.
//Найти объем куба V = a3 и площадь его поверхности S = 6·a

func main() {
	var a float64 = 6.5
	V := a * a * a
	fmt.Println(V)
}
