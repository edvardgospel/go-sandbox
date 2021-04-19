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
	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}

	if n.FirstChild != nil && n.FirstChild.Data != "script" && n.FirstChild.Data != "style" {
		texts = visit(texts, n.FirstChild)
	}
	if n.NextSibling != nil && n.NextSibling.Data != "script" && n.NextSibling.Data != "style" {
		texts = visit(texts, n.NextSibling)
	}

	return texts
}
