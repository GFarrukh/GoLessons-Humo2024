package main

import (
	"fmt"
)

//Рассчет выплат по кредиту с фиксированной ставкой
//Начальная сумма кредита равна 3000000 рублей.
//Реализуйте функции для ежемесячного расчета выплат по кредиту с фиксированной процентной ставкой 10%.
//Выводите остаток по кредиту после каждой выплаты.
//Если остаток по кредиту становится меньше 500000 рублей, выведите сообщение "Почти погашен кредит".

func main() {
	var start_credit uint64 = 3000000
	credit(&start_credit)
}
func credit(start_credit *uint64) {
	for {
		*start_credit = *start_credit - *start_credit*10/100
		fmt.Println(*start_credit)
		if *start_credit < 500000.0 {
			fmt.Println("Почти погашен кредит")
			break
		}
	}
}
