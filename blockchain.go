package main

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type Blockchain struct {
	Chain        []*Block
	Transactions []Transaction
	mu           sync.Mutex
}

// LoadBlockchain loads the blockchain from a file
func LoadBlockchain(filename string) (*Blockchain, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var chain []*Block
	if err := json.Unmarshal(data, &chain); err != nil {
		return nil, err
	}
	return &Blockchain{Chain: chain}, nil
}

// SaveBlockchain saves the blockchain to a file
func (bc *Blockchain) SaveBlockchain(filename string) error {
	data, err := json.Marshal(bc.Chain)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (bc *Blockchain) AddTransaction(sender, recipient string, amount float64) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	transaction := Transaction{Sender: sender, Recipient: recipient, Amount: amount}
	bc.Transactions = append(bc.Transactions, transaction)
}

// MineBlock mines a new block and adds it to the blockchain
func (bc *Blockchain) MineBlock() {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	newBlock := NewBlock(len(bc.Chain), bc.Transactions, bc.Chain[len(bc.Chain)-1].Hash)
	bc.Chain = append(bc.Chain, newBlock)
	bc.Transactions = []Transaction{} // Reset transactions
}
