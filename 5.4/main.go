// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var linkAttrs = map[string][]string{
	"a":      {"href"},
	"link":   {"href"},
	"script": {"src"},
	"img":    {"src"},
	"iframe": {"src"},
	"form":   {"action"},
	"html":   {"manifest"},
	"video":  {"src", "poster"},
	// ...
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch05/ex04: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		for k, v := range linkAttrs {
			if n.Data == k {
				for _, attr := range v {
					for _, a := range n.Attr {
						if a.Key == attr {
							links = append(links, a.Val)
						}
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
