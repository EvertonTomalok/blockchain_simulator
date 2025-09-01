package cmd

import (
	"fmt"
	"time"

	"github.com/EvertonTomalok/blockchain_simulator/internal"
)

func Execute() {
	// Initialize blockchain
	blockchain := internal.NewBlockchain()
	fmt.Println("Blockchain initialized with genesis block")

	// Create transaction pool (batch size = 10 transactions)
	pool := internal.NewTransactionPool(blockchain, 10)

	// Create transaction producer (generates transactions)
	producer := internal.NewTransactionProducer(pool)

	// Start the system
	pool.Start()
	producer.Start()

	fmt.Println("Transaction pool and producer started...")
	fmt.Println("Press Ctrl+C to stop or wait 20 seconds for demo")

	// Run for 20 seconds for demonstration
	time.Sleep(20 * time.Second)

	// Stop the system
	producer.Stop()
	pool.Stop()

	// Display final blockchain state
	fmt.Printf("\n\n-----------------\n\n-> Final blockchain state <-\n")
	fmt.Printf("Number of blocks: %d\n", len(blockchain.Blocks))
	fmt.Printf("Blockchain is valid: %t\n\n-----------------\n\n", blockchain.IsValid())

	// Display all blocks
	for i, block := range blockchain.Blocks {
		fmt.Printf("Block %d: %d transactions, Hash: %s\n",
			i, len(block.Transactions), block.Hash[:16]+"...")
	}
}
