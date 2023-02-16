package main

import (
	"sync"
)

func main() {
	var wgForLinks sync.WaitGroup
	visitedURLs := make(map[string]bool)
	var h1s []string
	h1Ch := make(chan string)

	baseURL := "http://books.toscrape.com/"

	doc := GetDocumentFrom(baseURL)

	wgForLinks.Add(1)

	// Is actually already the worker
	// Next step: Extraction of H1s. Channel to collect H1s?
	go ScrapeHTMLDocument(doc, &visitedURLs, baseURL, &wgForLinks, h1Ch)

	for {
		newH1, open := <-h1Ch
		h1s = append(h1s, newH1)
		if !open {
			break
		}
	}

	wgForLinks.Wait()
}
