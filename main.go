package main

import (
	"fmt"
	"os"
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
	fmt.Printf("starting crawl of %s", baseURL)
}