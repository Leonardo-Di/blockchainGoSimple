package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
}

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	PreviousHash string
	Hash         string
	Nonce        int
}

// NewBlock creates a new block
func NewBlock(index int, transactions []Transaction, previousHash string) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PreviousHash: previousHash,
	}
	block.Hash, block.Nonce = block.calculateHash()
	return block
}

// Difficulty level for proof of work
const Difficulty = 4

// calculateHash calculates the hash of the block
func (b *Block) calculateHash() (string, int) {
	nonce := 0
	var hash string
	for {
		record := string(b.Index) + b.Timestamp + b.PreviousHash + string(nonce)
		h := sha256.New()
		h.Write([]byte(record))
		hash = hex.EncodeToString(h.Sum(nil))
		if hash[:Difficulty] == "0000" { // Adjust difficulty here
			break
		}
		nonce++
	}
	return hash, nonce
}
