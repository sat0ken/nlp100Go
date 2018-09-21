package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "I am an NLPer"
	arr := strings.Fields(word)

	fmt.Print("単語:bi-gram ")
	for i := 0; i < len(arr)-1; i++ {
		fmt.Print(arr[i], "-", arr[i+1], " ")
	}

	fmt.Println("")
	fmt.Print("文字:bi-gram ")
	for _, v := range arr {
		if 2 <= len(v) {
			for i := 0; i < len(v)-1; i++ {
				fmt.Print(v[i:i+2], " ")
			}
		}
	}

}
