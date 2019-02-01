package main

import (
	"fmt"
	"os"
)

// IsDir check whether path is a directory
func IsDir(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

// IsFile check whether path is a regular file
func IsFile(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return stat.Mode().IsRegular(), nil
}

// IsExist check whether path exists
func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return true, err
}

func main() {
	path := "D:/sonnet18.txt"

	// 是否文件夹
	isDir, _ := IsDir(path)
	fmt.Printf("isDir: %v\n", isDir)
	// 是否普通文件
	isFile, _ := IsFile(path)
	fmt.Printf("isFile: %v\n", isFile)
	// 是否存在
	isExist, _ := IsExist(path)
	fmt.Printf("exists: %v\n", isExist)
}
