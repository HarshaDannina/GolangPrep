package main

import (
	"fmt"
	"net/http"

	"github.com/HarshaDannina/GolangPrep/RSSaggregator/internal/auth"
	"github.com/HarshaDannina/GolangPrep/RSSaggregator/internal/database"
)

func (apiConfig *apiConfig) middlewareAuth(handler func(http.ResponseWriter, *http.Request, database.User)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiConfig.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 500, fmt.Sprintf("Auth error: %v", err))
			return
		}
		handler(w, r, user)
	}
}
