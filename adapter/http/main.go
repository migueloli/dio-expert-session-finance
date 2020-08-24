package http

import (
	"net/http"

	"github.com/migueloli/dio-expert-session-finance/adapter/http/actuator"
	"github.com/migueloli/dio-expert-session-finance/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init is used to create the endpoints and start the server.
func Init() {
	http.HandleFunc("/health", actuator.Health)
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.ListenAndServe(":8080", nil)
}
