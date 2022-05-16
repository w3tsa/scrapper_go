package main

import (
	"github.com/gocolly/colly"
)

type Item struct {
	Title 		string `json:"title"`
	Excerpt 	string `json:"excerpt"`
	Link 		string `json:"link"`
	Image 		string `json:"image"`
}


func scrapper() []Item {
	c := colly.NewCollector(
		colly.AllowedDomains("www.fodors.com"),
	)

	var items []Item

	c.OnHTML("article", func(h *colly.HTMLElement) {
		item := Item {
			Title: h.ChildText("h2.entry-title"),
			Excerpt: h.ChildText(".entry-summary p"),
			Link: h.ChildAttr(".entry-title > a", "href"),
			Image: h.ChildAttr(".entry-photo > a > img", "src"),
		}
		items = append(items, item)
	})

	c.Visit("https://www.fodors.com/news/category/family")

	return items

}
