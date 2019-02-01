package main

import (
	"os"
	"path/filepath"
)

func main() {
	path := "D:/etgo1/new_file.txt"
	createWithDir(path)
}

func createWithDir(path string) {
	dir := filepath.Dir(path)
	if fi, err := os.Stat(dir); err == nil {
		// 文件夹路径存在，但该路径不是文件夹
		if !fi.IsDir() {
			panic(dir + " is not a directory")
		}
	} else if os.IsNotExist(err) { // 文件夹不存在，尝试新建
		os.MkdirAll(dir, 0644)
	} else { // 其他情况，比如没有查看权限
		panic(err)
	}
	// 文件夹已经建好，在其中新建文件
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	file.Close()
}
