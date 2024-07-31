package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Println("Hi, I`m %s and I'm %d years old.", p.Name, p.Age)
}

type Employee struct {
	Person
	position string
}

// Introduce
// Предопределение метода Introduce для Emloyee
func (e Employee) Introduce() {
	fmt.Printf("Hi, I`m %s and I'm %d years old.", e.Name, e.Age)
}
func main() {
	p := Person{Name: "Alice", Age: 30}
	e := Employee{Person: Person{Name: "Bob", Age: 40}, position: "Manager"}
	p.Introduce() //Output: Hi, I`m %s and I'm %d years old. Alice 30
	e.Introduce() //Output: Hi, I`m Bob and I'm 40 years old.
	e.Person.Introduce()
}
