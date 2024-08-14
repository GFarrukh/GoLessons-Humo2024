package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error open file:", err)
		return
	}
	os.WriteFile("test.txt", []byte("text"), 0644)
	l, errr := os.ReadFile("test.txt")
	if errr != nil {
		fmt.Println("Error read file:", err)
		return
	}
	defer file.Close()
	fmt.Println(string(l))
}
