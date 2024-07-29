package main

import "fmt"

func fibonacci() func() int {
	fmt.Println("1")
	return func() int {
		f0 := 123456789
		fmt.Println("2")
		return f0
	}
}
func main() {
	f := fibonacci()
	f()
}
