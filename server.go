package main

import (
	"encoding/json"
	"net/http"
)

// getBlocksHandler handles HTTP requests to get the blockchain
func (bc *Blockchain) getBlocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bc.Chain)
}

// addTransactionHandler handles HTTP requests to add a transaction
func (bc *Blockchain) addTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Basic authentication
	username, password, ok := r.BasicAuth()
	if !ok || !Authenticate(username, password) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	bc.AddTransaction(transaction.Sender, transaction.Recipient, transaction.Amount)
	w.WriteHeader(http.StatusAccepted)
}

// mineBlockHandler handles HTTP requests to mine a new block
func (bc *Blockchain) mineBlockHandler(w http.ResponseWriter, r *http.Request) {
	bc.MineBlock()
	w.WriteHeader(http.StatusAccepted)
}
