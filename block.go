package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Index        int
	TimeStamp    string
	Transactions []Transaction
	PrevHash     string
	Hash         string
}

// CalculateHash calculates the hash of a block
func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.TimeStamp + block.PrevHash
	for _, tx := range block.Transactions {
		record += tx.TimeStamp + tx.From + tx.To + tx.Amount
	}
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// CreateBlock creates a new block
func CreateBlock(prevBlock Block, transactions []Transaction) Block {
	newBlock := Block{
		Index:        prevBlock.Index + 1,
		TimeStamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     prevBlock.Hash,
	}
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock
}
