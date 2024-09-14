package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshall JSON response %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "appplication/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	log.Printf("Responding with %v error: %v", code, message)
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errResponse{Error: message})
}
