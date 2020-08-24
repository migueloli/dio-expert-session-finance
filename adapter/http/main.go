package http

import (
	"net/http"

	"github.com/migueloli/dio-expert-session-finance.git/adapter/http/actuator"
	"github.com/migueloli/dio-expert-session-finance.git/adapter/http/transaction"
)

// Init is used to create the endpoints and start the server.
func Init() {
	http.HandleFunc("/health", actuator.Health)

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.ListenAndServe(":8080", nil)
}
