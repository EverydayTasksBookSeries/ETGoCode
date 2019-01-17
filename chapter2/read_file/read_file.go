package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	f, err := os.Open("E:/a.txt")
	// 错误处理
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 新建缓冲
	buf := make([]byte, 10)
	// 读取文件内容到缓冲buf里
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		fmt.Print(string(buf[:n]))
		// 错误处理
		if err != nil {
			panic(err)
		}
	}
}
