package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	num := []int{1, 5, 6, 7, 8, 9, 15, 16, 19}
	word = strings.Replace(word, ".", "", -1)
	arr := strings.Fields(word)
	wordMap := make(map[string]int)

	for i, v := range arr {
		for _, vv := range num {
			if i+1 == vv {
				wordMap[string(v[0])] = i + 1
			} else {
				wordMap[string(v[0:2])] = i + 1
			}
		}
	}

	fmt.Println(wordMap)

}
