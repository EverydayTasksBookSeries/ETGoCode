package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Open file
	file, err := os.Open("E:/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// make a buffer with enough capacity
	bufCap := 128 * 1024
	buf := make([]byte, bufCap, bufCap)

	// make a scanner with the buffer
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buf, bufCap)

	// Scanner.scan() returns an enumeration of lines
	for scanner.Scan() {
		// One line of text
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
