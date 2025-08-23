package main

import (
	"fmt"
	"strings"
)

/*
params:
	1.rawCurrentURL: 
	  current URL we are crawling

	2.rawBaseURL:          
	  root URL of the website we are crawling

crawlPage algorithm:   
	1.
	2.
	3.
	4.
	5.

HTML URL Examples- 
	1.Absolute URLs
	  https://example.com/about
	  http://anotherdomain.org/page
	2.Root-Relative URLs
	  /about
	  /contact/index.html
	3.Path-Relative URLs (relative to the current page’s path)
	  section.html
	  ../images/pic.png
	  ./sibling-page
	4.Fragments only (reference within current page)
	  #section1
	5.Non-HTTP(S) & Invalid Schemes
	  mailto:someone@example.com
	  tel:+123456789

TODO- 
	1. Need to handle URL Type 3, 4 and 5 
*/

func crawlPage(rawBaseURL, rawCurrentURL string, crawledPages map[string]int) {
	fmt.Printf("Crawling %s\n", rawCurrentURL)

	if !strings.HasPrefix(rawCurrentURL, rawBaseURL) {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("normalization error - url: %s, error: %v\n", rawCurrentURL, err)
		return
	}

	_, ok := crawledPages[normalizedCurrentURL]
	if ok {
		crawledPages[normalizedCurrentURL]++;
		return
	}
	crawledPages[normalizedCurrentURL] = 1

	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("getHTML error - url: %s, error: %v\n", rawCurrentURL, err)
		return 
	}

	linksToCrawl, err := getURLsFromHTML(currentHTML, rawBaseURL)
	if err != nil {
		fmt.Printf("getURLsFromHTML error - url: %s, error: %v\n", rawCurrentURL, err)
		return 
	}

	for _, link := range linksToCrawl {
		crawlPage(rawBaseURL, link, crawledPages)
	}
}