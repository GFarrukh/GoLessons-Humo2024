package main

import (
	"errors"
	"fmt"
)

type Calculator struct{}

// Add
func (c Calculator) Add(a, b int) int {
	return a + b
}

// Divide
// метод с двумя параметарами и двумя возвращаеми значение
func (c Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	calc := Calculator{}
	sum := calc.Add(10, 5)
	fmt.Println("Sum: ", sum)

	quotient, err := calc.Divide(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Quotient: ", quotient)
	}

	_, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
