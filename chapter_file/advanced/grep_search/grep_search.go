package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func grepFile(file string, pattern string) (int64, error) {
	patCount := int64(0)
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	r, _ := regexp.Compile("M.*Bennet")
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			patCount++
			fmt.Println(file, ":", scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return patCount, err
	}
	return patCount, nil
}

func grepSearch(dir string, pattern string) int64 {
	var total int64
	filepath.Walk(dir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			count, err := grepFile(path, pattern)
			if err != nil {
				fmt.Println(err)
			}
			total = total + count
		}
		return nil
	})
	return total
}

func main() {
	pat := "said"
	total := grepSearch("E:/books/gutenberg/", pat)
	fmt.Printf("\nFound %d lines containing pattern %v\n", total, pat)
}
