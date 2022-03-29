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
	nodes := linkNodes(rootNode)
	for _, node := range nodes {
		links = append(links, buildLink(node))
		fmt.Println(node)
	}

	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, a := range n.Attr {
		if a.Key == "href" {
			ret = Link{Href: a.Val}
			break  // ignore other attributes
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			ret.Text = c.Data
			break
		}
	}
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var nodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...)
	}
	return nodes
}


// Initial logic for generalized parsing of links with href and their text
/* var links []Link
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
	} else if n.Type == html.TextNode {
		return n.Data
	}
	return ""
}
f(rootNode, "") */
