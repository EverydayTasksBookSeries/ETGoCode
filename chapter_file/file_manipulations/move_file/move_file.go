package main

import "os"

func main() {

	// 新建目录
	err := os.MkdirAll("D:/tmp1", 0644)
	if err != nil {
		panic(err)
	}

	// 把文件移动到新目录
	err = os.Rename("D:/tmp/test_move.txt", "D:/tmp1/test_move.txt")
	if err != nil {
		panic(err)
	}
}
