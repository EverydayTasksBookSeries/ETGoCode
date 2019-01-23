package main

import (
	"encoding/json"
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
	dec := json.NewDecoder(f)
	enc := json.NewEncoder(os.Stdout)
	var books []book
	if err := dec.Decode(&books); err != nil {
		if err == io.EOF {
			return
		}
		panic(err)
	}
	for _, v := range books {
		if err := enc.Encode(&v); err != nil {
			panic(err)
		}
	}
}
