package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	files := findFileExt("D:/tmp", ".txt")
	fmt.Println(files)
}

// 在dir目录中，找到符合ext后缀的文件
func findFileExt(dir string, ext string) []string {
	var found []string
	filepath.Walk(dir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			if filepath.Ext(path) == ext {
				found = append(found, path)
			}
		}
		return nil
	})
	return found
}
