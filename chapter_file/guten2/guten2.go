package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// 判断正文开始的分隔行
func isStartDelim(line string) bool {
	return strings.HasPrefix(line, "*** START OF")
}

// 判断正文结尾的分隔行
func isEndDelim(line string) bool {
	return strings.HasPrefix(line, "*** END OF")
}

// 判断是否为章节标题行
func isChapterTitle(line string) bool {
	return strings.HasPrefix(line, "Chapter")
}

// 提取章节编号
func extractChapterTitle(line string) string {
	return strings.TrimSpace(line[len("Chapter"):])
}

type bookmeta struct {
	title    string
	author   string
	language string
}

// 用来收集一个章节的数据结构
type chapter struct {
	title    string
	contents []string
}

// 打印章节的统计信息
func statsChapter(ch *chapter) {
	if ch.title != "" { // 忽略掉第一章之前那段无效文字
		fmt.Printf("Chapter %v has %v paragraphs.\n", ch.title, len(ch.contents))
		fmt.Println("Last Paragraph:", ch.contents[len(ch.contents)-1])
		fmt.Println()
	}
}

// 检查书籍信息
func checkMeta(meta *bookmeta, line string) {
	titlePrefix := "Title: "
	authorPrefix := "Author: "
	langPrefix := "Language: "
	if strings.HasPrefix(line, titlePrefix) {
		meta.title = line[len(titlePrefix):]
	} else if strings.HasPrefix(line, authorPrefix) {
		meta.author = line[len(authorPrefix):]
	} else if strings.HasPrefix(line, langPrefix) {
		meta.language = line[len(langPrefix):]
	}
}

func main() {
	// 打开文件
	file, err := os.Open("E:/books/gutenberg/pride_and_prejudice.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 使用scanner按行读取
	scanner := bufio.NewScanner(file)
	isContent := false
	var curChapter chapter
	var meta bookmeta
	// 需要一个slice来存放当前段落的每一行，遇到新空白行的时候，将之前收集到的多行合并为一行。
	var curPara []string
	for scanner.Scan() {
		// 当前行
		line := strings.TrimSpace(scanner.Text())

		// 如果遇到正文开始分隔行，则进入正文
		if isStartDelim(line) {
			isContent = true
			continue
		}

		// 如果遇到正文结束分隔行，则退出循环
		if isEndDelim(line) {
			// 处理正文结束分隔行前一句多余的内容
			lastPos := len(curChapter.contents) - 1
			lastPara := curChapter.contents[lastPos]
			if strings.HasPrefix(lastPara, "End of the Project Gutenberg EBook") {
				curChapter.contents = curChapter.contents[:lastPos]
			}
			break
		}

		// 读取meta信息
		if !isContent {
			checkMeta(&meta, line)
		}

		// 如果遇到新章节的标题
		if isChapterTitle(line) {
			// 处理上一章节
			statsChapter(&curChapter)
			// 创建新章节
			chapterNum := extractChapterTitle(line)
			curChapter = chapter{chapterNum, []string{}}
			continue
		}

		// 如果是正文，则当前行加入当前章节数据中
		if isContent {
			if len(line) > 0 { // 文本
				curPara = append(curPara, line)
			} else { // 空行
				if len(curPara) > 0 {
					curChapter.contents = append(curChapter.contents, strings.Join(curPara, " "))
					curPara = []string{}
				}
			}
		}

	}
	// 处理最后一章
	statsChapter(&curChapter)

	// 打印书籍信息
	fmt.Printf("Book Meta: %#v", meta)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
