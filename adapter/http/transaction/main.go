package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/migueloli/dio-expert-session-finance/model/transaction"
	"github.com/migueloli/dio-expert-session-finance/util"
)

// GetTransactions is the endpoint to process a GET request and return a list of transactions.
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	transactions := transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2020-08-23T20:25:56"),
		},
	}

	json.NewEncoder(w).Encode(transactions)
}

// CreateTransaction is the endpoint to process a POST request and create a new transactions.
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res := transaction.Transactions{}

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &res)

	fmt.Print(res)
}
