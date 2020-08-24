package transaction

import "time"

// Transaction is the definition of a transaction structure.
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"value"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// Transactions is a split of transactions.
type Transactions []Transaction
