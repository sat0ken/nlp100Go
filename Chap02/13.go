package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arr1, arr2 := []string{}, []string{}

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

	arr1 = read("col1.txt")
	arr2 = read("col2.txt")

	write := func(row1, row2 []string, text string) {
		f, err := os.Create(text)
		if err != nil {
			panic(err)
		}
		w := bufio.NewWriter(f)
		for i, _ := range row1 {
			fmt.Fprint(w, row1[i], "\t", row2[i], "\n")
		}
		w.Flush()
	}

	write(arr1, arr2, "merge.txt")

}
