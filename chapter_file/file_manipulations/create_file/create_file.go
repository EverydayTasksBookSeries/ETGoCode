package main

import "os"

func main() {
	file, err := os.Create("D:/etgo/new_file.txt")
	if err != nil {
		panic(err)
	}
	file.Close()
}
