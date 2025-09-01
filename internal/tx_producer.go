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
	users       []string
	isRunning   bool
	mu          sync.Mutex
	stopChannel chan bool
}

// NewTransactionProducer creates a new transaction producer
func NewTransactionProducer(pool *TransactionPool) *TransactionProducer {
	users := []string{
		"Everton", "Amanda", "Gabriel", "Marc", "Mavie",
		"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank",
		"Grace", "Henry",
	}

	return &TransactionProducer{
		pool:        pool,
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

	var wg sync.WaitGroup
	for range 5 { // 5 workers
		go tp.produce(&wg)
	}
	wg.Wait()
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
func (tp *TransactionProducer) produce(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	intervalMs := getRandomMs(400, 800)
	ticker := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
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

func getRandomMs(min, max int) int {
	rand.Seed(time.Now().Unix())
	possibilities := make([]int, max-min)
	for i := min; i < max; i++ {
		possibilities = append(possibilities, i)
	}
	choice := possibilities[rand.Intn(len(possibilities))]

	if choice < min {
		return min
	}
	return choice
}
