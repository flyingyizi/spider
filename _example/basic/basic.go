package main

import (
	"fmt"

	"github.com/flyingyizi/spider"
	"github.com/flyingyizi/spider/core"
	"github.com/flyingyizi/spider/query"
)

func main() {
	// Instantiate default collector
	c := spider.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		spider.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *query.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *core.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}
