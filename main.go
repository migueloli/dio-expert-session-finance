package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", getTransactions)

	http.ListenAndServe(":8080", nil)
}

// Transaction is the definition of a transaction structure.
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"value"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type Transactions []Transaction

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var layout = "2006-01-02T15:04:05"
	newTime, _ := time.Parse(layout, "2020-08-23T20:25:56")

	var transactions = Transactions{
		Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: newTime,
		},
	}

	json.NewEncoder(w).Encode(transactions)
}
