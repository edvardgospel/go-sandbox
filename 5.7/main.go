package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

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

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			fmt.Printf(" %v=\"%v\"", attr.Key, attr.Val)
		}
		if isSelfClosableTag(n.Data) {
			fmt.Printf("/>\n")
		} else {
			fmt.Printf(">\n")
		}
		depth++
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	} else if n.Type == html.TextNode {
		re := regexp.MustCompile(`^[ \t]*$`)
		for _, str := range strings.Split(n.Data, "\n") {
			if !re.MatchString(str) {
				fmt.Printf("%*s%s\n", depth*2, "", strings.TrimSpace(str))
			}
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if !isSelfClosableTag(n.Data) {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

var closableTags = []string{"area", "base", "br", "col", "command", "embed", "hr", "img", "input", "keygen", "link", "meta", "param", "source", "track", "wbr"}

func isSelfClosableTag(tagName string) bool {
	for _, v := range closableTags {
		if v == tagName {
			return true
		}
	}
	return false
}
