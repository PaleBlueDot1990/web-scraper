package main

import (
	"fmt"
	"sort"
)

type visitedURL struct {
	normalizedURL string 
	visitedCount  int 
}

func printReport(crawledPages map[string]int, rawBaseURL string) {
	fmt.Printf("=============================\n")
	fmt.Printf("REPORT for %s\n", rawBaseURL)
	fmt.Printf("=============================\n")

	visitedURLs := make([]visitedURL, 0)
	for url, count := range crawledPages {
		visitedURLs = append(visitedURLs, visitedURL{
			normalizedURL: url,
			visitedCount:  count,
		})
	}

	sort.Slice(visitedURLs, func(i, j int) bool {
		if visitedURLs[i].visitedCount != visitedURLs[j].visitedCount {
			return visitedURLs[i].visitedCount > visitedURLs[j].visitedCount
		}
		return visitedURLs[i].normalizedURL < visitedURLs[j].normalizedURL
	})

	for _, u := range visitedURLs {
		fmt.Printf("Found %d internal links to https://%s\n", u.visitedCount, u.normalizedURL)
	}
}
