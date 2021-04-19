package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := "hello hello hello hey"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)

	wordCount := make(map[string]int)

	for scanner.Scan() {
		wordCount[scanner.Text()]++
	}

	for k, v := range wordCount {
		fmt.Println(k, v)
	}

}
