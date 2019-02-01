package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 找到文件夹之下的所有子节点
	files, err := filepath.Glob("D:/tmp/*")
	if err != nil {
		panic(err)
	}
	// 逐一删除每个文件或文件夹
	n := 0
	for _, f := range files {
		err := os.RemoveAll(f)
		if err != nil {
			panic(err)
		}
		n++
	}
	fmt.Printf("%v nodes removed.\n", n)
}
