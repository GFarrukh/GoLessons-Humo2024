package main

import "fmt"

//Даны два ненулевых числа.
//Найти сумму, разность, произведение и частное их квадратов.

func main() {
	var a float64 = 5
	var b float64 = 18
	Sum := a*a + b*b
	Dif := a*a - b*b
	Prod := a * a * b * b
	Quot := a * a / b * b
	fmt.Println(Sum)
	fmt.Println(Dif)
	fmt.Println(Prod)
	fmt.Println(Quot)
}
