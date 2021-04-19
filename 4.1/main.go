package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {

	if args := os.Args; len(args) != 3 {
		fmt.Fprintln(os.Stderr, "2 arguments required")
		os.Exit(1)
	}

	fmt.Println(sha256count(os.Args[1], os.Args[2]))
}

func sha256count(a, b string) int {
	sha1 := sha256.Sum256([]byte(a))
	sha2 := sha256.Sum256([]byte(b))
	return pop(sha1, sha2)
}

func pop(a, b [32]byte) int {
	count := 0
	for i := range a {
		count += int(pc[a[i]^b[i]])
	}
	return count
}
