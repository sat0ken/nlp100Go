package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filepath string) ([]string, error) {

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
		return nil, serr
	}
	return lines, nil

}

func main() {

	file, err := readFile("../data/hightemp.txt")
	if err == nil {
		fmt.Println(len(file))
	} else {
		fmt.Println(err)
	}

}
