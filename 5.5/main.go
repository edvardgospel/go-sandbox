// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return countWordsAndImages(n.NextSibling)
	} else if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}

func main() {
	words, images, err := CountWordsAndImages("http://golang.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	fmt.Printf("words:  %d\n", words)
	fmt.Printf("images: %d\n", images)
}
