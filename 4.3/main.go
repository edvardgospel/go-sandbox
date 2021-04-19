package main

import "fmt"

func reverse(p *[3]int) {
	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}

func main() {
	arr := [3]int{1, 2, 1}
	reverse(&arr)
	fmt.Println(arr)
}
