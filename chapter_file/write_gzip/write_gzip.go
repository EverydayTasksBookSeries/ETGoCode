package main

import (
	"bufio"
	"compress/gzip"
	"os"
)

func main() {

	// 建立压缩文件
	f, err := os.OpenFile("D:/guanju.utf8.txt.gz", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 建立gzip.Writer
	gw := gzip.NewWriter(f)
	defer gw.Close()

	// 建立bufio.Writer
	bw := bufio.NewWriter(gw)

	// 文本
	text := `关雎 
关关雎鸠，在河之洲。
窈窕淑女，君子好逑。
参差荇菜，左右流之。
窈窕淑女，寤寐求之。
求之不得，寤寐思服。
悠哉悠哉，辗转反侧。
参差荇菜，左右采之。
窈窕淑女，琴瑟友之。
参差荇菜，左右芼之。
窈窕淑女，钟鼓乐之。`

	// 写入文本并Flush()
	bw.WriteString(text)
	bw.Flush()
}
