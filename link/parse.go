package link

import "io"

// Link contains data of link <a href="">""</a>
type Link struct {
	Href string
	Text string
}

// Parse will take an html document and will return slice of
// Link parsed from it or error
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
