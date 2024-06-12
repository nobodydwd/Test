package main

import (
	"fmt"
	"time"
)

// Initialize the blockchain with the genesis block
func initBlockchain() []Block {
	genesisBlock := Block{
		Index:        0,
		Transactions: []Transaction{},
		TimeStamp:    time.Now().String(),
		PrevHash:     "",
		Hash:         "",
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	return []Block{genesisBlock}
}

func main() {
	Blockchain := initBlockchain()

	// Simulate user input for a transfer
	from := "alex"
	to := "nasry"
	amount := 100

	// Perform the transfer
	err := Transfer(from, to, amount)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a new transaction and add it to a new block
	transaction := NewTransaction(from, to, fmt.Sprintf("%d", amount))
	newBlock := CreateBlock(Blockchain[len(Blockchain)-1], []Transaction{transaction})
	Blockchain = append(Blockchain, newBlock)

	// Print out the blockchain
	for _, block := range Blockchain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.TimeStamp)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("Transactions:")
		for _, tx := range block.Transactions {
			fmt.Printf(" - From: %s, To: %s, Amount: %s, Time: %s\n", tx.From, tx.To, tx.Amount, tx.TimeStamp)
		}
		fmt.Println()
	}

	fmt.Println("nice")

	// Start the JSON-RPC server
	go startRPCServer()

	// Prevent the main function from exiting immediately
	select {}
}
