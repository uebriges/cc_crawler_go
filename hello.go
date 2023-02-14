package main

import (
	"fmt"
	"sync"
)

func main() {
	visitedURLs := make(map[string]bool)
	var h1s []string
	var wg sync.WaitGroup

	urlCh := make(chan string)
	h1Ch := make(chan string)

	baseURL := "http://books.toscrape.com/"

	doc := GetDocumentFrom(baseURL)

	AddLinksFromDocumentUnique(doc, &visitedURLs)

	go func() {
		for currentUrl := range visitedURLs {
			wg.Add(1)
			fmt.Println(fmt.Sprintf("%s%s", baseURL, currentUrl))
			urlCh <- currentUrl
			go worker(urlCh, h1Ch, &visitedURLs, h1s)
		}
	}()

	for k := range visitedURLs {
		if visitedURLs[k] == false {
			// fmt.Printf("Checking %s\n", fmt.Sprintf("%s%s", baseURL, k))
			doc = GetDocumentFrom(fmt.Sprintf("%s%s", baseURL, k))
			AddLinksFromDocumentUnique(doc, &visitedURLs)
			h1s = append(h1s, ExtractH1FromDocument(doc)...)
			visitedURLs[k] = true
		}
	}

	fmt.Println("These are the extracted H1s")
	for _, v := range h1s {
		fmt.Printf("%s\n", v)
	}

}

func worker(url <-chan string, h1s chan<- string, visitedURLs *map[string]bool, h1 []string) {
	currentURL := <-url
	fmt.Printf("Worker with %s\n", currentURL)
	doc := GetDocumentFrom(currentURL)
	AddLinksFromDocumentUnique(doc, *&visitedURLs)
}
