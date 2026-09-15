package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

// Quote stores one scraped quote block: its text, author, and tags.
type Quote struct {
	Text   string   `json:"text"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}

func main() {
	var quotes []Quote

	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	// Target container: .quote — fires once per quote block on the page.
	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		var tags []string
		e.ForEach(".tag", func(_ int, tag *colly.HTMLElement) {
			tags = append(tags, tag.Text)
		})

		quote := Quote{
			Text:   e.ChildText(".text"),
			Author: e.ChildText(".author"),
			Tags:   tags,
		}
		quotes = append(quotes, quote)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	if err := c.Visit("https://quotes.toscrape.com/"); err != nil {
		log.Fatal(err)
	}

	// Print the scraped data to stdout.
	fmt.Printf("\nScraped %d quotes:\n\n", len(quotes))
	for _, q := range quotes {
		fmt.Printf("%q\n  - %s %v\n\n", q.Text, q.Author, q.Tags)
	}

	// Serialize to JSON and print it to validate.
	data, err := json.MarshalIndent(quotes, "", "  ")
	if err != nil {
		log.Fatal("failed to marshal quotes to JSON:", err)
	}
	fmt.Println("JSON output:")
	fmt.Println(string(data))

	// Write scraped data to output.json.
	if err := os.WriteFile("output.json", data, 0644); err != nil {
		log.Fatal("failed to write output.json:", err)
	}
	fmt.Printf("\nWrote %d quotes to output.json\n", len(quotes))
}