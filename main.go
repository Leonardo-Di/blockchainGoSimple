package main

import (
	"log"
	"net/http"
)

func main() {
	// Load the blockchain from a file or create a new one
	bc, err := LoadBlockchain("blockchain.json")
	if err != nil {
		log.Println("No existing blockchain found, creating a new one.")
		bc = NewBlockchain()
	}

	// Save the blockchain on exit
	defer func() {
		if err := bc.SaveBlockchain("blockchain.json"); err != nil {
			log.Println("Error saving blockchain:", err)
		}
	}()

	// Set up HTTP routes
	http.HandleFunc("/blocks", bc.getBlocksHandler)
	http.HandleFunc("/transactions", bc.addTransactionHandler)
	http.HandleFunc("/mine", bc.mineBlockHandler)

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
