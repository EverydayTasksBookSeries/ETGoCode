package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	file := "D:/tmp/test.txt"

	// 修改文件权限
	err := os.Chmod(file, 0777)
	if err != nil {
		fmt.Println(err)
	}

	// 修改atime和mtime
	err = os.Chtimes(file, time.Now(), time.Now())
	if err != nil {
		fmt.Println(err)
	}

	// 修改文件所有者，参数是uid，用户ID，以及gid，组ID
	// 注意这个函数只有在UNIX系统才有实际作用，Windows下返回错误
	err = os.Chown(file, 1, 1)
	if err != nil {
		fmt.Println(err)
	}
}
