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
	maxDepth := flag.Int("depth", 3, "maximum depth links the program traverse")

	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	for _, page := range pages {
		fmt.Println(page)
	}
}

func bfs(urlStr string, depth int) []string {
	seen := make(map[string]bool)
	var q map[string]bool
	nq := map[string]bool {
		urlStr: true,
	}

	for i := 0; i <= depth; i++ {
	// for len(nq) != 0 {  // when you want to go all the way
		q, nq = nq, make(map[string]bool)
		for page := range q {
			// fmt.Println("=> ", page)
			if _, ok := seen[page]; ok {
				continue
			}
			seen[page] = true
			for _, p := range get(page) {
				// fmt.Println("==> ", page)
				nq[p] = true
			}
		}
	}

	ret := make([]string, 0, len(seen))
	for k := range seen {
		ret = append(ret, k)
	}

	return ret
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

	links := filter(hrefs(resp.Body, base), withPrefix(base))

	return links

}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string

	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}

	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)

	/* for _, link := range links {
		fmt.Println("==> ", link)
	} */

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
