package internal

import (
	"sync"
	"time"
)

// Blockchain represents a simple blockchain
type Blockchain struct {
	mu     sync.Mutex
	Blocks []Block
}

// NewBlockchain creates a new blockchain with a genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := &Block{
		Index:        0,
		Transactions: []Transaction{},
		PrevHash:     "0",
		Timestamp:    time.Now(),
	}
	genesisBlock.Hash = genesisBlock.calculateHash()

	return &Blockchain{
		Blocks: []Block{*genesisBlock},
	}
}

// AddBlock adds a new block with the given transactions to the blockchain
func (bc *Blockchain) AddBlock(transactions []Transaction) {
	defer bc.mu.Unlock()
	bc.mu.Lock()

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, *newBlock)
}

// IsValid verifies the integrity of the blockchain
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		// Check if current block's hash is valid
		if currentBlock.Hash != currentBlock.calculateHash() {
			return false
		}

		// Check if current block's previous hash matches the previous block's hash
		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}

		// Check if block index is sequential
		if currentBlock.Index != prevBlock.Index+1 {
			return false
		}
	}

	// Verify genesis block
	genesisBlock := bc.Blocks[0]
	if genesisBlock.Index != 0 || genesisBlock.PrevHash != "0" {
		return false
	}

	return true
}
