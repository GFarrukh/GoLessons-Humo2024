package main

import "fmt"

//Создайте функцию GetDiscount, которая возвращает скидку
//"10%" - при amount > 1000,
//"5%" - при 500 < amount <=100
//или "0%" - при amount <= 500 в зависимости от суммы покупки amount.
//Переменную amount вводить с консоли внутри функции.

func main() {
	fmt.Println(GetDiscount())
}
func GetDiscount() int {
	var amount int
	fmt.Scanf("%d", &amount)
	if amount > 1000 {
		return 10
	}
}
