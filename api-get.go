package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Structure pour représenter l'objet JSON
type MyObject struct {
	// Ajoutez ici les champs de l'objet JSON que vous attendez
	Type  string `json:"type"`
	Date string `json:"created_at"`
}

func main() {
	// URL de l'API
	url := "https://api.github.com/users/ava-labs/events"

	// Effectuer la requête HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la requête:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Vérifier le code de statut HTTP
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erreur HTTP:", resp.Status)
		os.Exit(1)
	}

	// Décoder la réponse JSON
	var objects []MyObject
	if err := json.NewDecoder(resp.Body).Decode(&objects); err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		os.Exit(1)
	}

	// Vérifier qu'il y a au moins un objet dans la réponse
	if len(objects) == 0 {
		fmt.Println("Aucun objet trouvé dans la réponse")
		os.Exit(1)
	}

	// Afficher le premier objet
	firstObject := objects[0]
	fmt.Println("Premier objet:", firstObject)
}
