package main

import "io/ioutil"

func main() {
	data := []byte("Hello World!\n")
	ioutil.WriteFile("E:/hello_out.txt", data, 0644)
}
