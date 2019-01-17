package main

import (
	"bufio"
	"fmt"
	"io"
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

	r := bufio.NewReader(f)

	for {
		s, err := r.ReadString('\n')
		// 注意读到的'\n'也保留在s中了，所以这里用fmt.Print而不是fmt.Println
		fmt.Print(s)
		if err == io.EOF {
			break
		}
	}
}
