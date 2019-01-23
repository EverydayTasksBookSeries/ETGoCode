package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	// 打开压缩问价
	f, err := os.Open("D:/sonnet18.txt.gz")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 打开gzip的Reader，套在文件Reader之上
	gr, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	defer gr.Close()

	// 新建一个Scanner，读取gzip.Reader
	scanner := bufio.NewScanner(gr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
