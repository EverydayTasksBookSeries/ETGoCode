package main

import (
	"os"
)

func main() {
	err := os.Remove("D:/test.txt")
	if err != nil { // 删除失败
		panic(err)
	}
}
