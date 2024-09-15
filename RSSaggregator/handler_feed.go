package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/HarshaDannina/GolangPrep/RSSaggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating feed : %v", err))
		return
	}

	respondWithJson(w, 201, databaseFeedToResponse(feed))
}

func (apiConfig *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiConfig.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Can't get feeds: %v", err))
		return
	}
	respondWithJson(w, 200, databaseFeedsToResponse(feeds))
}

func (apiConfig *apiConfig) handlerDeleteFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_id_str := chi.URLParam(r, "feedId")
	feed_id, err := uuid.Parse(feed_id_str)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed id with error: %v", err))
		return
	}
	err = apiConfig.DB.DeleteFeed(r.Context(), database.DeleteFeedParams{
		ID:     feed_id,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Couldn't delete feed: %v", err))
		return
	}
	respondWithJson(w, 200, struct{}{})
}
