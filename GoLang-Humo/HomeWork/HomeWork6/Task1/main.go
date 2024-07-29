package main

import "fmt"

//Учет накопительных счетов с ежемесячным пополнением
//Начальный баланс накопительного счета равен 100000 рублей.
//	Реализуйте функции для пополнения счета каждый месяц на
//фиксированную сумму.
//	Выводите баланс после каждого пополнения.
//	Если баланс становится больше 500000 рублей, выведите сообщение
//"Достигнут лимит накоплений".

func main() {
	var balance uint64 = 100000
	var n uint64
	var month int
	fmt.Print("На какую сумму пополнить счет каждый месяц?\n")
	fmt.Scan(&n)
	fmt.Print("на сколько месяцов открыт накопитьельный счот?\n")
	fmt.Scan(&month)
	reffill_balance(n, month, &balance)
}
func reffill_balance(n uint64, month int, balance *uint64) {
	for i := 0; i < month; i++ {
		*balance = *balance + n
		fmt.Printf("%d месяц баланс равен %d\n", i+1, *balance)
		if *balance > 500000 {
			fmt.Printf("Достигнут лимит накоплений")
			break
		}
	}
}
