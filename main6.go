package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"net/http"
	
   "github.com/gocolly/colly"
)

// Structure pour représenter l'objet JSON


type githubInfos struct {
	// Ajoutez ici les champs de l'objet JSON que vous attendez
	Type  string `json:"type"`
	Date string `json:"created_at"`
}

type gitLink struct {
	Username string `json:"username"`
	Link string `json:"link"`
	LastActivity *githubInfos `json:"last_activity"`
}

type item struct {
	Name string `json:"name"`
	Cmc string `json:"cmc"`
	Github []gitLink `json:"github"`
}

func main() {
   c := colly.NewCollector()
 
   // Find and visit all links
   c.OnHTML("td a.cmc-link", func(e *colly.HTMLElement) {
     e.Request.Visit(e.Attr("href"))
   })
 
   var items []item
   c.OnRequest(func(r *colly.Request) {
		if (strings.Contains(r.URL.String(), "currencies") && len(getGithub(r.URL.String())) > 0){
			// fmt.Println("Visiting", r.URL, getGithub(r.URL.String()))
			split := strings.Split(r.URL.String(), "/")
			items = append(items, item{Name: split[5], Cmc: r.URL.String(),Github: getGithub(r.URL.String())})
		}
   })

   c.Visit("https://coinmarketcap.com/fr/?page=1")

	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("github.json", content, 0644)
}

func getGithub(link string) []gitLink {
	g := colly.NewCollector()
 
	// Find and visit all links
	g.OnHTML("div[data-role='body'] a[href^='https://github.com/']", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
		// fmt.Println("=> "+e.Attr("href"));
	})

	var gitLinks []gitLink
	g.OnRequest(func(l *colly.Request) {
		if (strings.Contains(l.URL.String(), "github")){
			split := strings.Split(l.URL.String(), "/")
			
			firstObject, err := getGitHubInfos(split[3])
			if err != nil {
				fmt.Println("Erreur:", err)
				os.Exit(1)
			}else {
				// fmt.Println("LastActivity:", split[3], firstObject)
				gitLinks = append(gitLinks, gitLink{Username:split[3], Link:l.URL.String(), LastActivity:firstObject})
			}
		}
	})

	g.Visit(link)
	
	return gitLinks
}

func getGitHubInfos(username string) (*githubInfos, error) {
	// URL de l'API
	url := "https://api.github.com/users/"+username+"/events" // ava-labs
	fmt.Println(url)

	// Effectuer la requête HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la requête:", err)
		// os.Exit(1)
	}
	defer resp.Body.Close()

	// Vérifier le code de statut HTTP
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erreur HTTP:", resp.Status)
		// os.Exit(1)
	}

	// Décoder la réponse JSON
	var objects []githubInfos
	if err := json.NewDecoder(resp.Body).Decode(&objects); err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		// os.Exit(1)
	}

	// Vérifier qu'il y a au moins un objet dans la réponse
	if len(objects) == 0 {
		fmt.Println("Aucun objet trouvé dans la réponse")
		// os.Exit(1)
	}

	return &objects[0], nil
}
