package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func main() {
	const sampleHTML = `
		<!DOCTYPE html>
		<html>
			<body>
				<a href="https://example.com">Example</a>
				<a href="https://google.com">Google</a>
				<p>No link here</p>
			</body>
		</html>`

	doc, err := html.Parse(strings.NewReader(sampleHTML))
	if err != nil {
		panic(err)
	}

	count := countLinks(doc)
	fmt.Printf("Found %d links\n", count)
}

func countLinks(n *html.Node) int {
	if n.Type == html.ElementNode && n.Data == "a" {
		return 1 + countLinks(n.FirstChild) + countLinks(n.NextSibling)
	}
	count := 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count += countLinks(c)
	}
	return count
}
