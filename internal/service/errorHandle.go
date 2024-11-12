package service

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func ErrorHandling(err string, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")

	response := ErrorResponse{Error: err}
	json.NewEncoder(w).Encode(response)
}
