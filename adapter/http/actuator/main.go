package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody is a struct to return the status health for the application.
type HealthBody struct {
	Status string `json:"status"`
}

// Health is the endpoint to verify if the application is working ok.
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	status := HealthBody{"alive"}

	json.NewEncoder(w).Encode(status)
}
