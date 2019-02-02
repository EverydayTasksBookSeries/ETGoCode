package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src := "D:/tmp/test.txt"
	dest := "E:/test.txt"

	// 先在目标驱动器新建文件
	oFile, err := os.Create(dest)
	if err != nil {
		panic(fmt.Sprintf("Couldn't open dest file: %v", err))
	}
	defer oFile.Close()

	// 复制文件
	iFile, err := os.Open(src)
	if err != nil {
		panic("Couldn't open source file: " + src)
	}
	_, err = io.Copy(oFile, iFile)
	iFile.Close()

	if err != nil {
		panic(fmt.Sprintf("Writing to output file failed: %v", err))
	}

}
