package main

import (
	"fmt"
)

func cipher(str string) []rune {
	result := []rune{}
	for _, y := range str {
		if y >= 96 && y < 123 {
			result = append(result, 219-y)
		}
	}
	return result
}

func main() {
	a := "abc"
	b := cipher(a)
	fmt.Println(string(b[:]))
}
