package internal

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TransactionProducer generates random transactions
type TransactionProducer struct {
	pool        *TransactionPool
	interval    time.Duration
	users       []string
	isRunning   bool
	mu          sync.Mutex
	stopChannel chan bool
}

// NewTransactionProducer creates a new transaction producer
func NewTransactionProducer(pool *TransactionPool, intervalMs int) *TransactionProducer {
	users := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}

	return &TransactionProducer{
		pool:        pool,
		interval:    time.Duration(intervalMs) * time.Millisecond,
		users:       users,
		stopChannel: make(chan bool),
	}
}

// Start begins producing transactions
func (tp *TransactionProducer) Start() {
	tp.mu.Lock()
	if tp.isRunning {
		tp.mu.Unlock()
		return
	}
	tp.isRunning = true
	tp.mu.Unlock()

	go tp.produce()
}

// Stop stops producing transactions
func (tp *TransactionProducer) Stop() {
	tp.mu.Lock()
	if !tp.isRunning {
		tp.mu.Unlock()
		return
	}
	tp.isRunning = false
	tp.mu.Unlock()

	tp.stopChannel <- true
}

// produce generates random transactions at regular intervals
func (tp *TransactionProducer) produce() {
	ticker := time.NewTicker(tp.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Generate random transaction
			from := tp.users[rand.Intn(len(tp.users))]
			to := tp.users[rand.Intn(len(tp.users))]

			// Ensure from and to are different
			for to == from {
				to = tp.users[rand.Intn(len(tp.users))]
			}

			amount := rand.Float64() * 100 // Random amount between 0-100
			tx := NewTransaction(from, to, amount)

			fmt.Printf("Generated transaction: %s -> %s: %.2f\n", from, to, amount)
			tp.pool.AddTransaction(tx)

		case <-tp.stopChannel:
			return
		}
	}
}
