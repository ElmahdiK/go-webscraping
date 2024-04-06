package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://www.pokepedia.fr/Ectoplasma"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	link := findArtworkLink(doc)
	if link != "" {
		fmt.Println(link)
	} else {
		fmt.Println("Aucun lien d'artwork officiel trouvé.")
	}
}

func findArtworkLink(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "class" && strings.Contains(attr.Val, "image") {
				for _, attr2 := range n.Attr {
					if attr2.Key == "href" {
						return attr2.Val
					}
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := findArtworkLink(c)
		if result != "" {
			return result
		}
	}

	return ""
}
