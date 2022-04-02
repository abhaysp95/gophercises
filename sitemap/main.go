package main

import (
	"flag"
	"fmt"
	"gophercises/link"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "provide url to sitemap")

	flag.Parse()

	pages := get(*urlFlag)
	for _, page := range pages {
		fmt.Println(page)
	}
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL {
		Scheme: reqUrl.Scheme,
		Host: reqUrl.Host,
	}
	base := baseUrl.String()

	links := filter(base, hrefs(resp.Body, base))

	return links

}

func filter(base string, links []string) []string {
	var ret []string


	return ret
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)

	for _, link := range links {
		fmt.Println("==> ", link)
	}

	var ret []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			ret = append(ret, base + link.Href)
		case strings.HasPrefix(link.Href, "http"):
			ret = append(ret, link.Href)
			// default is do nothing
		}
	}
	return ret
}


// wrote this filter logic so that maybe this can work with different url scheme
/* baseURL, err := url.Parse(base)
if err != nil {
	panic(err)
}

for _, link := range links {
	linkURL, err := url.Parse(link)
	if err != nil {
		panic(err)
	}
	if linkURL.Host == baseURL.Host {
		fmt.Println("linkURL: ", linkURL)
		ret = append(ret, linkURL.String())
	}
} */
