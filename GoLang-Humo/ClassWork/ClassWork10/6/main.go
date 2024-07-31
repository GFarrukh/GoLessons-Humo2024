package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

// NewRectangle
// Функция-конструктор для Rectangle
func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{width: width, height: height}
}
func main() {
	rect := NewRectangle(10, 5)
	fmt.Printf("Rectangle: %+v\n", rect)
}
