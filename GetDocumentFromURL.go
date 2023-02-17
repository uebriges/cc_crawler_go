package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

// Possible improvements:
// - error handling without termination of app

// ------------------------------------------------
// Loads and returns document from URL
// ------------------------------------------------
func GetDocumentFromURL(url string) *goquery.Document {
	fmt.Printf("Getting doc from: %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code %d: %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}
