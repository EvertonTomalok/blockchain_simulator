package internal

import (
	"crypto/sha256"
	"fmt"
	"sync/atomic"
	"time"
)

var nextTransactionId atomic.Int64

type Transaction struct {
	ID        int64
	From      string
	To        string
	Amount    float64
	hash      string
	Timestamp time.Time
}

func (t *Transaction) setHash() {
	data := fmt.Sprintf("%d%s%s%f%d", t.ID, t.From, t.To, t.Amount, t.Timestamp.Unix())
	hash := sha256.Sum256([]byte(data))
	t.hash = fmt.Sprintf("%x", hash)
}

func (t *Transaction) GetHash() string {
	return t.hash
}

func NewTransaction(from, to string, amount float64) *Transaction {
	block := &Transaction{
		ID:        nextTransactionId.Add(1),
		From:      from,
		To:        to,
		Amount:    amount,
		Timestamp: time.Now(),
	}
	block.setHash()

	return block
}
