package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/MarcosMRod/goserver2/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err:= decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "😟 - Err parsing JSON")
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w, 500, "😟 - Err creating user")
		return
	}
	respondWithJSON(w, 201, databaseUserToUser(user))
}
func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r*http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		respondWithError(w, 400, "😟 - Err getting posts for user")
		return
	}
	respondWithJSON(w, 200, databasePostsToPosts(posts))
}