package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	word = strings.Replace(word, ",", "", -1)
	word = strings.Replace(word, ".", "", -1)

	arr := strings.Fields(word)
	list := []int{}

	for _, v := range arr {
		list = append(list, len(v))
	}
	fmt.Println(list)
}
