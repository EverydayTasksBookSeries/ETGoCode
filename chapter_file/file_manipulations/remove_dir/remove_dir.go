package main

import "os"

func main() {
	err := os.RemoveAll("D:/tmp")
	if err != nil { // 删除失败
		panic(err)
	}
}
