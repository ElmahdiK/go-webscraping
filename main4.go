
package main

import (
	"fmt"
	"github.com/nfx/go-htmltable"
)
type Ticker struct {
    Symbol   string `header:"Symbol"`
    Security string `header:"Security"`
}

func main() {
	url := "https://en.wikipedia.org/wiki/List_of_S%26P_500_companies"
	out, _ := htmltable.NewSliceFromURL[Ticker](url)
	fmt.Println("\n--Results--")
	fmt.Println(out[0].Symbol)
	fmt.Println(out[0].Security)

	// Output: 
	// MMM
	// 3M
}