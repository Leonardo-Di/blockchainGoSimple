package main

import (
	"encoding/json"
	"net/http"
)

func (bc *Blockchain) getBlocksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bc.Chain)
}

func (bc *Blockchain) addTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username, password, ok := r.BasicAuth()
	if !ok || !Authenticate(username, password) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	bc.AddTransaction(transaction.Sender, transaction.Recipient, transaction.Amount)
	LogInfo("Transaction added: " + transaction.Sender + " -> " + transaction.Recipient)
	w.WriteHeader(http.StatusAccepted)
}

func (bc *Blockchain) mineBlockHandler(w http.ResponseWriter, r *http.Request) {
	bc.MineBlock()
	LogInfo("New block mined")
	w.WriteHeader(http.StatusAccepted)
}
