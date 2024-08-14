package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error create file:", err)
		return
	}
	for i := 0; i < 10; i++ {
		file.Write([]byte(string(i)))
	}

	fmt.Println("Data written successfully")
	for i := 0; i < 10; i++ {
		fmt.Fprintf(file, "\n%d", i)
	}

}
