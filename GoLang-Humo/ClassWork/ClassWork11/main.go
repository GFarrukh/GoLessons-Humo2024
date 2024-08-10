package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello, World!"
	fmt.Println(str[0])
	b := str[0]
	fmt.Println(string(b))

	sub := str[7:12]
	fmt.Println(sub)

	contains1 := strings.Contains(str, "World!")
	fmt.Println(contains1)
	contains2 := strings.Contains(str, "World!!")
	fmt.Println(contains2)
	contains3 := strings.Contains(str, "World")
	fmt.Println(contains3)
	index := strings.Index(str, "Wo")
	fmt.Println(index)

	s := "Hello, 法魯克"
	bytes := []byte(s)
	fmt.Println(bytes)
	fmt.Println(string(bytes))

	for _, v := range s {
		fmt.Println(v)
		fmt.Println(string(v))
	}

	for i := 0; i < len(s); i++ {
		//fmt.Println(s[i])
		fmt.Println(string(s[i]))
	}

	str1 := "Aé世😀"
	for len(str1) > 0 {
		r, size := utf8.DecodeRuneInString(str1)
		fmt.Printf("Руна: %c, Unicode: U+%04X, Занимает %d байт\n", r, r, size)
		str1 = str1[size:]
	}
	var m1 []string
	m := "Hqeqlqlqo"
	m1 = strings.Split(m, "q")
	fmt.Println(m1)
	fmt.Println(m1[3])
}
