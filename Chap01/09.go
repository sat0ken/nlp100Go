package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func main() {
	str := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	strArr := strings.Fields(str)
	for _, word := range strArr {
		if 4 < len(word) {
			strFirst := string(word[0])
			strLast := string(word[len(word)-1])
			//fmt.Println(string(word[0]), string(word[1:len(word)-2]), string(word[len(word)-1]))
			//fmt.Println(string(word[len(word)-1]))
			rand.Seed(time.Now().UnixNano())
			d := permutations(string(word[1 : len(word)-1]))
			strShufle := d[rand.Intn(len(word))]

			fmt.Print(strFirst + strShufle + strLast + " ")
		} else {
			fmt.Print(word + " ")
		}
	}
}
