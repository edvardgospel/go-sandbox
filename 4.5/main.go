package main

import "fmt"

func removeAdjacent(s []string) []string {
	i := 1
	for i < len(s) {
		if s[i-1] == s[i] {
			copy(s, append(s[:i], s[i+1:]...))
			s = s[:len(s)-1]
		} else {
			i++
		}
	}
	return s
}

func main() {
	s := []string{"1", "1", "1", "2", "2", "3"}
	fmt.Println(removeAdjacent(s))
}
