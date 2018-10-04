package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Sentence struct {
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

	kaiseki := func(line []string) []Sentence {
		slice := []Sentence{}
		for _, v := range line {
			if v != `EOS` {
				sent := Sentence{}
				tmp := strings.Split(v, "\t")
				sent.surface = tmp[0]
				tmp1 := strings.Split(tmp[1], ",")
				sent.pos = tmp1[0]
				sent.pos1 = tmp1[1]
				sent.base = tmp1[6]
				slice = append(slice, sent)
			}
		}
		return slice
	}

	neko := kaiseki(open("../data/neko.txt.mecab"))
	for _, v := range neko {
		if v.pos == "動詞" {
			fmt.Println(v.surface)
		}
	}

}
