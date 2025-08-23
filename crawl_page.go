package main

import (
	"fmt"
	"strings"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func ()  {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.hasCrawledMaxNumberOfPages() {
		return 
	}

	if !strings.HasPrefix(rawCurrentURL, cfg.rawBaseURL) {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("normalization error - url: %s, error: %v\n", rawCurrentURL, err)
		return
	}

	isVisitingFirstTime := cfg.addPageVisit(normalizedCurrentURL)
	if !isVisitingFirstTime {
		return
	}

	fmt.Printf("Crawling %s\n", rawCurrentURL)

	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("getHTML error - url: %s, error: %v\n", rawCurrentURL, err)
		return 
	}

	linksToCrawl, err := getURLsFromHTML(currentHTML, cfg.rawBaseURL)
	if err != nil {
		fmt.Printf("getURLsFromHTML error - url: %s, error: %v\n", rawCurrentURL, err)
		return 
	}

	for _, link := range linksToCrawl {
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}
}