package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("D:/sonnet18.txt")
	check(err)

	N := 10
	b := make([]byte, N)
	n, err := f.Read(b)
	check(err)
	fmt.Printf("read %d bytes: [%s]\n", n, string(b))
}
