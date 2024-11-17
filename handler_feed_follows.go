package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Yadav106/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(
	w http.ResponseWriter, r *http.Request, user database.User,
) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
    UserID: user.ID,
    FeedID: params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
    return
	}

	respondWithJSON(w, 201, DatabaseFeedFollowToFeedFollow(feed_follow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(
	w http.ResponseWriter, r *http.Request, user database.User,
) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
    return
	}

	respondWithJSON(w, 201, DatabaseFeedFollowsToFeedFollows(feedFollows))
}

