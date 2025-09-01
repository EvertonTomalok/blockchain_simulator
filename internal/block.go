package internal

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var nextBlokcId atomic.Int64

type Block struct {
	Index        int64
	Transactions []Transaction
	PrevHash     string
	Hash         string
	Timestamp    time.Time
}

// calculateHash computes the SHA256 hash of the block
func (b *Block) calculateHash() string {
	var txHashes []string
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.GetHash())
	}
	transactionsHash := strings.Join(txHashes, "")

	data := fmt.Sprintf("%d%s%s%d", b.Index, b.PrevHash, transactionsHash, b.Timestamp.Unix())
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// NewBlock creates a new block with the given transactions and previous hash
func NewBlock(transactions []Transaction, prevHash string) *Block {
	block := &Block{
		Index:        nextBlokcId.Add(1),
		Transactions: transactions,
		PrevHash:     prevHash,
		Timestamp:    time.Now(),
	}
	block.Hash = block.calculateHash()
	return block
}

func (b *Block) AddTransaction(transactions []Transaction) {
	for _, transaction := range transactions {
		b.Transactions = append(b.Transactions, transaction)
	}
	// Recalculate hash after adding transactions
	b.Hash = b.calculateHash()
}

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
