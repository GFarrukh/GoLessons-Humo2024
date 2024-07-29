package main

import "fmt"

//Создайте функцию GetTemperatureDescription,
//которая возвращает "Холодно" (меньше 10) ,
//"Тепло" (с 10 до 25)
//или "Жарко" в зависимости от значения переменной temperature.
//Переменную temperature вводить с консоли внутри функции.

func main() {
	fmt.Println(GetTemperatureDescription())
}
func GetTemperatureDescription() int {
	var temperature int
	fmt.Scanf("%d", &temperature)
	if temperature < 10 {
		return 10
	}
}
