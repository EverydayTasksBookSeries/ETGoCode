package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开CSV文件
	f, err := os.Open("D:/test.csv")
	if err != nil {
		panic(err)
	}

	// 建立csv.Reader
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// 遇到EOF跳出
		if err == io.EOF {
			break
		}
		// 打印record，它是一个[]string数组
		fmt.Println(record)
		// 遍历打印其中每一个元素
		for value := range record {
			fmt.Printf("  %v\n", record[value])
		}
	}
}
