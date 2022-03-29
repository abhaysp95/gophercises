package link

import (
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

// TODO:
// Update main as command
// Provide flags and other stuffs for multiple files usage and to provide file path
// Use better method for string concationation (maybe buffer/stream)
// Update for any other edge cases (example included) if you find any

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
			ret = Link{
				Href: a.Val,
				Text: getText(n),
			}
			break  // ignore other attributes
		}
	}
	return ret
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data)
	}
	if n.Type != html.ElementNode {  // don't really care about any other type
		return ""
	}

	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += " " + getText(c)  // not an efficient way
	}
	return strings.Join(strings.Fields(text), " ")
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
