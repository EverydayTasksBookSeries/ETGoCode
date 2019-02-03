package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func grepFile(file string, pattern string) (int64, error) {
	patCount := int64(0)
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if bytes.Contains(scanner.Bytes(), []byte(pattern)) {
			patCount++
			fmt.Println(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		return patCount, err
	}
	return patCount, nil
}

func main() {
	pat := "said"
	total, err := grepFile("E:/books/gutenberg/pride_and_prejudice.txt", pat)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nFound %d lines containing pattern %v\n", total, pat)
}
