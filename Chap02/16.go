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

	write := func(row []string, text string) {
		f, err := os.Create(text)
		if err != nil {
			panic(err)
		}
		w := bufio.NewWriter(f)
		for _, v := range row {
			fmt.Fprint(w, v, "\n")
		}
		w.Flush()
	}

	text := read("../data/hightemp.txt")
	n, _ := strconv.Atoi(os.Args[1])
	splitcnt := len(text) / n

	tmp := []string{}
	for i := 1; i <= len(text); i++ {
		tmp = append(tmp, text[i-1])
		if i%splitcnt == 0 {
			write(tmp, strconv.Itoa(i)+".txt")
			tmp = []string{}
		}
	}

}
