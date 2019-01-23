package main

import "io/ioutil"

func main() {
	// 准备文本
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

	// 将文本转为二进制数据
	data := []byte(text)

	// 使用方便函数ioutil.WriteFile直接写入文件
	ioutil.WriteFile("D:/guanju.utf8.txt", data, 0644)
}
