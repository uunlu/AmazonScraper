package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println("AmazonScraper started")
	Start()
}

// Start : scrapper starting point
func Start() {
	c := colly.NewCollector()
	const baseURL = "https://www.amazon.com"
	// c.OnHTML("[cel_widget_id='MAIN-SEARCH_RESULTS']", func(e *colly.HTMLElement) {
	c.OnHTML("[data-asin*='B']", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Chil)
		asin := e.Attr("data-asin")
		link := e.ChildAttr("a", "href")

		urlLink := ""
		if len(strings.TrimSpace(link)) != 0 {
			urlLink = baseURL + strings.TrimSpace(link)
		}

		title := e.ChildText("h2")
		fmt.Println(title)
		fmt.Println(urlLink)
		fmt.Println(asin)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	const AmazonURL = "https://www.amazon.com/s?k=macbook&ref=nb_sb_noss_2"
	c.Visit(AmazonURL)
}
