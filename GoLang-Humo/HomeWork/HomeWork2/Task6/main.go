package main

import "fmt"

//Даны длины ребер a, b, c прямоугольного параллелепипеда.
//Найти его объем V = a·b·c и площадь поверхности S = 2·(a·b + b·c + a·c).

func main() {
	var a float64 = 5
	var b float64 = 6
	var c float64 = 8
	V := a * b * c
	S := 2 * (a*b + b*c + a*c)
	fmt.Println(V)
	fmt.Println(S)
}
