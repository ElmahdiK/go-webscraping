
package main

import (
	"fmt"
	"github.com/nfx/go-htmltable"
)
type Ticker struct {
    Nom   string `header:"Nom"`
    Prix string `header:"Prix"`
}

func main() {
	url := "https://coinmarketcap.com/fr/?page=1"
	out, _ := htmltable.NewSliceFromURL[Ticker](url)
	fmt.Println("\n--Results--")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d > "+out[i].Nom+" - "+out[i].Prix, i)
		fmt.Println("\n")
		fmt.Println(out[i])
    }
}