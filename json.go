package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Printf("Responding with %v error: %v", status, msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w,status, errorResponse{
		Error: msg,
	} )
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}