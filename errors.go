package main

import "fmt"

type BlockchainError struct {
	Message string
}

func (e *BlockchainError) Error() string {
	return fmt.Sprintf("Blockchain Error: %s", e.Message)
}
