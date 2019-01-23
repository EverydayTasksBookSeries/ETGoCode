package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type book struct {
	ID     int64
	Title  string
	Author string
	Intro  string
}

func main() {
	f, err := os.Open("D:/books.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	d := json.NewDecoder(f)
	level := 0
	var prev string
	for {
		t, tokenErr := d.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				break
			}
		}

		switch t.(type) {
		case json.Delim:
			ts := t.(json.Delim).String()
			// 判断层级
			if ts == "{" || ts == "[" {
				level++
			}
			if ts == "}" || ts == "]" {
				level--
			}
		default:
			// 我们的需求：获取级别2的"author"对应的值
			if prev == "author" && level == 2 {
				fmt.Println(t)
			}
		}

		// 记录键，注意这里凡是字符串类型都记录了，所以必须配合级别来避免错误
		switch t.(type) {
		case string:
			prev = t.(string)
		}
	}
}
