package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		err := outlineHTML(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "outline: %v\n", err)
			continue
		}
	}
}

func outlineHTML(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	forEachNode(doc, startElement, endElement)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if n.FirstChild == nil {
		childElement(n)
	} else {
		if pre != nil {
			pre(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, pre, post)
		}
		if post != nil {
			post(n)
		}
	}
}

var depth int

func childElement(n *html.Node) {
	if n.Type == html.ElementNode || n.Type == html.TextNode || n.Type == html.CommentNode {
		fmt.Printf("%*s<%s/>\n", depth*4, "", n.Data)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode || n.Type == html.TextNode || n.Type == html.CommentNode {
		fmt.Printf("%*s<%s %v>\n", depth*4, "", n.Data, n.Attr)
		depth++
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode || n.Type == html.TextNode || n.Type == html.CommentNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*4, "", n.Data)
	}
}
