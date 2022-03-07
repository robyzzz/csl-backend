package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robyzzz/csl-backend/model"
	"github.com/robyzzz/csl-backend/utils"
)

// GET /api/player/{steamid} - Returns player's game stats from database
func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	player, err := model.GetPlayer(mux.Vars(r)["steamid"])
	if err != nil {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusNotFound, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(player)
}
