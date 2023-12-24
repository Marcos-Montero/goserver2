package server

import (
	"net/http"

	"github.com/MarcosMRod/goserver2/utils"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 500, "😟 - Something went wrong")
}