package controller

import (
	"encoding/json"
	"net/http"

	"github.com/robyzzz/csl-backend/config"
	"github.com/solovev/steam_go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	opId := steam_go.NewOpenId(r)

	switch opId.Mode() {
	case "cancel":
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	case "":
		// redirect to steam auth
		http.Redirect(w, r, opId.AuthUrl(), http.StatusTemporaryRedirect)
	default:
		// login success

		// get user
		user, err := opId.ValidateAndGetUser(config.STEAM_API_KEY)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		config.CreateSessionID(w, r, user.SteamId)

		// TODO: store user info in database

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}