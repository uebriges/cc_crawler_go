package main

import (
	"github.com/PuerkitoBio/goquery"
)

// ------------------------------------------------
// Collects all links in a document and returns
// a slice of links
// ------------------------------------------------
func AddLinksFromDocumentUnique(document *goquery.Document, linkMap *map[string]bool) {
	document.Find("a").Each(func(i int, s *goquery.Selection) {
		s.Each(func(i int, aTag *goquery.Selection) {
			href, exists := aTag.Attr("href")
			if exists && !LinkExists(*linkMap, href) {
				// fmt.Printf("Adding link: %s\n", href)
				(*linkMap)[href] = false
			}
		})
	})
}
