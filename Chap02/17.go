package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

	makemap := func(arr []string) map[string]bool {
		tmp := make(map[string]bool)
		for _, v := range arr {
			_, ok := tmp[v]
			if ok == false {
				tmp[v] = true
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
	for i, _ := range makemap(row1) {
		fmt.Println(i)
	}
}
