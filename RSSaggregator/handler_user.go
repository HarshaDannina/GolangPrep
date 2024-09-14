package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/HarshaDannina/GolangPrep/RSSaggregator/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// Create struct for request body
	type parameters struct {
		Name string `json:"name"`
	}
	// Decode body into a json
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	// Marshall json to struct
	err := decoder.Decode(&params)
	// Bad request if error in marshall
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't decode json: %v", err))
		return
	}
	// Create user in database using sqlc generated method
	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't create user: %v", err))
		return
	}
	respondWithJson(w, 200, databaseUserToResponse(user))
}
