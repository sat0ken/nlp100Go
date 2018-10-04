package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Sentence struct {
	surface string
	base    string
	pos     string
	pos1    string
}

type WordCount struct {
	sent string
	cnt  int
}

type ByCnt []WordCount

func (a ByCnt) Len() int {
	return len(a)
}

func (a ByCnt) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByCnt) Less(i, j int) bool {
	return a[i].cnt > a[j].cnt
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
	count := make(map[string]int)
	for i := 0; i < len(neko); i++ {
		_, ok := count[neko[i].surface]
		if ok == false {
			count[neko[i].surface] = 1
		} else {
			count[neko[i].surface]++
		}
	}
	arr := []WordCount{}
	for k, v := range count {
		arr = append(arr, WordCount{k, v})
	}
	sort.Sort(ByCnt(arr))
	for i := 0; i < 10; i++ {
		fmt.Println(arr[i].sent, arr[i].cnt)
	}
}
