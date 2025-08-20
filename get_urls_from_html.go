package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	rawBaseURL = strings.TrimSuffix(rawBaseURL, "/")
	htmlReader := strings.NewReader(htmlBody)
	rootNode, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err 
	}

	var links []string
	collectLinksFromHTMLTree(rootNode, &links)
	processRelativeURLs(&links, rawBaseURL)
	return links, nil 
}

func collectLinksFromHTMLTree(currNode *html.Node, links *[]string) {
	if currNode.Type == html.ElementNode && currNode.Data == "a" {
		for _, attr := range currNode.Attr {
			if attr.Key == "href" {
				*links = append(*links, attr.Val)
			}
		}
	}

	for chilNode := currNode.FirstChild; chilNode != nil; chilNode = chilNode.NextSibling {
		collectLinksFromHTMLTree(chilNode, links)
	}
}

func processRelativeURLs(links *[]string, rawBaseURL string) {
	for idx, link := range *links {
		if !strings.HasPrefix(link, "http") {
			(*links)[idx] = rawBaseURL + link 
		}
	}
}