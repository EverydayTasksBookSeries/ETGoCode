package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	// 打开文件
	f, err := os.Open("D:/guanju.gbk.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 使用GBK编码
	enc := simplifiedchinese.GBK
	// 建立一个GBK编码的转换器Reader
	r := transform.NewReader(f, enc.NewDecoder())

	// 经过转换之后，就可以当做普通UTF-8文本来进行按行读取了
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fmt.Printf("Read line: %s\n", sc.Bytes())
	}
	if err = sc.Err(); err != nil {
		panic(err)
	}
}
