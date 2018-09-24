package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type Country struct {
	Text  string
	Title string
}

func main() {
	open := func(path string) []Country {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", path, err)
		}
		defer f.Close()

		jsonline := []Country{}
		r := bufio.NewReader(f)
		for {
			b, err := r.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			tmp := Country{}
			json.Unmarshal([]byte(b), &tmp)
			jsonline = append(jsonline, tmp)
		}
		return jsonline
	}

	replace := func(reg_pattern []string, line string) string {
		var rtn string = line
		for _, v := range reg_pattern {
			reg := regexp.MustCompile(v)
			rtn = reg.ReplaceAllString(rtn, "$1")
		}
		return rtn
	}

	file := open("../data/jawiki-country.json")
	pattern := []string{`\[\[Category:`, `]]`, `\|\*`}
	for _, country := range file {
		if country.Title == "イギリス" {
			for _, v := range strings.Split(country.Text, "\n") {
				if strings.Contains(v, "Category") {
					fmt.Println(replace(pattern, v))
				}
			}
		}
	}
}
