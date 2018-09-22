package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	text := read("../data/hightemp.txt")
	n, _ := strconv.Atoi(os.Args[1])
	for i := len(text) - n; i < len(text); i++ {
		fmt.Println(text[i])
	}

}
