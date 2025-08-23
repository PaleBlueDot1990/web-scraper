package main

import (
	"fmt"
	"os"
)

/*
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
	1. Need to handle URL Type 3, 4 and 5 while crawling
*/

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(arguments) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := arguments[0]
	maxConcurrency := 5
	cfg := configure(rawBaseURL, maxConcurrency)

	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("starting crawl of %s\n", rawBaseURL)
	fmt.Println("----------------------------------------------------------------------")

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("finished crawling %s\n", rawBaseURL)
	fmt.Println("----------------------------------------------------------------------")
	
	for crawledURL, visitedCount := range cfg.crawledPages {
		fmt.Printf("%s page is crawled %d times\n", crawledURL, visitedCount)
	}
}