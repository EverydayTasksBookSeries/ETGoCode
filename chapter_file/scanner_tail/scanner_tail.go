package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// 读取最后5行
	N := 5

	// 先建立一个大小为0的对列
	queue := make([]string, N)

	file, err := os.Open("D:/sonnet18.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// 逐行扫描全文
	for scanner.Scan() {
		// 将当前行加入到对列
		queue = append(queue, scanner.Text())
		// 如果对列长度超过N，则踢出最早进入对列的一行
		if len(queue) > N {
			queue = queue[1:]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// 全文扫描结束，从queue中读取最后N行进行处理
	for _, line := range queue {
		fmt.Println(line)
	}
}
