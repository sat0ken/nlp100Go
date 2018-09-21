package main

import (
	"fmt"
)

func main() {
	wordA := "パトカー"
	wordB := "タクシー"
	runeWordA := []rune(wordA)
	runeWordB := []rune(wordB)

	for i := 0; i < len(runeWordA); i++ {
		fmt.Print(string(runeWordA[i]), string(runeWordB[i]))
	}
}
