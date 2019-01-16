package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// read first N lines
	M := 1
	N := 3
	file, err := os.Open("E:/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	n := 0
	// Scanner.scan() returns an enumeration of lines
	for scanner.Scan() {
		n++
		// if current line number is less than M, ignore and continue
		if n <= M {
			continue
		}
		// if current line number is greater than N, quit
		if n > N {
			break
		}
		// One line of text
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
