package jsonutils

import (
	"encoding/json"
	"net/http"
)

func JSONify(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if response == nil {
		return
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
	}
}
