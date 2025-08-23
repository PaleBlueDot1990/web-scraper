package main

import (
	"os"
	"fmt"
)

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
	baseURL := arguments[0]
	
	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("starting crawl of %s\n", baseURL)
	fmt.Println("----------------------------------------------------------------------")

	crawledPages := make(map[string]int)
	crawlPage(baseURL, baseURL, crawledPages)

	fmt.Println("----------------------------------------------------------------------")
	fmt.Printf("finished crawling %s\n", baseURL)
	fmt.Println("----------------------------------------------------------------------")
	
	for key, val := range crawledPages {
		fmt.Printf("%s page is crawled %d times\n", key, val)
	}
}