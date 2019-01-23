package main

import "os"

func main() {
	// 准备文本
	text := []string{
		"Sonnet 18 by William Shakespeare\n",
		"Shall I compare thee to a summer's day?\n",
		"Thou art more lovely and more temperate:\n",
		"Rough winds do shake the darling buds of May,\n",
		"And summer's lease hath all too short a date:\n",
		"Sometime too hot the eye of heaven shines,\n",
		"And often is his gold complexion dimm'd;\n",
		"And every fair from fair sometime declines,\n",
		"By chance or nature's changing course untrimm'd;\n",
		"But thy eternal summer shall not fade\n",
		"Nor lose possession of that fair thou owest;\n",
		"Nor shall Death brag thou wander'st in his shade,\n",
		"When in eternal lines to time thou growest:\n",
		"So long as men can breathe or eyes can see,\n",
		"So long lives this and this gives life to thee.",
	}

	// 创建文件
	file, err := os.Create("D:/sonnet18.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 逐行写入
	for _, line := range text {
		file.WriteString(line)
	}

	file.Sync()
}
