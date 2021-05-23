package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Provide url and id.")
		os.Exit(1)
	}
	url := os.Args[1]
	id := os.Args[2]
	doc, err := getDocument(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	el := ElementByID(doc, id)
	fmt.Println(el)
}

func getDocument(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return doc, nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, match, id)
}

func forEachNode(n *html.Node, match func(n *html.Node, id string) bool, id string) *html.Node {
	if match(n, id) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		n := forEachNode(c, match, id)
		if n != nil {
			return n
		}
	}
	return nil
}

func match(n *html.Node, id string) bool {
	for _, attr := range n.Attr {
		if attr.Key == "id" {
			if attr.Val == id {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
