package main

import (
	"os"
)

func main() {
	if err := os.Mkdir("D:/tmp", 0644); err != nil {
		panic(err)
	}
}
