package main

import (
	"fmt"

	"github.com/Dushyantbha012/GO-IMDB/server"
	"github.com/Dushyantbha012/GO-IMDB/store"
)

func main() {
	address := ":6379"
	fmt.Println("Starting GO-IMDB server on", address)

	store := store.NewStore()
	if err := server.Start(address, store); err != nil {
		fmt.Println("Error:", err)
	}
}
