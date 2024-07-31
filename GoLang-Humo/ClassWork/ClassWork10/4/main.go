package main

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func main() {
	tempC := Celsius(32)
	tempF := tempC.ToFahrenheit()
	fmt.Println(tempF)
}
