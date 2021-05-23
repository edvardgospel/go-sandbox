package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(expand("$abcdef", helper))
}

func expand(s string, f func(string) string) string {
	re := regexp.MustCompile(`\$[^\s]+`)
	return re.ReplaceAllStringFunc(s, func(x string) string {
		return f(x[1:])
	})
}

func helper(s string) string {
	return s + s
}
