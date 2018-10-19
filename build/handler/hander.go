package handler

import (
	"fmt"
	"net/http"
)

func Hello_Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"status\":\"OK\"}")
}
