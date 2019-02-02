package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files := findFile("D:/tmp", "*.txt")
	fmt.Println(files)
}

// 在dir目录中，找到符合pattern的文件
func findFile(dir string, pattern string) []string {
	var files []string
	filepath.Walk(dir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			matched, err := filepath.Match(pattern, filepath.Base(path))
			if err != nil {
				log.Printf("Error in matching pattern %v for file %v", pattern, path)
				return err
			}
			if matched {
				files = append(files, path)
			}
		}
		return nil
	})
	return files
}
