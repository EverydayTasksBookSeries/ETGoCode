package main

import (
	"fmt"
	"os"
)

func main() {

	stat, err := os.Stat("D:/sonnet18.txt")
	if err != nil {
		panic(err)
	}

	// 文件名
	fmt.Printf("fileName: %v", stat.Name())
	// 是否文件夹
	fmt.Printf("isDir: %v\n", stat.Mode().IsDir())
	// 上面同样功能的便捷函数
	fmt.Printf("isDir: %v\n", stat.IsDir())
	// 是否普通文件
	fmt.Printf("isRegular: %v\n", stat.Mode().IsRegular())
	// 文件权限
	fmt.Printf("permission: %v\n", stat.Mode().Perm())
	// 修改时间
	fmt.Printf("lastModTime: %v\n", stat.ModTime())
	// 文件大小
	fmt.Printf("size: %v\n", stat.Size())

}
