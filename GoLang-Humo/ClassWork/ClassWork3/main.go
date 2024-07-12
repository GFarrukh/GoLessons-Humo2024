package main

import "fmt"

func sum(a, b int) int {
	c := a + b
	return c
}

func printhello(name string) {
	fmt.Println("Hello", name)
}
func printhello2() {
	fmt.Println("Hello")
}

func inc(a int) {
	a++
	fmt.Println(a)
}

func main() {
	fmt.Println(sum(5, 4))
	fmt.Println(sum(6, 7))
	fmt.Println(sum(6, 9))

	c := 6
	d := float64(c)
	fmt.Println(d)
	d = d + 0.005
	fmt.Println(d)

	x := 4.9
	l := int(x)
	fmt.Println(l)

	printhello("Vlad")

	k := sum(8, 9)
	fmt.Println(k)
	_ = sum(9, 7)

	printhello2()

	a := 6
	inc(a)
	fmt.Println(a)

	p := sum2(8)
	i := p()
	fmt.Println(i)
	summ, multiply := SumAndMultiply(5, 8)
	fmt.Println(summ, multiply)
	fmt.println()
}

func sum2(a int) func() int {
	b := 6
	return func() int {
		return a + b
	}
}

func SumAndMultiply(a, b int) (int, int) {
	sum := a + b
	multiply := a * b
	return sum, multiply
}
