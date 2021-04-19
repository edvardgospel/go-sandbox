// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for link, nr := range visit(map[string]int{}, doc) {
		fmt.Println(link, nr)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode {
		links[n.Data]++
	}

	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)

	return links
}
