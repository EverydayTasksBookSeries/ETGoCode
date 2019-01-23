package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {
	readGBK("E:/gbk_sample.txt")
}

// Read UTF-8 from a GBK encoded file.
func readGBK(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Use GBK encoding
	enc := simplifiedchinese.GBK
	// make a transform reader with GBK decoder
	r := transform.NewReader(f, enc.NewDecoder())

	// Read converted UTF-8 from `r` as needed.
	// As an example we'll read line-by-line showing what was read:
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fmt.Printf("Read line: %s\n", sc.Bytes())
	}
	if err = sc.Err(); err != nil {
		panic(err)
	}

}
