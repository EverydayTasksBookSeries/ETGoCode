package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 找到符合通配符的文件
	files, err := filepath.Glob("D:/tmp/test*.txt")
	if err != nil {
		panic(err)
	}
	// 逐一删除每个文件
	n := 0
	for _, f := range files {
		err := os.Remove(f)
		if err != nil {
			panic(err)
		}
		n++
	}
	fmt.Printf("%v files removed.\n", n)
}
