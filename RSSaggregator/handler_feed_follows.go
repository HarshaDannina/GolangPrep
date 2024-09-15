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

func (apiConfig *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Can't decode json: %v", err))
		return
	}
	feedFollow, err := apiConfig.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Can't create feed follow: %v", err))
		return
	}
	respondWithJson(w, 201, databaseFeedFollowToResponse(feedFollow))
}

func (apiConfig *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedfollows, err := apiConfig.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Can't get feed follows: %v", err))
		return
	}
	respondWithJson(w, 200, databaseFeedFollowsToResponse(feedfollows))
}

func (apiConfig *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follow_id_str := chi.URLParam(r, "feedFollowId")
	feed_follow_id, err := uuid.Parse(feed_follow_id_str)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id with error: %v", err))
		return
	}
	err = apiConfig.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feed_follow_id,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	respondWithJson(w, 200, struct{}{})
}
