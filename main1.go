package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL du site à analyser
	url := "https://www.pokepedia.fr/Fantominus"

	// Charger le document HTML de l'URL
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal("Erreur lors du chargement de la page :", err)
	}

	// Sélectionner la table avec la classe spécifiée
	table := doc.Find("table.liensrougesreduits")

	// Parcourir les lignes de la table
	table.Find("tr").Each(func(i int, row *goquery.Selection) {
		// Sélectionner l'image dans la colonne correspondante
		image := row.Find("td a img")

		// Récupérer l'attribut src de l'image
		src, exists := image.Attr("src")
		if exists {
			fmt.Println(src)
		}
	})
}
