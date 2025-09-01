package internal

import (
	"fmt"
	"sync"
	"time"
)

// TransactionPool manages incoming transactions and batches them into blocks
type TransactionPool struct {
	txChannel   chan *Transaction
	blockchain  *Blockchain
	batchSize   int
	mu          sync.Mutex
	isRunning   bool
	stopChannel chan bool
}

// NewTransactionPool creates a new transaction pool
func NewTransactionPool(blockchain *Blockchain, batchSize int) *TransactionPool {
	return &TransactionPool{
		txChannel:   make(chan *Transaction, 100), // Buffered channel
		blockchain:  blockchain,
		batchSize:   batchSize,
		stopChannel: make(chan bool),
	}
}

// Start begins the transaction processing
func (tp *TransactionPool) Start() {
	tp.mu.Lock()
	if tp.isRunning {
		tp.mu.Unlock()
		return
	}
	tp.isRunning = true
	tp.mu.Unlock()

	go tp.consumer()
}

// Stop stops the transaction processing
func (tp *TransactionPool) Stop() {
	tp.mu.Lock()
	if !tp.isRunning {
		tp.mu.Unlock()
		return
	}
	tp.isRunning = false
	tp.mu.Unlock()

	tp.stopChannel <- true
	close(tp.txChannel)
}

// AddTransaction adds a transaction to the pool
func (tp *TransactionPool) AddTransaction(tx *Transaction) {
	select {
	case tp.txChannel <- tx:
		// Transaction added successfully
	default:
		fmt.Println("Transaction pool is full, dropping transaction")
	}
}

// consumer batches transactions and creates blocks
func (tp *TransactionPool) consumer() {
	batch := make([]Transaction, 0, tp.batchSize)
	ticker := time.NewTicker(5 * time.Second) // Timeout for partial batches
	defer ticker.Stop()

	for {
		select {
		case tx, ok := <-tp.txChannel:
			if !ok {
				// Channel closed, process remaining batch if any
				if len(batch) > 0 {
					tp.createBlock(batch)
				}
				return
			}

			batch = append(batch, *tx)
			fmt.Printf("Added transaction to batch. Batch size: %d/%d\n", len(batch), tp.batchSize)

			if len(batch) >= tp.batchSize {
				tp.createBlock(batch)
				batch = make([]Transaction, 0, tp.batchSize)
				ticker.Reset(5 * time.Second)
			}

		case <-ticker.C:
			// Timeout - create block with partial batch if any transactions exist
			if len(batch) > 0 {
				fmt.Printf("Timeout reached, creating block with %d transactions\n", len(batch))
				tp.createBlock(batch)
				batch = make([]Transaction, 0, tp.batchSize)
			}

		case <-tp.stopChannel:
			// Process remaining batch before stopping
			if len(batch) > 0 {
				tp.createBlock(batch)
			}
			return
		}
	}
}

// createBlock creates a new block with the given transactions
func (tp *TransactionPool) createBlock(transactions []Transaction) {
	tp.blockchain.AddBlock(transactions)
	fmt.Printf("Created new block with %d transactions. Blockchain length: %d\n",
		len(transactions), len(tp.blockchain.Blocks))
}
