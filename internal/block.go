package internal

import (
	"crypto/sha256"
	"fmt"
	"strings"
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
	b.Transactions = append(b.Transactions, transactions...)
	// Recalculate hash after adding transactions
	b.Hash = b.calculateHash()
}
