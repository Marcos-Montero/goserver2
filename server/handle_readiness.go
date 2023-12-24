package server

import (
	"net/http"

	"github.com/MarcosMRod/goserver2/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, map[string]string{"status": "ok"})
}