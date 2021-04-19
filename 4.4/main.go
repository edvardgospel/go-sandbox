package main

import "fmt"

func rotate(s []int, i int) {
	i = i % len(s)
	copy(s, append(s[i:], s[:i]...))
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	rotate(arr, 0)
	fmt.Println(arr)
}
