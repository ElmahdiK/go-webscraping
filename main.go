package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name string `json:"name"`
}

// command => go run main.gp
func main() {

	scrapeUrl := "https://www.poesie-francaise.fr/poemes-auteurs/"

	c := colly.NewCollector(colly.AllowedDomains("www.poesie-francaise.fr"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US;q=0.9")
		fmt.Printf("Visiting %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while scrapin: %s\n", err.Error())
	})

	var items []item
	c.OnHTML("ul[id^=poemes_] li", func(h *colly.HTMLElement) {
		item := item{
			Name: h.Text,
		}
		items = append(items, item)
	})

	c.Visit(scrapeUrl)
	//fmt.Println(items)

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("authors.json", content, 0644)
}
