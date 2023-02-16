package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"sync"
)

// ------------------------------------------------
// Extracts a tags recursively and scrapes h1 tags
// ------------------------------------------------
func ScrapeHTMLDocument(document *goquery.Document, linkMap *map[string]bool, baseURL string, wgForLinks *sync.WaitGroup, h1Ch chan string) {
	defer wgForLinks.Done()

	aTags := document.Find("a")
	aTags.Each(func(i int, aTag *goquery.Selection) {
		href, hasHref := aTag.Attr("href")

		if hasHref && !LinkExists(*linkMap, href) {
			wgForLinks.Add(1)
			// fmt.Printf("Adding link: %s\n", href)
			(*linkMap)[href] = false
			currentDoc := GetDocumentFrom(fmt.Sprintf("%s%s", baseURL, href))
			h1s := ExtractElementsFromDocument(currentDoc, "h1")
			for _, h1 := range h1s {
				fmt.Printf("Add h1: %s\n", h1)
				h1Ch <- h1 // needs to be closed at some point
			}
			go ScrapeHTMLDocument(currentDoc, linkMap, baseURL, wgForLinks, h1Ch)
		}
	})
}
