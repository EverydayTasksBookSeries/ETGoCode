package main

import (
	"encoding/json"
	"fmt"
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

	// 跳过"{"
	_, err = d.Token()
	if err != nil {
		panic(err)
	}

	// 一个个解析book
	var b book
	for d.More() {
		err := d.Decode(&b)
		if err != nil {
			panic(err)
		}
		fmt.Println(b.Title)
	}

	// 跳过"}""
	_, err = d.Token()
	if err != nil {
		panic(err)
	}

}
