package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MarcosMRod/goserver2/internal/database"
	"github.com/MarcosMRod/goserver2/utils"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err:= decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, "😟 - Err parsing JSON")
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		utils.RespondWithError(w, 500, "😟 - Err creating user")
		return
	}
	utils.RespondWithJSON(w, 201, utils.DatabaseUserToUser(user))
}
func (apiCfg *ApiConfig) HandlerGetUser(w http.ResponseWriter, r*http.Request, user database.User) {
	utils.RespondWithJSON(w, 200, utils.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		utils.RespondWithError(w, 400, "😟 - Err getting posts for user")
		return
	}
	utils.RespondWithJSON(w, 200, utils.DatabasePostsToPosts(posts))
}