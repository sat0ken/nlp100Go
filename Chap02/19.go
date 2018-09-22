package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Hightemp struct {
	pref string
	cnt  int
}

type ByCnt []Hightemp

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

	read := func(path string) []string {
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

	makemap := func(arr []string) map[string]int {
		tmp := make(map[string]int)
		for _, v := range arr {
			_, ok := tmp[v]
			if ok == false {
				tmp[v] = 1
			} else {
				tmp[v] = tmp[v] + 1
			}
		}
		return tmp
	}

	txt := read("../data/hightemp.txt")
	var row1 = []string{}
	for _, v := range txt {
		tmp := strings.Fields(v)
		row1 = append(row1, tmp[0])
	}

	arr := []Hightemp{}
	for i, v := range makemap(row1) {
		tmp := Hightemp{i, v}
		arr = append(arr, tmp)
	}
	sort.Sort(ByCnt(arr))
	for _, v := range arr {
		fmt.Printf("%v %v\n", v.pref, v.cnt)
	}
}
