// crawler downloads page and visit all links on the page for each downloaded page
// go run main.go http://google.com https://ya.ru
package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

var expr = regexp.MustCompile(`<a\s+(?:[^>]*?\s+)?href=["'](.*?)["']`)
var tokens = make(chan struct{}, 5)

func extractLinks(url string) ([]string, error) {
	tokens <- struct{}{}
	defer func() {
		<-tokens
	}()
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(res.Body)
	_ = res.Body.Close()

	if err != nil {
		return nil, err
	}

	matches := expr.FindAllSubmatch(content, -1)
	links := make([]string, len(matches))

	for i, match := range matches {
		link := string(match[1])

		// to always get absolute url
		u, err := res.Request.URL.Parse(link)
		if err != nil {
			continue
		}
		links[i] = u.String()
	}

	return links, nil
}

func crawl(url string) []string {
	log.Printf("parsing %q", url)
	links, err := extractLinks(url)
	if err != nil {
		log.Printf("can't get links from %s: %v", url, err)
		return []string{}
	}

	return links
}

func main() {
	var lists = make(chan []string)
	n := 0

	n++
	go func() {
		lists <- os.Args[1:]
	}()

	var seen = make(map[string]bool)
	for ; n > 0; n-- {
		list := <-lists
		for _, url := range list {
			if !seen[url] {
				seen[url] = true
				n++
				go func(url string) {
					lists <- crawl(url)
				}(url)
			}
		}
	}
}
