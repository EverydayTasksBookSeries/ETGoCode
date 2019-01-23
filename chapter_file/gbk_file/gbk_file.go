package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	readGBK("E:/hello.GBK.txt")
}

// 读取GBK编码的文件
func readGBK(filename string) {
	// 打开文件
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// GBK解码器
	enc := simplifiedchinese.GBK
	r := transform.NewReader(f, enc.NewDecoder())

	// GBK编码器
	// o, err := os.Open(filename)
	w := transform.NewWriter(f, enc.NewEncoder())

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fmt.Printf("Read line: %s\n", sc.Bytes())

		//
		w.Write(sc.Bytes())
	}
	if err = sc.Err(); err != nil {
		panic(err)
	}

}
