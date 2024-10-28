package main

import (
	"fmt"
	"log"
	"os"

	ButterCMS "github.com/ButterCMS/buttercms-go"
)

func main() {
	// Initialize the SDK with API key from environment variable
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable is not set")
	}

	// Create a new ButterCMS client
	ButterCMS.SetAuthToken(apiKey)

	// Fetch a single page
	params := map[string]string{
		"preview": "1",
	}
	page, err := ButterCMS.GetPage("*", "test-page-1", params)
	if err != nil {
		log.Fatalf("Error fetching page: %v", err)
	}

	// Print the results to console
	fmt.Printf("Page Slug: %s\n", page.Page.Slug)
	fmt.Printf("Page Status: %s\n", page.Page.Status)
	fmt.Printf("Page Scheduled Date: %s\n", page.Page.Scheduled)
}
