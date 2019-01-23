package main

import (
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("D:/sonnet18.txt")
	check(err)

	N := 100

	// 跳转到倒数第N个字节
	f.Seek(-int64(N), io.SeekEnd)

	b := make([]byte, 10)
	nRead := 0
	for {
		n, err := f.Read(b)
		if n == 0 {
			break
		}
		check(err)
		// 记录总共读了多少字节
		nRead += n
		// 如果还没足够N，处理掉当前读到的数据，并继续
		if nRead < N {
			// 处理当前的一段
			fmt.Printf("read %d bytes: [%s]\n", n, string(b))
		} else { // 足够N了，处理掉最后读到的一小段数据（直到N为止），然后跳出
			// 处理最后读到的一小段
			nTail := n - (nRead - N)
			fmt.Printf("read %d bytes: [%s]\n", nTail, string(b[:nTail]))
			break
		}
	}
}
