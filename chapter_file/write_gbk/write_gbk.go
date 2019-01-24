package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 把标准编码UTF8的数据，转化成GBK编码
func encodeGBK(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func main() {

	// 打开UTF-8文件
	f, err := os.Open("D:/guanju.utf8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 新建GBK文件
	fout, err := os.Create("D:/guanju.gbk.txt")
	if err != nil {
		panic(err)
	}
	defer fout.Close()
	// 新建输出bufio.Writer
	bw := bufio.NewWriter(fout)
	defer bw.Flush()

	scanner := bufio.NewScanner(f)
	// 逐行读取
	for scanner.Scan() {
		line := scanner.Text()
		// 将当前行转换为GBK编码
		gbkBytes, err := encodeGBK([]byte(line))
		// 把GBK编码的数据写出到GBK文件中去
		if err != nil {
			fmt.Println(err)
		} else {
			bw.Write(gbkBytes)
			bw.WriteString("\n")
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
