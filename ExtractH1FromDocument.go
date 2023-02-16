package main

import (
	"github.com/PuerkitoBio/goquery"
)

// ------------------------------------------------
// Extracts h1 texts from a document and returns a
// slice of texts
// ------------------------------------------------
func ExtractElementsFromDocument(document *goquery.Document, elementsSelector string) (elements []string) {
	var h1Text string

	document.Find(elementsSelector).Each(func(i int, s *goquery.Selection) {
		s.Each(func(i int, h1 *goquery.Selection) {
			h1Text = h1.Text()
			elements = append(elements, h1Text)
		})
	})
	return
}
