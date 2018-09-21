package main

import (
	"fmt"
)

func main() {
	word := "パタトクカシーー"
	runeWord := []rune(word)

	for i := 0; i < len(runeWord); i++ {
		if i%2 == 0 {
			fmt.Print(string(runeWord[i]))
		}
	}
}
