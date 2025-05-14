package link

import (
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Path string 
	Text string
}

func Parse(file_name string) []Link {
	file, err := os.Open(file_name)	
	check(err)
	defer file.Close()

	root_node, err := html.Parse(file)
	check(err)

	var links []Link
	var in_href = false

	for node := range root_node.Descendants() {
		var path, text string
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					in_href = true
					path = attr.Val
					break
				}
			}

			if in_href {
				for child_node := range node.Descendants() {
					if child_node.Type == html.TextNode {
						text = text + child_node.Data
					}
				}
			}
			links = append(links, Link{path, strings.Trim(strings.TrimSpace(text), "\n")})
			in_href = false
		}
	}

	return links
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}