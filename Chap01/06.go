package main

import (
	"fmt"
)

func syugou(map1, map2 map[string]bool, check string) {
	wa, seki, sa := []string{}, []string{}, []string{}

	for k1, _ := range map1 {
		wa = append(wa, k1)
		_, ok := map2[k1]
		if ok == false {
			sa = append(sa, k1)
		} else {
			seki = append(seki, k1)
		}
	}
	for k2, _ := range map2 {
		_, ok := map1[k2]
		if ok == false {
			wa = append(wa, k2)
		}
	}

	fmt.Print(wa, "\n", seki, "\n", sa, "\n")
	fmt.Printf("X in 'se' is %v\n", map1[check])
	fmt.Printf("Y in 'se' is %v\n", map2[check])

}

func main() {
	word1 := "paraparaparadise"
	word2 := "paragraph"

	bigram := func(str string) []string {
		str_bigram := []string{}
		for i := 0; i < len(str)-1; i++ {
			str_bigram = append(str_bigram, str[i:i+2])
		}
		return str_bigram
	}
	word1_bi, word2_bi := bigram(word1), bigram(word2)
	fmt.Println("X:", word1_bi)
	fmt.Println("Y:", word2_bi)

	makemap := func(arr []string) map[string]bool {
		tmp_map := make(map[string]bool)
		for _, v := range arr {
			_, ok := tmp_map[v]
			if ok == false {
				tmp_map[v] = true
			}
		}
		return tmp_map
	}

	syugou(makemap(word1_bi), makemap(word2_bi), "se")

}
