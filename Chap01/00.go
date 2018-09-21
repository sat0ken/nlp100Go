package main

import (
	"fmt"
)

func main() {
	word := "stressed"

	for i := len(word) - 1; i >= 0; i-- {
		fmt.Print(string(word[i]))
	}
}
