package main

import (
	"encoding/json"
	"fmt"
	"os"
	
   "github.com/gocolly/colly"
)

type item struct {
	Link string `json:"link"`
}

func main() {
   c := colly.NewCollector()
 
   // Find and visit all links
   c.OnHTML("div[data-role='body'] a[href^='https://github.com/']", func(e *colly.HTMLElement) {
     e.Request.Visit(e.Attr("href"))
   })
 
   var items []item
   c.OnRequest(func(r *colly.Request) {
    // fmt.Println("Visiting", r.URL)
	item := item{
		Link: r.URL.String(),
	}
	items = append(items, item)
   })

   c.Visit("https://coinmarketcap.com/fr/currencies/decentraland/")

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("authors.json", content, 0644)
}