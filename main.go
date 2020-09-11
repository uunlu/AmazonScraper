package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("AmazonScraper started")
	c := colly.NewCollector()

	c.OnHTML("[cel_widget_id='MAIN-SEARCH_RESULTS'] h2", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Chil)

		link := e.ChildAttr("a", "href")
		title := strings.TrimSpace(link)
		fmt.Println(title)
		fmt.Println(link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	const AmazonURL = "https://www.amazon.com/s?k=macbook&ref=nb_sb_noss_2"
	c.Visit(AmazonURL)
}
