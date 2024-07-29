package main

import "fmt"

/*
	func checkbalance(balancePtr *float64) {
		if *balancePtr < 5000 {
			fmt.Println("Баланс низкий")
		} else {
			fmt.Printf("Текущий баланс равен %f", *balancePtr)
		}
	}

type Age int

	func main(){
		var myAge Age = 25
		fmt.Println(myAge)
	}

//-------------

type MyInt = int

	func main() {
		var a MyInt = 42
		fmt.Println(a)
	}

//---------------

	type Person struct {
		Name string
		Age  int
	}

	func main() {
		p := new(Person)
		p.Name = "Alice"
		p.Age = 25

		p1 := Person{"Bob", 30}
		fmt.Println(p1)

		p2 := Person{"lee", 23}
		fmt.Printf("%+v", p2)
	}

	type Address struct {
		City string
		State string
	}

	type Person1 struct {
		Name string
		Age int
		Address   //можно еще так   Address Address
	}

	func nestred(){
		p := Person1
		Name: "Lee"
		Age: "27"
		Address: Address {
			City: "New York",
			State: "NY",
		}

		fmt.Println(p.Name)
		fmt.Println(p.Age)
		fmt.Println(p.Address.State)
		fmt.Println(p.Address.City)

		var address Address = Address{
			City: "New York",
			State: "NY",
		}
	}

//--------------------------

	type Node struct {
		Value int
		Next  *Node
	}

	func main() {
		n1 := &Node{Value: 1}
		n2 := &Node{Value: 2}
		n3 := &Node{Value: 3}

		n1.Next = n2
		n2.Next = n3
		//n3.Next = n1 //бесконечный цикл

		current := n1
		for current != nil {
			fmt.Println(current.Value)
			current = current.Next
		}
	}

//--------------------------

func main() {
	person := struct {
		Name string
		Age  int
	}{
		Name: "John Doe",
		Age:  42,
	}
	fmt.Println(person)
}
//------------------------------
*/

type Age int
type Number int
type Score int
type Temperature float64
type Price float64

func main() {
	/*
						//1
						var age Age
						fmt.Scan(&age)
						fmt.Println(AGE(age))
						//2
						var n Number
						fmt.Scan(&n)
						fmt.Println(Num(n))
						//3
						var n Score
						fmt.Scan(&n)
						fmt.Println(diapazon(n))

					//4
					var n Count
					fmt.Scan(&n)
					Countdown(n)

				//5
				var n Temperature
				fmt.Scan(&n)
				fmt.Println(temperature(n))

			//6
			var n Price
			fmt.Scan(&n)
			fmt.Println(price(n))

		//7
		p := new(User)
		fmt.Scan(&p.Name)
		fmt.Scan(&p.Age)
		fmt.Printf(user(*p))
	*/
	//8
	l := new(Product)
	fmt.Scan(&l.Name)
	fmt.Scan(&l.Price)
	fmt.Printf(product(*l))
}

// Определение возраста совершеннолетия
// Определите тип Age на основе int. Напишите функцию,
// которая принимает возраст и возвращает сообщение о том,
// является ли человек совершеннолетним (возраст 18 лет и старше) или нет.
func AGE(age Age) string {
	if age < 18 {
		return "Не является совершеннолетним"
	} else {
		return "Является совершеннолетним"
	}
}

// Проверка на четность
// Определите тип Number на основе int. Напишите функцию, которая принимает число и возвращает сообщение о том,
// является ли оно четным или нечетным.
func Num(n Number) string {
	if n%2 == 0 {
		return "Четный"
	} else {
		return "Нечетный"
	}
}

// Проверка допустимого диапазона
// // Определите тип Score на основе int. Напишите функцию, которая принимает оценку и возвращает сообщение,
// // находится ли она в допустимом диапазоне от 0 до 100
func diapazon(n Score) string {
	if n > 0 && n < 100 {
		return "ДА"
	} else {
		return "НЕТ"
	}
}

// Обратный отсчет
// Создайте псевдоним для типа int под названием Count. Напишите функцию, которая принимает Count и выводит
// обратный отсчет до нуля.
type Count int

func Countdown(d Count) {
	for i := d; i > 0; i-- {
		fmt.Println(i)
	}
}

// Проверка температуры
// Создайте псевдоним для типа float64 под названием Temperature. Напишите функцию, которая принимает
// Temperature и выводит сообщение о состоянии (ниже нуля, выше нуля или ноль).
func temperature(n Temperature) string {
	if n == 0 {
		return "ноль"
	} else if n < 0 {
		return "ниже нуля"
	} else {
		return "выше нуля"
	}
}

// Применение скидки
// Создайте псевдоним для типа float64 под названием Price. Напишите функцию, которая принимает Price
// и возвращает новую цену с 10% скидкой.
func price(n Price) Price {
	l := n - n*0.1
	return l
}

// Информация о пользователе
// Создайте структуру User с полями Name (строка) и Age (целое число). Напишите функцию, которая
// принимает User и выводит информацию о нем.
type User struct {
	Name string
	Age  int
}

func user(u User) string {
	str := fmt.Sprintf("Имя: %s\nВозраст: %d", u.Name, u.Age)
	return str
}

// Продукт и его стоимость
// Создайте структуру Product с полями Name (строка) и Price (тип Price). Напишите функцию, которая
// принимает Product и возвращает сообщение о его стоимости.
type Product struct {
	Name  string
	Price float64
}

func product(p Product) string {
	str := fmt.Sprintf("Продукт %s стоит %.2f", p.Name, p.Price)
	return str
}
