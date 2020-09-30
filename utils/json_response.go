package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse : return json response
func JSONResponse(statusCode int, response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
