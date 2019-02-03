package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// isExist 判断文件是否存在
func isExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return true, err
}

// copy 复制文件内容
func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

// 将文件从src目录同步到dest目录，只拷贝更新的内容
func syncDir(src string, dest string) error {

	srcDirPath, err := filepath.Abs(src)
	if err != nil {
		panic(err)
	}
	// 遍历src目录
	filepath.Walk(src, func(path string, f os.FileInfo, _ error) error {
		// 对src目录中的每一个文件/目录，判断dest中是否存在
		// 找到dest中对应的path
		srcPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		relPath := srcPath[len(srcDirPath):]
		destPath := filepath.Join(dest, relPath)

		fmt.Println(destPath)
		// 如果是目录，就尝试新建目录
		if f.IsDir() {
			// 检查dest对象是否已经存在
			exist, _ := isExist(destPath)
			if exist {
				// TODO: 处理已存在，但是目标不是文件夹的情况
			} else { // 新建文件夹
				os.MkdirAll(destPath, f.Mode())
			}
		} else { // 如果是文件，尝试新建文件
			exist, _ := isExist(destPath)
			stat, _ := os.Stat(destPath)
			if exist {
				// 处理已存在，但是目标不是文件夹的情况
				srcMTime := f.ModTime()
				destMTime := stat.ModTime()
				srcSize := f.Size()
				destsSize := stat.Size()
				if !srcMTime.Equal(destMTime) || srcSize != destsSize {
					copy(srcPath, destPath)
				}
			} else { // 新建文件
				fmt.Println("create dest file:", destPath)
				copy(srcPath, destPath)
				os.Chtimes(destPath, time.Now(), f.ModTime())
			}
		}
		return nil
	})
	return nil
}

func main() {
	src := "D:/tmp/"
	dest := "D:/tmp1/"

	err := syncDir(src, dest)
	if err != nil {
		panic(err)
	}
}
