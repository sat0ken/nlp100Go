package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Morph struct {
	surface string
	base    string
	pos     string
	pos1    string
}

func main() {
	open := func(path string) []string {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", path, err)
		}
		defer f.Close()

		lines := []string{}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if serr := scanner.Err(); serr != nil {
			fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", path, err)
		}
		return lines
	}

	kaiseki := func(line []string) [][]Morph {
		slice := [][]Morph{}
		mo_array := []Morph{}
		for _, v := range line {
			if v == `EOS` {
				slice = append(slice, mo_array)
				mo_array = []Morph{}
			} else if v != `EOS` && strings.HasPrefix(v, "*") == false {
				mo := Morph{}
				tmp := strings.Split(v, "\t")
				tmp1 := strings.Split(tmp[1], ",")
				mo.surface = tmp[0]
				mo.pos = tmp1[0]
				mo.pos1 = tmp1[1]
				mo.base = tmp1[6]
				mo_array = append(mo_array, mo)
			}
		}
		return slice
	}

	neko := kaiseki(open("../data/neko.txt.cabocha"))
	for _, v := range neko[3] {
		fmt.Println(v)
	}

}
