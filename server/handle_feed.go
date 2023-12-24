package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MarcosMRod/goserver2/internal/database"
	"github.com/MarcosMRod/goserver2/utils"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`

	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err:= decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, "😟 - Err parsing JSON")
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, 500, "😟 - Err creating feed")
		return
	}
	utils.RespondWithJSON(w, 201, utils.DatabaseFeedToFeed(feed))
}
func (apiCfg *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		utils.RespondWithError(w, 500, "😟 - Err getting feeds")
		return
	}
	utils.RespondWithJSON(w, 201, utils.DatabaseFeedsToFeeds(feeds))
}
