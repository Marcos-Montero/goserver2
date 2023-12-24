package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MarcosMRod/goserver2/internal/database"
	"github.com/MarcosMRod/goserver2/utils"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err:= decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, "😟 - Err parsing JSON")
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		utils.RespondWithError(w, 500, "😟 - Err creating feed")
		return
	}
	utils.RespondWithJSON(w, 201, utils.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *ApiConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("😟 - Err getting feed follows %v", err))
		return
	}
	utils.RespondWithJSON(w, 200, utils.DatabaseFeedFollowsToFeedFollows(feedFollows))
}
func (apiCfg *ApiConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		utils.RespondWithError(w, 400, "😟 - Err parsing feed follow ID")
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("😟 - Err deleting feed follow: %v", err))
		return
	}
	utils.RespondWithJSON(w, 200, struct{}{})

}
