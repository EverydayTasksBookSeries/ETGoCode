package main

import (
	"os"
)

func main() {
	if err := os.MkdirAll("D:/tmp/sub", 0644); err != nil {
		panic(err)
	}
}
