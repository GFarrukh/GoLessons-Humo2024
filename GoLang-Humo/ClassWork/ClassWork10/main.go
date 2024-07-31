package main

import "fmt"

// Rectangle
// Определение структуры
type Rectangle struct {
	Width, Height float64
	//or width float64 \n height float64
}

// Area
// Метод с value receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Метод с pointer
func (r *Rectangle) scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Area()
	fmt.Printf("Area: %f\n", area)

	rect.scale(2)

	fmt.Printf("Scale: %d\n", rect.Area())
}

/*
type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func main() {
	c := Counter{value: 10}
	c.Increment()
	fmt.Println("Value after Increment:", c.value) // Output: Value after Increment: 11
}
*/
