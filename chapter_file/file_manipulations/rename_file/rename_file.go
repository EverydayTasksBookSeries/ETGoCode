package main

import "os"

func main() {

	err := os.Rename("D:/tmp/test.txt", "D:/tmp/test1.txt")
	if err != nil {
		panic(err)
	}
}
