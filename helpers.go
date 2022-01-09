package main

import (
	"encoding/json"
	"net/http"
)

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func responseErrorWithJSON(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorHandle{
		ErrorStatus:  code,
		ErrorMessage: message,
		ErrorInput:   nil,
		ErrorFrom:    "",
	})
}
