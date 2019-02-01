package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {

	stat, err := os.Stat("D:/sonnet18.txt")
	if err != nil {
		panic(err)
	}

	wstat := stat.Sys().(*syscall.Win32FileAttributeData)
	fmt.Printf("atime: %v\n", time.Unix(0, wstat.LastAccessTime.Nanoseconds()))
	fmt.Printf("mtime: %v\n", time.Unix(0, wstat.LastWriteTime.Nanoseconds()))
	fmt.Printf("ctime: %v\n", time.Unix(0, wstat.CreationTime.Nanoseconds()))
}
