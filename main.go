package main

import (
	"fmt"
	"os"
	"strconv"
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
	rawBaseURL, maxConcurrency, maxPagesToCrawl := getCommandLineArgs()
	cfg := configure(rawBaseURL, maxConcurrency, maxPagesToCrawl)

	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("starting crawl of %s\n", rawBaseURL)
	fmt.Println("----------------------------------------------------------------------")

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("finished crawling %s\n", rawBaseURL)
	fmt.Println("----------------------------------------------------------------------")
	
	printReport(cfg.crawledPages, rawBaseURL)
}

func getCommandLineArgs() (rawBaseURL string, maxConcurrency, maxPagesToCrawl int) {
	arguments := os.Args[1:]
	if len(arguments) < 3 {
		fmt.Printf("not enough arguments provided\n")
		fmt.Printf("please provide: <base_url> <max_concurrent_threads> <max_pages_to_crawl>\n")
		os.Exit(1)
	}
	if len(arguments) > 3 {
		fmt.Printf("too many arguments provided\n")
		fmt.Printf("please provide: <base_url> <max_concurrent_threads> <max_pages_to_crawl>\n")
		os.Exit(1)
	}

	rawBaseURL = arguments[0]

	maxConcurrency, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Printf("invalid value for <max_concurrent_threads>: %v\n", arguments[1])
		os.Exit(1)
	}

	maxPagesToCrawl, err = strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Printf("invalid value for <max_pages_to_crawl>: %v\n",arguments[2])
		os.Exit(1)
	}

	return rawBaseURL, maxConcurrency, maxPagesToCrawl
}