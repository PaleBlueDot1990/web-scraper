package main 

import (
	"sync"
)

type config struct {
	crawledPages       map[string]int
	rawBaseURL         string 
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPagesToCrawl    int 
}

func configure(rawBaseURL string, maxConcurrency, maxPagesToCrawl int) *config {
	return &config{
		crawledPages:       make(map[string]int),
		rawBaseURL:         rawBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPagesToCrawl:    maxPagesToCrawl,
	}
} 

func (cfg *config) addPageVisit(normalizedCurrentURL string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, ok := cfg.crawledPages[normalizedCurrentURL]
	if ok {
		cfg.crawledPages[normalizedCurrentURL]++
		return false 
	}

	cfg.crawledPages[normalizedCurrentURL] = 1
	return true 
}

func (cfg *config) hasCrawledMaxNumberOfPages() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	return len(cfg.crawledPages) >= cfg.maxPagesToCrawl
}

