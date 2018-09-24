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

	file := open("../data/jawiki-country.json")
	reg := regexp.MustCompile(`^\[\[File:(.*?)\|`)

	for _, country := range file {
		if country.Title == "イギリス" {
			for _, v := range strings.Split(country.Text, "\n") {
				tmp := reg.FindStringSubmatch(v)
				if tmp != nil {
					fmt.Println(tmp[0])
				}
			}
		}
	}

}
