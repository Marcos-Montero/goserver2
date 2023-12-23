package main

import (
	"net/http"

	"github.com/MarcosMRod/goserver2/internal/auth"
	"github.com/MarcosMRod/goserver2/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User) 

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, "😟 - Err getting API key")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, "😟 - Err getting user")
			return
		}

		handler(w, r, user)
	}
}