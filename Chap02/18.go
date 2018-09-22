package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hightemp struct {
	pref string
	town string
	temp float64
	date string
}

type ByTemp []Hightemp

func (a ByTemp) Len() int {
	return len(a)
}

func (a ByTemp) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByTemp) Less(i, j int) bool {
	return a[i].temp < a[j].temp
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

	txt := read("../data/hightemp.txt")
	arr := []Hightemp{}
	for _, v := range txt {
		tmp := strings.Fields(v)
		ondo, _ := strconv.ParseFloat(tmp[2], 64)
		tmp1 := Hightemp{tmp[0], tmp[1], ondo, tmp[3]}
		arr = append(arr, tmp1)
	}
	sort.Sort(ByTemp(arr))
	for _, v := range arr {
		fmt.Printf("%v\t%v\t%v\t%v\n", v.pref, v.town, v.temp, v.date)
	}
}
