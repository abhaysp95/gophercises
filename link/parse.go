package link

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

// Link contains data of link <a href="">""</a>
type Link struct {
	Href string
	Text string
}

// Parse will take an html document and will return slice of
// Link parsed from it or error
func Parse(r io.Reader) ([]Link, error) {
	rootNode, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	var links []Link
	var f func(*html.Node, string) string
	f = func(n *html.Node, pad string) string {
		var data string
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			data = f(c, pad + "  ")
		}
		fmt.Printf("\t%s%+v\n", pad, n.Data)
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println("here")
					links = append(links, Link{
						Href: a.Val,
						Text: data,
					})
					break  // other attributes of <a></a> tag doesn't matter
				}
			}
		}
		return n.Data
	}
	f(rootNode, "")

	return links, nil
}
