package main

import "fmt"

//Даны целые положительные числа A и B (A > B).
//На отрезке длины A размещено максимально возможное количество отрезков длины B (без наложений).
//Используя операцию взятия остатка от деления нацело, найти длину незанятой части отрезка A.

func main() {
	var A uint64 = 256
	var B uint64 = 13
	n := A % B
	fmt.Println(n)
}
