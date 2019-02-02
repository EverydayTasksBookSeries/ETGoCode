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
	func() { // scope for iFile.Close()
		iFile, err := os.Open(src)
		if err != nil {
			panic("Couldn't open source file: " + src)
		}
		defer iFile.Close()
		_, err = io.Copy(oFile, iFile)

		if err != nil {
			panic(fmt.Sprintf("Writing to output file failed: %v", err))
		}
	}()
	// 复制成功，删除源文件
	err = os.Remove(src)
	if err != nil {
		panic(fmt.Sprintf("Failed removing original file: %v", err))
	}
}
